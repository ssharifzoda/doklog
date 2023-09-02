package doklog

func GetLogLevel(level LogLevel) string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case FatalLevel:
		return "FATAL"
	case ErrorLevel:
		return "ERROR"
	case InfoLevel:
		return "INFO"
	case WarningLevel:
		return "WARNING"
	default:
		return "UNKNOWN"
	}
}
