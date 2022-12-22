package domain

import (
	. "EasyDI/ds"
	. "EasyDI/log"
	"errors"
)

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg.Log(message)
}

func (sl SimpleLogic) SayHello(userId string) (string, error) {
	sl.l.Log("In SayHello for " + userId)
	name, ok := sl.ds.UserNameForId(userId)
	if !ok {
		return "", errors.New("unknown user")
	} else {
		return "Hello, " + name, nil
	}
}

func (sl SimpleLogic) SayGoodBye(userId string) (string, error) {
	sl.l.Log("in SayGoodBye for " + userId)
	name, ok := sl.ds.UserNameForId(userId)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye, " + name, nil
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}
