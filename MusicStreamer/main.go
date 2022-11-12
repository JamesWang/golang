package main

import (
	. "MusicStreamer/jokebox"
)

func main() {
	//fmt.Printf("%v", MusicTracks.Tracks)
	go StartServer(HOST, PORT)
	StartHttpServer("0.0.0.0:8088")
}
