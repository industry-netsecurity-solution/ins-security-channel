package ins

import (
	"bytes"
	"container/list"
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
		t = &tm
	}

	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

func TimeYYmmdd(t *time.Time) string {
	if t == nil {
		tm := time.Now()
		t = &tm
	}

	return fmt.Sprintf("%04d-%02d-%02d",
		t.Year(), t.Month(), t.Day())
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

func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
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

func Pack(args ...interface{}) interface{} {
	if args == nil {
		return nil
	}
	if 0 < len(args) {
		return args[0]
	}

	return nil
}

func PackAsString(args ...interface{}) string {
	if str, err := AsString(Pack(args...)); err == nil {
		return str
	}

	return ""
}

func ListToStringArray(src *list.List) []string {
	results := make([]string, src.Len())

	var i int = 0
	for e := src.Front(); e != nil; e = e.Next() {
		results[i] = e.Value.(string)
		i++
	}

	return results
}

func ListToSlice(src *list.List) []interface{} {
	results := make([]interface{}, src.Len())

	var i int = 0
	for e := src.Front(); e != nil; e = e.Next() {
		results[i] = e.Value
		i++
	}

	return results
}

func SliceToList(src []interface{}) *list.List {
	results := list.New()
	for _, e := range src {
		results.PushBack(e)
	}

	return results
}

func AppendSliceToList(dst *list.List, src []interface{}) *list.List {
	for _, e := range src {
		dst.PushBack(e)
	}

	return dst
}

func AppendToList(dst *list.List, src ...interface{}) *list.List {
	for _, e := range src {
		dst.PushBack(e)
	}

	return dst
}

func StringJoin(elems []string, sep string, filter func(string) string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}

	var b bytes.Buffer

	if filter == nil {
		b.WriteString(elems[0])
	} else {
		b.WriteString(filter(elems[0]))
	}

	for _, s := range elems[1:] {
		b.WriteString(sep)
		if filter == nil {
			b.WriteString(s)
		} else {
			b.WriteString(filter(s))
		}
	}
	return b.String()
}
