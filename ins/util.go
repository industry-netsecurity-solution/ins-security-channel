package ins

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

func TimeYYmmddHHMMSS(t *time.Time) string {
	if t == nil {
		tm := time.Now()
		t = &tm;
	}

	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

//func GetType(myvar interface{}) string {
//	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
//		return "*" + t.Elem().Name()
//	} else {
//		return t.Name()
//	}
//}

func GetType(myvar interface{}) string {
	return reflect.TypeOf(myvar).String()
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func ReadFile(filepath string) (*bytes.Buffer, error) {
	var file *os.File = nil
	var err error = nil
	if file, err = os.Open(filepath); err != nil {
		return nil, err
	}
	defer file.Close()

	var n int64 = bytes.MinRead
	if fi, err := file.Stat(); err == nil {
		if size := fi.Size() + bytes.MinRead; size > n {
			n = size
		}

	}

	buffer := new(bytes.Buffer)
	if int64(int(n)) == n {
		buffer.Grow(int(n))
	}
	_, err = buffer.ReadFrom(file)

	return buffer, nil
}

func SumHash(filepath string, hash hash.Hash) (string, error) {
	var file *os.File = nil
	var err error
	if file, err = os.Open(filepath); err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func SumSHA256(filepath string) (string, error) {
	hash := sha256.New()
	return SumHash(filepath, hash)
}

func SumMD5(filepath string) (string, error) {
	hash := md5.New()
	return SumHash(filepath, hash)
}

func Pack(args...interface{}) interface{} {
	if args == nil {
		return nil
	}
	if 0 < len(args) {
		return args[0]
	}

	return nil
}

func PackAsString(args...interface{}) string {
	if str, err := AsString(Pack(args...)); err == nil {
		return str
	}

	return ""
}