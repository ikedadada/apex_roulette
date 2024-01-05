package service

type LogLevel string

const (
	LogLevelInfo     LogLevel = "INFO"
	LogLevelError    LogLevel = "ERROR"
	LogLevelCritical LogLevel = "CRITICAL"
)

type Logger interface {
	StructLog(level LogLevel, message string)
}
