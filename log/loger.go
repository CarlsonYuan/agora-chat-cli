/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package log

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	logger  zerolog.Logger
	verbose bool
	writer  *zerolog.ConsoleWriter
}

func (l *Logger) SetFieldsOrder(fieldsOrder []string) {
	l.writer.FieldsOrder = fieldsOrder
	l.logger = log.Output(*l.writer)
}

// NewLogger initializes and returns a new Logger instance
func NewLogger(verbose bool) *Logger {
	zerolog.TimestampFunc = time.Now().UTC
	writer := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logger := log.Output(writer)
	return &Logger{
		logger:  logger,
		verbose: verbose,
		writer:  &writer,
	}
}

// Info logs an informational message
func (l *Logger) Info(message string, fields map[string]interface{}) {
	event := l.logger.Info()
	for k, v := range fields {
		event.Interface(k, v)
	}
	event.Msg(message)
}

// Error logs an error message
func (l *Logger) Error(message string, fields map[string]interface{}) {
	event := l.logger.Error()
	for k, v := range fields {
		event.Interface(k, v)
	}
	event.Msg(message)
}

func (l *Logger) Warn(message string, fields map[string]interface{}) {
	event := l.logger.Warn()
	for k, v := range fields {
		event.Interface(k, v)
	}
	event.Msg(message)
}

// Debug logs a debug message
func (l *Logger) Debug(message string, fields map[string]interface{}) {
	event := l.logger.Debug()
	for k, v := range fields {
		event.Interface(k, v)
	}
	event.Msg(message)
}

// Verbose logs a message only if verbose mode is enabled
func (l *Logger) Verbose(message string, fields map[string]interface{}) {
	if l.verbose {
		event := l.logger.Debug() // You might want to log verbose messages at the debug level
		for k, v := range fields {
			event.Interface(k, v)
		}
		event.Msg(message)
	}
}

// Fatal logs a fatal message and exits the program
func (l *Logger) Fatal(message string, fields map[string]interface{}) {
	event := l.logger.Fatal()
	for k, v := range fields {
		event.Interface(k, v)
	}
	event.Msg(message)
	os.Exit(1) // Exit with status code 1 to indicate an error
}
