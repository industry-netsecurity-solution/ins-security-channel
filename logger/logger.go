package logger

import (
	"os"
	"log"
)

type LogWriter interface {
	Critiacal(args ... interface{})
	Error(args ... interface{})
	Warning(args ... interface{})
	Debug(args ... interface{})
	Info(args ... interface{})
	Println(args ... interface{})
}

type Logger struct {
	CRITICAL *log.Logger
	ERROR *log.Logger
	WARN *log.Logger
	DEBUG *log.Logger
	INFO *log.Logger
	NONE *log.Logger
}

func New() LogWriter {
	l := new (Logger)
	l.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	l.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	l.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	l.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)
	l.INFO = log.New(os.Stdout, "[INFO] ", 0)
	l.NONE = log.New(os.Stdout, "", 0)

	return l
}

func (v Logger) Critiacal(args ... interface{}) {
	v.CRITICAL.Println(args)
}

func (v Logger) Error(args ... interface{}) {
	v.ERROR.Println(args)
}

func (v Logger) Warning(args ... interface{}) {
	v.WARN.Println(args)
}

func (v Logger) Debug(args ... interface{}) {
	v.DEBUG.Println(args)
}

func (v Logger) Info(args ... interface{}) {
	v.INFO.Println(args)
}

func (v Logger) Println(args ... interface{}) {
	v.NONE.Println(args)
}
