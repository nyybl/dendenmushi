package lib

import (
	"log"
)

type Logger struct {
	ModuleName string 
}

func NewLogger(module string) *Logger {
	return &Logger{
		ModuleName: module,
	}
}

func (l *Logger) getPrefix(logType string) string {
	return " [" + logType + "/" + l.ModuleName + "] "
} 

func (l *Logger) Print(d ...any) {
	prefix := l.getPrefix("INFO")
	log.Println(prefix, d)
}

func (l *Logger) Error(d ...any) {
	prefix := l.getPrefix("ERROR")
	log.Println(prefix, d)
}