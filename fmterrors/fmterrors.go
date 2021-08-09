package fmterrors

import (
	"errors"
	"fmt"
)

func Error(args ... interface{}) error {
	return errors.New(fmt.Sprint(args...))
}

func Errorf(format string, args ... interface{}) error {
	return errors.New(fmt.Sprintf(format, args...))
}
