package main

import (
	. "EasyDI/log"
	"net/http"
)

type Logic interface {
	SayHello(userId string) (string, error)
}

type Controller struct {
	log   Logger
	logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.log.Log("In SayHello")
	userId := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{
		log:   l,
		logic: logic,
	}
}
