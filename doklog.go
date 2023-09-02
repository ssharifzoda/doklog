package doklog

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	logger *Logger
)

func NewLogger(file *os.File) *Logger {
	return &Logger{
		file: file,
	}
}

func init() {
	logsPath := filepath.Join(logsDir, logsFile)
	file, err := os.OpenFile(logsPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	logger = NewLogger(file)
}

func (l *Logger) Log(level LogLevel, args ...interface{}) {
	logTime := time.Now().Format("2006-01-02 15:04:05")
	logLevel := GetLogLevel(level)
	_, file, line, _ := runtime.Caller(2)
	filName := filepath.Base(file)
	logEntry := &LogEntry{
		Time:    logTime,
		Level:   logLevel,
		File:    fmt.Sprintf("%s:%d", filName, line),
		Message: args,
	}
	logJson, err := json.Marshal(logEntry)
	if err != nil {
		return
	}
	_, err = l.file.Write(append(logJson, '\n'))
	if err != nil {
		fmt.Println("Log not working")
	}
}

func Debug(args ...interface{}) {
	logger.Log(DebugLevel, args...)
}

func Info(args ...interface{}) {
	logger.Log(InfoLevel, args...)
}

func Warning(args ...interface{}) {
	logger.Log(WarningLevel, args...)
}
func Error(args ...interface{}) {
	logger.Log(ErrorLevel, args...)
}

func Fatal(args ...interface{}) {
	logger.Log(FatalLevel, args...)
}
