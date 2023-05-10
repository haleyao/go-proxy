package log

import (
	"fmt"
	"log"
	"os"
)

type Level string

const (
	INFO  Level = "INFO"
	WARN  Level = "WARN"
	ERROR Level = "ERROR"
	FATAL Level = "FATAL"
)

func (l Level) String() string {
	return string(l)
}

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "proxy", log.Ldate|log.Ltime|log.Lmicroseconds)
}

func printf(l Level, f string, args ...interface{}) string {
	return fmt.Sprintf(l.String()+": "+f, args...)
}

func Infof(f string, args ...interface{}) {
	logger.Output(3, printf(INFO, f, args...))
}

func Warnf(f string, args ...interface{}) {
	logger.Output(3, printf(WARN, f, args...))
}

func Errorf(f string, args ...interface{}) {
	logger.Output(3, printf(ERROR, f, args...))
}

func Fatalf(f string, args ...interface{}) {
	l := log.New(os.Stderr, "proxy", log.Ldate|log.Ltime|log.Lmicroseconds)
	l.Output(3, printf(FATAL, f, args...))
	os.Exit(1)
}
