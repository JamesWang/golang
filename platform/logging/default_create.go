package logging

import (
	"log"
	"os"
)

func NewDefaultLogger(level LogLevel) Logger {
	flags := log.Lmsgprefix | log.Ltime

	newLogger := func(prefix string) *log.Logger {
		return log.New(os.Stdout, prefix, flags)
	}

	return &DefaultLogger{
		minLevel: level,
		loggers: map[LogLevel]*log.Logger{
			Trace:       newLogger("TRACE"),
			Debug:       newLogger("DEBUG"),
			Information: newLogger("INFO"),
			Warning:     newLogger("WARN"),
			Fatal:       newLogger("FATAL"),
		},
		triggerPanic: true,
	}
}
