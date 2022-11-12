package jokebox

import "sync/atomic"

type JokeBoxState int

const (
	Playing JokeBoxState = iota
	Paused
)

type MusicCommand string

const (
	List     MusicCommand = "/list"
	Play     MusicCommand = "/play"
	Pause    MusicCommand = "/pause"
	Schedule MusicCommand = "/schedule"
)

type MusicCommandInfo struct {
	Command MusicCommand
	Data    []string
}

const (
	HOST = "0.0.0.0"
	PORT = "10255"
	TYPE = "tcp"
)

type count32 uint32

func (c *count32) Inc() uint32 {
	return atomic.AddUint32((*uint32)(c), 1)
}

func (c *count32) GetAndInc() uint32 {
	currValue := atomic.LoadUint32((*uint32)(c))
	c.Inc()
	return currValue
}
