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