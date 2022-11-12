package jokebox

import (
	. "MusicStreamer/tracks"
	"fmt"
	"log"
	"time"
)

var jokeBoxData JokeBoxData = JokeBoxData{
	currentState:   Paused,
	streamers:      map[uint32]KeyChannelPair{},
	playList:       nil,
	positionInFile: 0,
}

func RegisterListener(key uint32, musicChan chan []byte) {
	log.Printf("Listener[%d] is joining", key)
	jokeBoxData.streamers[key] = KeyChannelPair{id: key, channel: musicChan}
}
func UnRegisterListener(key uint32) {
	log.Printf("Listener[%d] is leaving", key)
	delete(jokeBoxData.streamers, key)
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
