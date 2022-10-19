package log

import (
	"log"
	"time"
)

func logPrefix(level string, v []interface{}) []interface{} {
	t := time.Now().Format("2006-01-02 15:04:05")
	arr := []interface{}{t, level}
	arr = append(arr, v...)
	return arr
}

func Debug(v ...interface{}) {
	log.Println(logPrefix("DEBUG", v))
}

func Info(v ...interface{}) {
	log.Println(logPrefix("INFO", v))
}

func Warn(v ...interface{}) {
	log.Println(logPrefix("WARN", v))
}

func Error(v ...interface{}) {
	log.Println(logPrefix("ERROR", v))
}
