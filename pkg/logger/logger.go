package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
)

type Logger interface {
	Debug(txt string)
	Info(txt string)
	Error(txt string)
}

type logger struct {
	lDebug *log.Logger
	lInfo  *log.Logger
	lErr   *log.Logger
}

func New(logPath string) Logger {

	logOutput := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     28,
		Compress:   true,
	}

	multiWriter := io.MultiWriter(os.Stdout, logOutput)
	// Устанавливаем все нужные флаги сразу
	lDebug := log.New(multiWriter, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	lInfo := log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	lErr := log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &logger{
		lDebug: lDebug,
		lInfo:  lInfo,
		lErr:   lErr,
	}
}

func (l *logger) Debug(txt string) {

	_ = l.lDebug.Output(2, txt)
}

func (l *logger) Info(txt string) {
	_ = l.lInfo.Output(2, txt)
}

func (l *logger) Error(txt string) {
	_ = l.lErr.Output(2, txt)
}
