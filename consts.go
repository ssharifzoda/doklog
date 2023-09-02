package doklog

const (
	logsDir  = "logs"
	logsFile = "all.log"
)

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)
