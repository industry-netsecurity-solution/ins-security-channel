package ins

import (
	"bytes"
	"fmt"
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
