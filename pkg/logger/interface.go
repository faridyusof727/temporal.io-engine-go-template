package logger

type Logger interface {
	Info(args ...any)
	InfoF(format string, args ...any)
	Error(args ...any)
	ErrorF(format string, args ...any)
	Debug(args ...any)
	DebugF(format string, args ...any)
	Warn(args ...any)
	WarnF(format string, args ...any)
}

type TemporalLogger interface {
	Debug(msg string, keyvals ...interface{})
	Info(msg string, keyvals ...interface{})
	Warn(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})
}
