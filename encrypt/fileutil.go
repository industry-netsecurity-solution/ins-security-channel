package encrypt

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
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
		nonce = make([]byte, 12)
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			return 0, err
		}
		logger.Debugf("Nonce 생성: %s", hex.EncodeToString(nonce))
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
		if nr > 0 {
			if writeNonce {
				writeNonce = false
				outBuf = aesgcm.Seal(nonce, nonce, buf[:nr], nil)

				logger.Debugf("Nonce: %s", hex.EncodeToString(nonce))
				logger.Debugf("Input(plaintext): %s", hex.EncodeToString(buf[:nr]))
				logger.Debugf("Output(encrypted): %s", hex.EncodeToString(outBuf))
			} else {
				outBuf = aesgcm.Seal(nil, nonce, buf[:nr], nil)
			}
			logger.Debugf("Block: %d byte", nr)
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
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
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

		if nr > 0 {
			if nonce == nil {
				nonceslice, ciphertext := buf[:nonceSize], buf[nonceSize:nr]
				nonce = make([]byte, len(nonceslice))
				copy(nonce, nonceslice)
				logger.Debugf("Nonce 확인: %s", hex.EncodeToString(nonce))

				outBuf, err = aesgcm.Open(nil, nonce, ciphertext, nil)
				if err != nil {
					logger.Error(err)
				}
				logger.Debugf("Nonce: %s", hex.EncodeToString(nonce))
				logger.Debugf("Input(encrypted): %s", hex.EncodeToString(ciphertext))
				logger.Debugf("Output(plaintext): %s", hex.EncodeToString(outBuf))
			} else {
				outBuf, err = aesgcm.Open(nil, nonce, buf[:nr], nil)
				if err != nil {
					logger.Error(err)
				}
			}
			logger.Debugf("Block: %d byte", nr)
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
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
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
