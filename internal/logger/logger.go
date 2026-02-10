package logger

import (
	"os"
	"log/slog"
)

type Options struct {   
	Level     slog.Leveler  
	AddSource bool
}

type Logger struct {
	log *slog.Logger
}

func New(opts Options) *Logger {

	hopts := &slog.HandlerOptions{
		Level:     opts.Level,
		AddSource: opts.AddSource,
	}

	h := slog.NewTextHandler(os.Stdout, hopts)
	
	return &Logger{log: slog.New(h)}
}

func (l *Logger) Debug(msg string, args ...any) {
	l.log.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.log.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.log.Warn(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.log.Error(msg, args...)
}

