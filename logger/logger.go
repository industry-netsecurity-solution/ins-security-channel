package logger

import (
	"io"
	"log"
	"os"
)

type LogWriter interface {
	Error(args ... interface{})
	Errorf(format string, args ... interface{})
	Errorln(args ... interface{})
	Warning(args ... interface{})
	Warningf(format string, args ... interface{})
	Warningln(args ... interface{})
	Debug(args ... interface{})
	Debugf(format string, args ... interface{})
	Debugln(args ... interface{})
	Info(args ... interface{})
	Infof(format string, args ... interface{})
	Infoln(args ... interface{})
	Print(args ... interface{})
	Printf(format string, args ... interface{})
	Println(args ... interface{})
}

type Logger struct {
	ERROR *log.Logger
	WARN *log.Logger
	DEBUG *log.Logger
	INFO *log.Logger
	NONE *log.Logger
}

func New(out io.Writer, flag int) *Logger {

	if out == nil {
		out = os.Stdout
	}

	l := new (Logger)
	l.ERROR = log.New(out, "[ERROR] ", flag)
	l.WARN = log.New(out, "[WARN]  ", flag)
	l.DEBUG = log.New(out, "[DEBUG] ", flag)
	l.INFO = log.New(out, "[INFO] ", flag)
	l.NONE = log.New(out, "", flag)

	return l
}

func (v Logger) Writer() LogWriter {
	return LogWriter(v)
}

func (v Logger) SetOutput(w io.Writer) {
	v.ERROR.SetOutput(w)
	v.WARN.SetOutput(w)
	v.DEBUG.SetOutput(w)
	v.INFO.SetOutput(w)
	v.NONE.SetOutput(w)
}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func (v Logger) Fatal(args ...interface{}) {
	v.NONE.Fatal(args...)
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func (v Logger) Fatalf(format string, args ...interface{}) {
	v.NONE.Fatalf(format, args...)
}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func (v Logger) Fatalln(args ...interface{}) {
	v.NONE.Fatalln(args...)
}

// Panic is equivalent to l.Print() followed by a call to panic().
func (v Logger) Panic(args ...interface{}) {
	v.NONE.Panic(args...)
}

// Panicf is equivalent to l.Printf() followed by a call to panic().
func (v Logger) Panicf(format string, args ...interface{}) {
	v.NONE.Panicf(format, args...)
}

// Panicln is equivalent to l.Println() followed by a call to panic().
func (v Logger) Panicln(args ...interface{}) {
	v.NONE.Panicln(args...)
}

// Flags returns the output flags for the logger.
// The flag bits are Ldate, Ltime, and so on.
func (v Logger) Flags() int {
	return v.NONE.Flags()
}

// SetFlags sets the output flags for the logger.
// The flag bits are Ldate, Ltime, and so on.
func (v Logger) SetFlags(flag int) {
	v.ERROR.SetFlags(flag)
	v.WARN.SetFlags(flag)
	v.DEBUG.SetFlags(flag)
	v.INFO.SetFlags(flag)
	v.NONE.SetFlags(flag)
}

func (v Logger) Error(args ... interface{}) {
	v.ERROR.Print(args...)
}

func (v Logger) Errorf(format string, args ... interface{}) {
	v.ERROR.Printf(format, args...)
}

func (v Logger) Errorln(args ... interface{}) {
	v.ERROR.Println(args...)
}

func (v Logger) Warning(args ... interface{}) {
	v.WARN.Print(args...)
}

func (v Logger) Warningf(format string, args ... interface{}) {
	v.WARN.Printf(format, args...)
}

func (v Logger) Warningln(args ... interface{}) {
	v.WARN.Println(args...)
}

func (v Logger) Debug(args ... interface{}) {
	v.DEBUG.Print(args...)
}

func (v Logger) Debugf(format string, args ... interface{}) {
	v.DEBUG.Printf(format, args...)
}

func (v Logger) Debugln(args ... interface{}) {
	v.DEBUG.Println(args...)
}

func (v Logger) Info(args ... interface{}) {
	v.INFO.Print(args...)
}

func (v Logger) Infof(format string, args ... interface{}) {
	v.INFO.Printf(format, args...)
}

func (v Logger) Infoln(args ... interface{}) {
	v.INFO.Println(args...)
}

func (v Logger) Print(args ... interface{}) {
	v.NONE.Print(args...)
}

func (v Logger) Printf(format string, args ... interface{}) {
	v.NONE.Printf(format, args...)
}

func (v Logger) Println(args ... interface{}) {
	v.NONE.Println(args...)
}


//-----------------------------------------------------

var std = New(os.Stderr, log.LstdFlags)


func Writer() LogWriter {
	return std.Writer()
}

func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func Fatalln(args ...interface{}) {
	std.Fatalln(args...)
}

// Panic is equivalent to l.Print() followed by a call to panic().
func Panic(args ...interface{}) {
	std.Panic(args...)
}

// Panicf is equivalent to l.Printf() followed by a call to panic().
func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

// Panicln is equivalent to l.Println() followed by a call to panic().
func Panicln(args ...interface{}) {
	std.Panicln(args...)
}

// Flags returns the output flags for the logger.
// The flag bits are Ldate, Ltime, and so on.
func Flags() int {
	return std.Flags()
}

// SetFlags sets the output flags for the logger.
// The flag bits are Ldate, Ltime, and so on.
func SetFlags(flag int) {
	std.SetFlags(flag)
}

func Error(args ... interface{}) {
	std.Error(args...)
}

func Errorf(format string, args ... interface{}) {
	std.Errorf(format, args...)
}

func Errorln(args ... interface{}) {
	std.Errorln(args...)
}

func Warning(args ... interface{}) {
	std.Warning(args...)
}

func Warningf(format string, args ... interface{}) {
	std.Warningf(format, args...)
}

func Warningln(args ... interface{}) {
	std.Warningln(args...)
}

func Debug(args ... interface{}) {
	std.Debug(args...)
}

func Debugf(format string, args ... interface{}) {
	std.Debugf(format, args...)
}

func Debugln(args ... interface{}) {
	std.Debugln(args...)
}

func Info(args ... interface{}) {
	std.Info(args...)
}

func Infof(format string, args ... interface{}) {
	std.Infof(format, args...)
}

func Infoln(args ... interface{}) {
	std.Infoln(args...)
}

func Print(args ... interface{}) {
	std.Print(args...)
}

func Printf(format string, args ... interface{}) {
	std.Printf(format, args...)
}

func Println(args ... interface{}) {
	std.Println(args...)
}
