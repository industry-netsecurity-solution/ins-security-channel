package ins

import (
	"bytes"
	"fmt"
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

	buffer := new(bytes.Buffer)
	buf := make([]byte, 4096)
	for {
		var rlen int
		if rlen, err = file.Read(buf); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if _, err = buffer.Write(buf[:rlen]); err != nil {
			return nil, err
		}
	}

	return buffer, nil
}
