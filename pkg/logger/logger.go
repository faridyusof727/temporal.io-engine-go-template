package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type DefaultFields struct {
	Program string
	Team    string
	ENV     string
}

type Options struct {
	DefaultFields *DefaultFields
	LogPath       string
	Level         int
}

type LoggerImpl struct {
	opts  Options
	entry *logrus.Entry
}

// Entry implements Logger.
func (l *LoggerImpl) Entry() *logrus.Entry {
	return l.entry
}

// Debug implements Logger.
func (l *LoggerImpl) Debug(args ...any) {
	l.entry.Debug(args...)
}

// DebugF implements Logger.
func (l *LoggerImpl) DebugF(format string, args ...any) {
	l.entry.Debugf(format, args...)
}

// Error implements Logger.
func (l *LoggerImpl) Error(args ...any) {
	l.entry.Error(args...)
}

// ErrorF implements Logger.
func (l *LoggerImpl) ErrorF(format string, args ...any) {
	l.entry.Errorf(format, args...)
}

// Info implements Logger.
func (l *LoggerImpl) Info(args ...any) {
	l.entry.Info(args...)
}

// InfoF implements Logger.
func (l *LoggerImpl) InfoF(format string, args ...any) {
	l.entry.Infof(format, args...)
}

// Warn implements Logger.
func (l *LoggerImpl) Warn(args ...any) {
	l.entry.Warn(args...)
}

// WarnF implements Logger.
func (l *LoggerImpl) WarnF(format string, args ...any) {
	l.entry.Warnf(format, args...)
}

func New(opts Options) (Logger, error) {
	l := logrus.New()

	l.SetLevel(logrus.Level(opts.Level))

	l.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg:   "MESSAGE",
			logrus.FieldKeyLevel: "LEVEL",
		},
	})

	entry := l.WithFields(logrus.Fields{})

	if opts.DefaultFields != nil {
		entry = entry.WithFields(logrus.Fields{
			"PROGRAM": opts.DefaultFields.Program,
			"TEAM":    opts.DefaultFields.Team,
			"ENV":     opts.DefaultFields.ENV,
		})
	}

	if len(opts.LogPath) > 0 {
		output, err := os.OpenFile(opts.LogPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			return nil, fmt.Errorf("cannot setup logger output: %w", err)
		}
		l.SetOutput(output)
	}

	return &LoggerImpl{
		opts:  opts,
		entry: entry,
	}, nil
}
