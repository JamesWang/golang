package tracks

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type TrackLoader interface {
	load(location string, filter func(string) bool) []string
}

type TrackInfo struct {
	Location string
	Tracks   []string
}

func ReadDir(dirname string) []os.FileInfo {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func (ti *TrackInfo) load(location string, filter func(string) bool) []string {
	ti.Location = location
	files := ReadDir(location)

	ti.Tracks = make([]string, len(files))
	for _, file := range files {
		ti.Tracks = append(ti.Tracks, file.Name())
	}
	return ti.Tracks
}

var MusicTracks *TrackInfo

func mp3(name string) bool {
	return strings.HasSuffix(name, ".mp3")
}

func init() {
	fmt.Println("initializing....")
	MusicTracks = &TrackInfo{}
	MusicTracks.load("V:\\MusicPhotos\\music", mp3)
}
