package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	LogNone  = 0
	LogDefault = 1
	LogTrace  = 1 << 1
	LogDebug = 1 << 2
	LogInfo  = 1 << 3
	LogWarn  = 1 << 4
	LogError = 1 << 5
	LogFatal = 1 << 6
	LogPanic = 1 << 7
)

type LogWriter interface {
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Warningln(args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
	Traceln(args ...interface{})
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
}

type Logger struct {
	logLevel   int
	PANIC *log.Logger
	FATAL *log.Logger
	ERROR *log.Logger
	WARN  *log.Logger
	DEBUG *log.Logger
	TRACE  *log.Logger
	INFO  *log.Logger
	DEFAULT  *log.Logger
}

var logLevel int = LogDefault  | LogInfo  | LogWarn | LogError | LogFatal | LogPanic

func New(out io.Writer, logLevel, flag int) *Logger {

	if out == nil {
		out = os.Stdout
	}

	l := new(Logger)
	l.logLevel = logLevel
	l.PANIC = log.New(out, "[PANIC] ", flag)
	l.FATAL = log.New(out, "[FATAL] ", flag)
	l.ERROR = log.New(out, "[ERROR] ", flag)
	l.WARN = log.New(out, "[WARN]  ", flag)
	l.DEBUG = log.New(out, "[DEBUG] ", flag)
	l.TRACE = log.New(out, "[TRACE] ", flag)
	l.INFO = log.New(out, "[INFO] ", flag)
	l.DEFAULT = log.New(out, "", flag)

	return l
}

func GetDefaultLogLevel() int {
	return logLevel
}

func SetDefaultLogLevel(loglevel int) {
	logLevel = loglevel | LogFatal | LogPanic
}

func SetDefaultLogLevels(nornal, info, trace, warn, err, fatal, panic bool) {
	logLevel = 0
	if nornal {
		logLevel |= LogDefault
	}

	if info {
		logLevel |= LogInfo
	}

	if trace {
		logLevel |= LogTrace
	}

	if warn {
		logLevel |= LogWarn
	}

	if err {
		logLevel |= LogError
	}

	if fatal {
		logLevel |= LogFatal
	}

	if panic {
		logLevel |= LogPanic
	}
}

func (v *Logger) Writer() LogWriter {
	return LogWriter(v)
}

func (v *Logger) SetOutput(w io.Writer) {
	v.PANIC.SetOutput(w)
	v.FATAL.SetOutput(w)
	v.ERROR.SetOutput(w)
	v.WARN.SetOutput(w)
	v.DEBUG.SetOutput(w)
	v.INFO.SetOutput(w)
	v.DEFAULT.SetOutput(w)
}
/*
// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func (v Logger) Fatal(args ...interface{}) {
	if LogLevel&LogFatal != LogFatal {
		return
	}
	v.FATAL.Fatal(args...)
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func (v Logger) Fatalf(format string, args ...interface{}) {
	if LogLevel&LogFatal != LogFatal {
		return
	}
	v.FATAL.Fatalf(format, args...)
}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func (v Logger) Fatalln(args ...interface{}) {
	if LogLevel&LogFatal != LogFatal {
		return
	}
	v.FATAL.Fatalln(args...)
}

// Panic is equivalent to l.Print() followed by a call to panic().
func (v Logger) Panic(args ...interface{}) {
	if LogLevel&LogPanic != LogPanic {
		return
	}
	v.PANIC.Panic(args...)
}

// Panicf is equivalent to l.Printf() followed by a call to panic().
func (v Logger) Panicf(format string, args ...interface{}) {
	if LogLevel&LogPanic != LogPanic {
		return
	}
	v.PANIC.Panicf(format, args...)
}

// Panicln is equivalent to l.Println() followed by a call to panic().
func (v Logger) Panicln(args ...interface{}) {
	if LogLevel&LogPanic != LogPanic {
		return
	}
	v.PANIC.Panicln(args...)
}
*/

func (v *Logger) GetLogLevel() int {
	return v.logLevel
}

func (v *Logger) SetLogLevel(loglevel int) {
	v.logLevel = loglevel | LogFatal | LogPanic
}

func (v *Logger) SetLogLevels(nornal, info, trace, warn, err, fatal, panic bool) {
	v.logLevel = 0
	if nornal {
		v.logLevel |= LogDefault
	}

	if info {
		v.logLevel |= LogInfo
	}

	if trace {
		v.logLevel |= LogTrace
	}

	if warn {
		v.logLevel |= LogWarn
	}

	if err {
		v.logLevel |= LogError
	}

	if fatal {
		v.logLevel |= LogFatal
	}

	if panic {
		v.logLevel |= LogPanic
	}
}

// Flags returns the output flags for the logger.
// The flag bits are Ldate, Ltime, and so on.
func (v *Logger) Flags() int {
	if v.logLevel&LogDefault != LogDefault {
		return 0
	}
	return v.DEFAULT.Flags()
}

// SetFlags sets the output flags for the logger.
// The flag bits are Ldate, Ltime, and so on.
func (v *Logger) SetFlags(flag int) {
	v.PANIC.SetFlags(flag)
	v.FATAL.SetFlags(flag)
	v.ERROR.SetFlags(flag)
	v.WARN.SetFlags(flag)
	v.DEBUG.SetFlags(flag)
	v.TRACE.SetFlags(flag)
	v.INFO.SetFlags(flag)
	v.DEFAULT.SetFlags(flag)
}

func (v *Logger) Panic(args ...interface{}) {
	s := fmt.Sprint(args...)
	v.PanicOutput(3, s)
}

func (v *Logger) Panicf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	v.PanicOutput(3, s)
}

func (v *Logger) Panicln(args ...interface{}) {
	s := fmt.Sprintln(args...)
	v.PanicOutput(3, s)
}

func (v *Logger) PanicOutput(calldepth int, s string) {
	if v.logLevel&LogPanic != LogPanic {
		return
	}
	v.PANIC.Output(calldepth, s)
	panic(s)
}

func (v *Logger) Fatal(args ...interface{}) {
	s := fmt.Sprint(args...)
	v.FatalOutput(3, s)
}

func (v *Logger) Fatalf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	v.FatalOutput(3, s)
}

func (v *Logger) Fatalln(args ...interface{}) {
	s := fmt.Sprintln(args...)
	v.FatalOutput(3, s)
}

func (v *Logger) FatalOutput(calldepth int, s string) {
	if v.logLevel&LogFatal != LogFatal {
		return
	}
	v.FATAL.Output(calldepth, s)
	os.Exit(1)
}

func (v *Logger) Error(args ...interface{}) {
	s := fmt.Sprint(args...)
	v.ErrorOutput(3, s)
}

func (v *Logger) Errorf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	v.ErrorOutput(3, s)
}

func (v *Logger) Errorln(args ...interface{}) {
	s := fmt.Sprintln(args...)
	v.ErrorOutput(3, s)
}

func (v *Logger) ErrorOutput(calldepth int, s string) {
	if v.logLevel&LogError != LogError {
		return
	}
	v.ERROR.Output(calldepth, s)
}

func (v *Logger) Warning(args ...interface{}) {
	s := fmt.Sprint(args...)
	v.WarningOutput(3, s)
}

func (v *Logger) Warningf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	v.WarningOutput(3, s)
}

func (v *Logger) Warningln(args ...interface{}) {
	s := fmt.Sprintln(args...)
	v.WarningOutput(3, s)
}

func (v *Logger) WarningOutput(calldepth int, s string) {
	if v.logLevel&LogWarn != LogWarn {
		return
	}
	v.WARN.Output(calldepth, s)
}

func (v *Logger) Debug(args ...interface{}) {
	s := fmt.Sprint(args...)
	v.DebugOutput(3, s)
}

func (v *Logger) Debugf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	v.DebugOutput(3, s)
}

func (v *Logger) Debugln(args ...interface{}) {
	s := fmt.Sprintln(args...)
	v.DebugOutput(3, s)
}

func (v *Logger) DebugOutput(calldepth int, s string) {
	if v.logLevel&LogDebug != LogDebug {
		return
	}
	v.DEBUG.Output(calldepth, s)
}

func (v *Logger) Trace(args ...interface{}) {
	s := fmt.Sprint(args...)
	v.TraceOutput(3, s)
}

func (v *Logger) Tracef(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	v.TraceOutput(3, s)
}

func (v *Logger) Traceln(args ...interface{}) {
	s := fmt.Sprintln(args...)
	v.TraceOutput(3, s)
}

func (v *Logger) TraceOutput(calldepth int, s string) {
	if v.logLevel&LogTrace != LogTrace {
		return
	}
	v.TRACE.Output(calldepth, s)
}

func (v *Logger) Info(args ...interface{}) {
	s := fmt.Sprint(args...)
	v.InfoOutput(3, s)
}

func (v Logger) Infof(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	v.InfoOutput(3, s)
}

func (v *Logger) Infoln(args ...interface{}) {
	s := fmt.Sprintln(args...)
	v.InfoOutput(3, s)
}

func (v *Logger) InfoOutput(calldepth int, s string) {
	if v.logLevel&LogInfo != LogInfo {
		return
	}
	v.INFO.Output(calldepth, s)
}

func (v *Logger) Print(args ...interface{}) {
	s := fmt.Sprint(args...)
	v.DefaultOutput(3, s)
}

func (v Logger) Printf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	v.DefaultOutput(3, s)
}

func (v Logger) Println(args ...interface{}) {
	s := fmt.Sprintln(args...)
	v.DefaultOutput(3, s)
}

func (v Logger) DefaultOutput(calldepth int, s string) {
	if v.logLevel&LogDefault != LogDefault {
		return
	}
	v.DEFAULT.Output(calldepth, s)
}

//-----------------------------------------------------

var std = New(os.Stderr, logLevel, log.LstdFlags|log.Lshortfile)

func Writer() LogWriter {
	return std.Writer()
}

func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

func GetLogLevel() int {
	return std.GetLogLevel()
}

func SetLogLevel(loglevel int) {
	std.SetLogLevel(loglevel | LogFatal | LogPanic)
}

func SetLogLevels(nornal, info, trace, warn, err, fatal, panic bool) {
	std.SetLogLevels(nornal, info, trace, warn, err, fatal, panic)
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

// Panic is equivalent to l.Print() followed by a call to panic().
func Panic(args ...interface{}) {
	std.PanicOutput(3, fmt.Sprint(args...))
}

// Panicf is equivalent to l.Printf() followed by a call to panic().
func Panicf(format string, args ...interface{}) {
	std.PanicOutput(3, fmt.Sprintf(format, args...))
}

// Panicln is equivalent to l.Println() followed by a call to panic().
func Panicln(args ...interface{}) {
	std.PanicOutput(3, fmt.Sprintln(args...))
}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func Fatal(args ...interface{}) {
	std.FatalOutput(3, fmt.Sprint(args...))
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func Fatalf(format string, args ...interface{}) {
	std.FatalOutput(3, fmt.Sprintf(format, args...))
}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func Fatalln(args ...interface{}) {
	std.FatalOutput(3, fmt.Sprintln(args...))
}

func Error(args ...interface{}) {
	std.ErrorOutput(3, fmt.Sprint(args...))

}

func Errorf(format string, args ...interface{}) {
	std.ErrorOutput(3, fmt.Sprintf(format, args...))

}

func Errorln(args ...interface{}) {
	std.ErrorOutput(3, fmt.Sprintln(args...))

}

func Warning(args ...interface{}) {
	std.WarningOutput(3, fmt.Sprint(args...))
}

func Warningf(format string, args ...interface{}) {
	std.WarningOutput(3, fmt.Sprintf(format, args...))

}

func Warningln(args ...interface{}) {
	std.WarningOutput(3, fmt.Sprintln(args...))

}

func Debug(args ...interface{}) {
	std.DebugOutput(3, fmt.Sprint(args...))
}

func Debugf(format string, args ...interface{}) {
	std.DebugOutput(3, fmt.Sprintf(format, args...))

}

func Debugln(args ...interface{}) {
	std.DebugOutput(3, fmt.Sprintln(args...))
}

func Trace(args ...interface{}) {
	std.TraceOutput(3, fmt.Sprint(args...))
}

func Tracef(format string, args ...interface{}) {
	std.TraceOutput(3, fmt.Sprintf(format, args...))

}

func Traceln(args ...interface{}) {
	std.TraceOutput(3, fmt.Sprintln(args...))
}

func Info(args ...interface{}) {
	std.InfoOutput(3, fmt.Sprint(args...))
}

func Infof(format string, args ...interface{}) {
	std.InfoOutput(3, fmt.Sprintf(format, args...))
}

func Infoln(args ...interface{}) {
	std.InfoOutput(3, fmt.Sprintln(args...))
}

func Print(args ...interface{}) {
	std.DefaultOutput(3, fmt.Sprint(args...))
}

func Printf(format string, args ...interface{}) {
	std.DefaultOutput(3, fmt.Sprintf(format, args...))
}

func Println(args ...interface{}) {
	std.DefaultOutput(3, fmt.Sprintln(args...))
}

func GetPanic() *log.Logger {
	return std.PANIC
}

func GetFatal() *log.Logger {
	return std.FATAL
}

func GetError() *log.Logger {
	return std.ERROR
}

func GetWarn() *log.Logger {
	return std.WARN
}

func GetDebug() *log.Logger {
	return std.DEBUG
}

func GetTrace() *log.Logger {
	return std.TRACE
}

func GetInfo() *log.Logger {
	return std.INFO
}

func GetDefault() *log.Logger {
	return std.DEFAULT
}
