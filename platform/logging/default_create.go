package logging

import (
	"log"
	"os"
	"platform/config"
	"strings"
)

func NewDefaultLogger(conf config.Configuration) Logger {
	var level LogLevel = Debug
	if configLevelString, found := conf.GetString("logging:level"); found {
		level = LogLevelFromString(configLevelString)
	}

	flags := log.Lmsgprefix | log.Ltime

	newLogger := func(prefix string) *log.Logger {
		return log.New(os.Stdout, prefix, flags)
	}

	return &DefaultLogger{
		minLevel: level,
		loggers: map[LogLevel]*log.Logger{
			Trace:       newLogger("TRACE "),
			Debug:       newLogger("DEBUG "),
			Information: newLogger("INFO "),
			Warning:     newLogger("WARN "),
			Fatal:       newLogger("FATAL "),
		},
		triggerPanic: true,
	}
}

func LogLevelFromString(val string) (level LogLevel) {
	switch strings.ToLower(val) {
	case "debug":
		level = Debug
	case "information":
		level = Information
	case "warning":
		level = Warning
	case "fatal":
		level = Fatal
	case "none":
		level = None
	default:
		level = Debug
	}
	return
}
