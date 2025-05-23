package logger

import (
	"fmt"
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
	l *log.Logger
}

func New(logPath string) Logger {
	err := os.MkdirAll("app", os.ModePerm)
	if err != nil {
		fmt.Println("Ошибка при создании папки:", err)
	} else {
		fmt.Println("Папка 'app' успешно создана.")
	}

	logOutput := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     28,
		Compress:   true,
	}

	multiWriter := io.MultiWriter(os.Stdout, logOutput)
	// Устанавливаем все нужные флаги сразу
	l := log.New(multiWriter, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &logger{l: l}
}

func (l *logger) Debug(txt string) {
	// depth = 3 — чтобы выйти из Debug → output → l.l.Output → в вызывающий код
	_ = l.l.Output(2, txt)
}

func (l *logger) Info(txt string) {
	_ = l.l.Output(2, txt)
}

func (l *logger) Error(txt string) {
	_ = l.l.Output(2, txt)
}
