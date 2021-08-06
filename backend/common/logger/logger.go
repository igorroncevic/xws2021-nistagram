package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	stdoutLog *logrus.Logger
	fileLog   *logrus.Logger
}

type LogLevel string

const (
	Debug LogLevel = "Debug"
	Info           = "Info"
	Warn           = "Warn"
	Error          = "Error"
	Fatal          = "Fatal"
)

func NewLogger() *Logger {
	StdoutFormatter := new(logrus.TextFormatter)
	StdoutFormatter.FullTimestamp = true
	StdoutFormatter.ForceColors = true

	// Logging to standard output
	stdoutLog := logrus.New()
	stdoutLog.SetOutput(os.Stdout)
	stdoutLog.SetFormatter(StdoutFormatter)
	stdoutLog.SetLevel(logrus.DebugLevel)

	// Logging to file
	fileLog := logrus.New()
	_, err := os.OpenFile("./server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0700) // R/W only for owner
	if err != nil {
		stdoutLog.Error("error opening file: %v", err)
		os.Exit(1)
	}
	fileLog.SetOutput(&lumberjack.Logger{
		Filename:   "./server.log",
		MaxSize:    10, // 10mb before deletion
		MaxAge:     30, // 30 days before deletion
		MaxBackups: 2,
		Compress:   true,
	})

	return &Logger{
		stdoutLog,
		fileLog,
	}
}

func (l *Logger) ToStdoutAndFile(function string, msg string, level LogLevel) {
	l.ToStdout(function, msg, level)
	l.ToFile(function, msg, level)
}

func (l *Logger) ToStdout(function string, msg string, level LogLevel) {
	messageToLog := function + ": " + msg
	switch level {
	case Debug:
		l.stdoutLog.Debug(messageToLog)
		break
	case Info:
		l.stdoutLog.Info(messageToLog)
		break
	case Warn:
		l.stdoutLog.Warn(messageToLog)
		break
	case Error:
		l.stdoutLog.Error(messageToLog)
		break
	case Fatal:
		l.stdoutLog.Fatal(messageToLog)
		break
	default:
		l.stdoutLog.Info(messageToLog)
	}
}

func (l *Logger) ToFile(function string, msg string, level LogLevel) {
	messageToLog := function + ": " + msg
	switch level {
	case Debug:
		l.fileLog.Debug(messageToLog)
		break
	case Info:
		l.fileLog.Info(messageToLog)
		break
	case Warn:
		l.fileLog.Warn(messageToLog)
		break
	case Error:
		l.fileLog.Error(messageToLog)
		break
	case Fatal:
		l.fileLog.Fatal(messageToLog)
		break
	default:
		l.fileLog.Info(messageToLog)
	}
}
