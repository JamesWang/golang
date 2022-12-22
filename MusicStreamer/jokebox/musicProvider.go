package jokebox

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

//var chunkChannel = make(chan []byte)

type CurrentFile struct {
	file   os.File
	reader bufio.Reader
	done   bool
}

func (cf CurrentFile) Base() string {
	return filepath.Base(cf.file.Name())
}

type KeyChannelPair struct {
	id      uint32
	channel chan []byte
}

type JokeBoxData struct {
	musicLocation    string
	currentState     JokeBoxState
	streamers        map[uint32]KeyChannelPair
	playList         []string
	positionInFile   int64
	currentMusicFile CurrentFile
}

func (jbd *JokeBoxData) finished() {
	jbd.currentMusicFile.file.Close()
	jbd.currentMusicFile.file = os.File{}
	jbd.positionInFile = 0
	jbd.currentMusicFile.done = true
	jbd.playList = jbd.playList[1:]
}

func (jbd *JokeBoxData) StreamMusicChunk() {
	if jbd.currentState == Paused && !jbd.currentMusicFile.done {
		return
	}
	if jbd.currentMusicFile.done && len(jbd.playList) == 0 {
		jbd.currentState = Paused
		return
	}

	if jbd.currentMusicFile.done || jbd.currentMusicFile.file == (os.File{}) {
		jbd.playNextMusic()
	}
	buffer := make([]byte, 0, 4096*4)
	count, err := jbd.currentMusicFile.reader.Read(buffer[:cap(buffer)])
	buffer = buffer[:count]
	if err != nil {
		if err == io.EOF {
			fmt.Printf("%s is done", jbd.currentMusicFile.Base())
			jbd.finished()
		} else {
			log.Fatalf("read music failed: %v", err)
		}
	} else {
		if count > 0 {
			jbd.Broadcasting(&buffer)
		}
	}
}

func (jbd *JokeBoxData) playNextMusic() {
	currFile := jbd.playList[0]
	file, err := os.Open(currFile)
	if err != nil {
		log.Fatalf("failed to open music: %s", currFile)
	}
	jbd.currentMusicFile = CurrentFile{
		file:   *file,
		reader: *bufio.NewReader(file),
		done:   false,
	}
}

func (jbd *JokeBoxData) Broadcasting(buffer *[]byte) {
	jbd.positionInFile += int64(len(*buffer))
	if len(*buffer) == 0 {
		jbd.finished()
		return
	}
	for _, keyChannel := range jbd.streamers {
		keyChannel.channel <- *buffer
	}
}
