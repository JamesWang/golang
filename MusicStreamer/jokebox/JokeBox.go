package jokebox

type JokeBoxStatus int

const (
	Playing JokeBoxStatus = iota
	Paused
)

type MusicCommand string

const (
	List     MusicCommand = "/list"
	Play     MusicCommand = "/play"
	Pause    MusicCommand = "/pause"
	Schedule MusicCommand = "/schedule"
)

const (
	HOST = "0.0.0.0"
	PORT = "10255"
	TYPE = "tcp"
)
