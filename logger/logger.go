package logger

import (
	"io"
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

func New(out io.Writer) LogWriter {

	if out == nil {
		out = os.Stdout
	}

	l := new (Logger)
	l.ERROR = log.New(out, "[ERROR] ", 0)
	l.CRITICAL = log.New(out, "[CRIT] ", 0)
	l.WARN = log.New(out, "[WARN]  ", 0)
	l.DEBUG = log.New(out, "[DEBUG] ", 0)
	l.INFO = log.New(out, "[INFO] ", 0)
	l.NONE = log.New(out, "", 0)

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
