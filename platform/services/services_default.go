package services

import (
	"fmt"
	"platform/config"
	"platform/logging"
)

func RegisterDefaultServices() {
	err := AddSingleton(func() (cfg config.Configuration) {
		cfg, loadErr := config.Load("C:\\dev\\github\\golang\\platform\\config\\config.json")
		if loadErr != nil {
			fmt.Print("here failed")
			panic(loadErr)
		}
		return
	})
	err = AddSingleton(func(appConfig config.Configuration) logging.Logger {
		return logging.NewDefaultLogger(appConfig)
	})
	if err != nil {
		panic(err)
	}
}
