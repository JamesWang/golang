package jokebox

import (
	"encoding/binary"
	"log"
	"net/http"
)

type MusicHandler struct {
	message string
}

var connectionCount count32

func closeHandler(key uint32, writer http.ResponseWriter) {
	select {
	case closed := <-writer.(http.CloseNotifier).CloseNotify():
		if closed {
			UnRegisterListener(key)
		}
	}
}
func (sh MusicHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "audio/mpeg")
	writer.Header().Set("Connection", "Keep-Alive")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Transfer-Encoding", "chunked")
	id := connectionCount.GetAndInc()
	var musicChannel = make(chan []byte)
	go RegisterListener(id, musicChannel)
	go closeHandler(id, writer)

	for {
		select {
		case chunk := <-musicChannel:
			binary.Write(writer, binary.BigEndian, chunk)
			flusher, ok := writer.(http.Flusher)
			if ok {
				flusher.Flush()
			}
		}
	}
}

func StartHttpServer(listenOn string) {
	http.Handle("/music", MusicHandler{"will play"})
	err := http.ListenAndServe(listenOn, nil)
	if err != nil {
		log.Fatalf("Cannot start music streamer, %v", err)
	}
}
