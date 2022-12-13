package log

import (
	"fmt"
)

func LogOutput(message string) {
	fmt.Println(message)
}

type Logger interface {
	Log(message string)
}
