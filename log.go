package glog

import (
	"fmt"
	"io"
	"log"
	"os"
)

type loglevel int

const (
	// Trace is log level for tracing usage
	Trace = loglevel(3)
	// Info is log level for delevering information
	Info = loglevel(2)
	// Error is log level to say there is an error at this process but the program can keep going
	Error = loglevel(1)
	// Fatal is log level to say that there is an error and the program should stop
	Fatal = loglevel(0)
)

const (
	detailFlag = log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile
	lessFlag   = log.Ldate | log.Lmicroseconds
)

// Logger is structur for a log management
type Logger struct {
	level loglevel
	log   *log.Logger
}

// New is func to create a log management at level info
func New(out io.Writer) *Logger {
	logger := &Logger{
		level: Info,
		log:   log.New(out, "", lessFlag),
	}
	return logger
}

// SetLevel is func to set level for witch message logger should print
func (m *Logger) SetLevel(level loglevel) {
	m.level = level
}

// Fatal is func to print fatal message
func (m *Logger) Fatal(v ...interface{}) {
	if m.level >= Fatal {
		m.log.SetPrefix("[FATAL]\t")
		m.log.SetFlags(detailFlag)
		m.output(fmt.Sprint(v...))
		os.Exit(1)
	}
}

// Fatalf is func to print fatal message ini formatted string
func (m *Logger) Fatalf(format string, v ...interface{}) {
	out := fmt.Sprintf(format, v...)
	m.Fatal(out)
}

// Error is func to print error message
func (m *Logger) Error(v ...interface{}) {
	if m.level >= Error {
		m.log.SetPrefix("[ERROR]\t")
		m.log.SetFlags(detailFlag)
		m.output(fmt.Sprint(v...))
	}
}

// Errorf is func to print formatted
func (m *Logger) Errorf(format string, v ...interface{}) {
	out := fmt.Sprintf(format, v...)
	m.Error(out)
}

// Info is func to print info message
func (m *Logger) Info(v ...interface{}) {
	if m.level >= Info {
		m.log.SetPrefix("[INFO]\t")
		if m.level == Info {
			m.log.SetFlags(lessFlag)
		} else {
			m.log.SetFlags(detailFlag)
		}
		m.output(fmt.Sprint(v...))
	}
}

// Infof is func to print formatted info message
func (m *Logger) Infof(format string, v ...interface{}) {
	out := fmt.Sprintf(format, v...)
	m.Info(out)
}

// Trace is func to print tracing message
func (m *Logger) Trace(v ...interface{}) {
	if m.level >= Trace {
		m.log.SetPrefix("[TRACE]\t")
		m.log.SetFlags(detailFlag)
		m.output(fmt.Sprint(v...))
	}
}

// Tracef is func to print formmated tracing message
func (m *Logger) Tracef(format string, v ...interface{}) {
	out := fmt.Sprintf(format, v...)
	m.Trace(out)
}

func (m *Logger) output(s string) {
	m.log.Output(3, s)
}
