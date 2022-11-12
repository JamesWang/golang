package jokebox

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
