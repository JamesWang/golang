package jokebox

import (
	. "MusicStreamer/tracks"
	"fmt"
)

func HandleCommands() {
	for {
		select {
		case commandInfo := <-commandChannel:
			fmt.Printf("handling commands: %v", commandInfo)
			switch commandInfo.Command {
			case List:
				listMusics()
			}
		}
	}
}

func listMusics() {
	fmt.Printf("list music:%v", MusicTracks.Tracks)
	cmdRespChannel <- MusicTracks.Tracks
}
