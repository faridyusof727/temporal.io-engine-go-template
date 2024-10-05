package logger

import (
	"github.com/sirupsen/logrus"
)

type TemporalLoggerImpl struct {
	entry *logrus.Entry
}

// Debug implements TemporalLogger.
func (t *TemporalLoggerImpl) Debug(msg string, keyvals ...interface{}) {
	t.entry.Debug(msg, keyvals)
}

// Error implements TemporalLogger.
func (t *TemporalLoggerImpl) Error(msg string, keyvals ...interface{}) {
	t.entry.Error(msg, keyvals)
}

// Info implements TemporalLogger.
func (t *TemporalLoggerImpl) Info(msg string, keyvals ...interface{}) {
	t.entry.Info(msg, keyvals)
}

// Warn implements TemporalLogger.
func (t *TemporalLoggerImpl) Warn(msg string, keyvals ...interface{}) {
	t.entry.Warn(msg, keyvals)
}

func NewTemporalLoggerAdapter(logger *LoggerImpl) TemporalLogger {
	l := logger.entry

	return &TemporalLoggerImpl{
		entry: l,
	}
}
