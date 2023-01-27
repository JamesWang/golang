package app

import (
	"log"
)

const Version = "1.0.0"

type Config struct {
	Port int
	Env  string
}

type Application struct {
	Config Config
	Logger *log.Logger
}
