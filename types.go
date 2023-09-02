package doklog

import "os"

type LogLevel int

type LogEntry struct {
	Time    string      `json:"time"`
	Level   string      `json:"level"`
	File    string      `json:"file"`
	Message interface{} `json:"message"`
}

type Logger struct {
	file *os.File
}
