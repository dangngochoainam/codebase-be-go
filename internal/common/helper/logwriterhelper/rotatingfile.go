package logwriterhelper

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

func NewRotatingFileWriter() io.Writer {
	return &lumberjack.Logger{
		Filename:   "logs/log",
		MaxSize:    50,   // megabytes
		MaxAge:     7,    // days
		MaxBackups: 7,    // files
		Compress:   true, // disabled by default
	}
}
