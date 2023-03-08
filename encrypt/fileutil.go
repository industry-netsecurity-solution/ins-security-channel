package encrypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/industry-netsecurity-solution/ins-security-channel/filesystem"
	"github.com/industry-netsecurity-solution/ins-security-channel/logger"
	"io"
	"os"
)

func Encrypt(dst io.Writer, src io.Reader, aesgcm cipher.AEAD) (written int64, err error) {
	return EncryptBuffer(dst, src, nil, aesgcm)
}

// copyBuffer is the actual implementation of Copy and CopyBuffer.
// if buf is nil, one is allocated.
func EncryptBuffer(dst io.Writer, src io.Reader, buf []byte, aesgcm cipher.AEAD) (written int64, err error) {

	var nonce []byte = nil
	if aesgcm == nil {
		// If the reader has a WriteTo method, use it to do the copy.
		// Avoids an allocation and a copy.
		if wt, ok := src.(io.WriterTo); ok {
			return wt.WriteTo(dst)
		}
		// Similarly, if the writer has a ReadFrom method, use it to do the copy.
		if rt, ok := dst.(io.ReaderFrom); ok {
			return rt.ReadFrom(src)
		}
	} else {
		// make random nonce
		nonceSize := aesgcm.NonceSize()
		nonce = make([]byte, nonceSize)
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			return 0, err
		}
	}

	writeNonce := true
	if buf == nil {
		size := 32 * 1024
		if l, ok := src.(*io.LimitedReader); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]byte, size)
	}
	for {
		var outBuf []byte = nil
		nr, er := src.Read(buf)
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}

		if nr < 0 {
			break
		} else if nr > 0 && nr <= len(buf) {
			tmpBuf := make([]byte, nr)
			copy(tmpBuf, buf)
			if writeNonce {
				writeNonce = false
				outBuf = aesgcm.Seal(nonce, nonce, tmpBuf, nil)
			} else {
				outBuf = aesgcm.Seal(nil, nonce, tmpBuf, nil)
			}
			nr = len(outBuf)

			nw, ew := dst.Write(outBuf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
				if ew == nil {
					ew = errors.New("invalid write result")
				}
			}
			written += int64(nw)
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
	}
	return written, err
}

func Decrypt(dst io.Writer, src io.Reader, aesgcm cipher.AEAD) (written int64, err error) {
	return DecryptBuffer(dst, src, nil, aesgcm)
}

// copyBuffer is the actual implementation of Copy and CopyBuffer.
// if buf is nil, one is allocated.
func DecryptBuffer(dst io.Writer, src io.Reader, buf []byte, aesgcm cipher.AEAD) (written int64, err error) {

	var nonce []byte
	if aesgcm == nil {
		// If the reader has a WriteTo method, use it to do the copy.
		// Avoids an allocation and a copy.
		if wt, ok := src.(io.WriterTo); ok {
			return wt.WriteTo(dst)
		}
		// Similarly, if the writer has a ReadFrom method, use it to do the copy.
		if rt, ok := dst.(io.ReaderFrom); ok {
			return rt.ReadFrom(src)
		}
	}

	nonceSize := aesgcm.NonceSize()

	blocksize := (32 * 1024) + 16
	if buf == nil {
		size := blocksize + 12
		if l, ok := src.(*io.LimitedReader); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]byte, size)
	}

	for {
		var outBuf []byte = nil
		var nr int
		var er error
		if nonce == nil {
			nr, er = src.Read(buf)
		} else {
			nr, er = src.Read(buf[:blocksize])
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}

		if nr < 0 {
			break
		} else if nr > 0 && nr <= len(buf) {
			tmpBuf := make([]byte, nr)
			copy(tmpBuf, buf)

			if nonce == nil {
				//nonceslice, encrypteddata := buf[:nonceSize], buf[nonceSize:nr]

				var typeBuffer = new(bytes.Buffer)
				typeBuffer.Write(buf)

				nonce = make([]byte, nonceSize)
				typeBuffer.Read(nonce)

				encrypteddata := make([]byte, nr-nonceSize)
				typeBuffer.Read(encrypteddata)

				outBuf, err = aesgcm.Open(nil, nonce, encrypteddata, nil)
				if err != nil {
					logger.Error(err)
				}
			} else {
				outBuf, err = aesgcm.Open(nil, nonce, tmpBuf, nil)
				if err != nil {
					logger.Error(err)
				}
			}
			nr = len(outBuf)

			nw, ew := dst.Write(outBuf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
				if ew == nil {
					ew = errors.New("invalid write result")
				}
			}
			written += int64(nw)
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
	}
	return written, err
}

func EncryptFile(sourcePath, destPath string, aesgcm cipher.AEAD, overwrite bool) (bool, error) {
	_, err := os.Stat(destPath)
	if overwrite == false && err == nil {
		return false, err
	}

	inputFile, err := os.OpenFile(sourcePath, os.O_RDONLY, 0644)
	if err != nil {
		return false, fmt.Errorf("couldn't open source file: %s", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(destPath)
	if err != nil {
		return false, fmt.Errorf("couldn't open dest file: %s", err)
	}

	if _, err = Encrypt(outputFile, inputFile, aesgcm); err != nil {
		outputFile.Close()
		return false, err
	}

	if err = outputFile.Sync(); err != nil {
		outputFile.Close()
		return false, fmt.Errorf("sync to output file failed: %s", err)
	}

	if err = outputFile.Close(); err != nil {
		return false, fmt.Errorf("writing to output file failed: %s", err)
	}

	atime, mtime, _, _, err := filesystem.StatTimes(sourcePath)
	if err != nil {
		return true, err
	}

	if _, err = os.Stat(destPath); err != nil {
		fmt.Println(err)
		return true, err
	}

	if err = os.Chtimes(destPath, atime, mtime); err != nil {
		return true, err
	}

	return true, nil
}

func DecryptFile(sourcePath, destPath string, aesgcm cipher.AEAD, overwrite bool) (bool, error) {
	_, err := os.Stat(destPath)
	if overwrite == false && err == nil {
		return false, err
	}

	inputFile, err := os.OpenFile(sourcePath, os.O_RDONLY, 0644)
	if err != nil {
		return false, fmt.Errorf("couldn't open source file: %s", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(destPath)
	if err != nil {
		return false, fmt.Errorf("couldn't open dest file: %s", err)
	}

	if _, err = Decrypt(outputFile, inputFile, aesgcm); err != nil {
		outputFile.Close()
		return false, err
	}

	if err = outputFile.Sync(); err != nil {
		outputFile.Close()
		return false, fmt.Errorf("sync to output file failed: %s", err)
	}

	if err = outputFile.Close(); err != nil {
		return false, fmt.Errorf("writing to output file failed: %s", err)
	}

	atime, mtime, _, _, err := filesystem.StatTimes(sourcePath)
	if err != nil {
		return true, err
	}

	if _, err = os.Stat(destPath); err != nil {
		fmt.Println(err)
		return true, err
	}

	if err = os.Chtimes(destPath, atime, mtime); err != nil {
		return true, err
	}

	return true, nil
}
