package jokebox

import (
	. "MusicStreamer/tracks"
	"fmt"
	"time"
)

var jokeBoxData JokeBoxData = JokeBoxData{
	currentState:   Paused,
	streamers:      nil,
	playList:       nil,
	positionInFile: 0,
}

func RegisterListener(musicChan chan []byte) {
	jokeBoxData.streamers = append(jokeBoxData.streamers, musicChan)
}

func HandleCommands() {
	for {
		select {
		case commandInfo := <-commandChannel:
			fmt.Printf("handling commands: %v", commandInfo)
			switch commandInfo.Command {
			case List:
				listMusics()
			case Play:
				jokeBoxData.currentState = Playing
			case Pause:
				jokeBoxData.currentState = Paused
			case Schedule:
				jokeBoxData.playList = append(jokeBoxData.playList, commandInfo.Data...)
			}
		}
	}
}

func listMusics() {
	fmt.Printf("list music:%v", MusicTracks.Tracks)
	cmdRespChannel <- MusicTracks.Tracks
}

func init() {
	go func() {
		for _ = range time.Tick(200 * time.Millisecond) {
			jokeBoxData.StreamMusicChunk()
		}
	}()
}
