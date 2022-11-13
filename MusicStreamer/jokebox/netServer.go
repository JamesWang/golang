package jokebox

import (
	. "MusicStreamer/tracks"
	"bufio"
	"fmt"
	"log"
	"net"
	"path/filepath"
	"strings"
)

var count = 0

func StartServer(host string, port string) {
	listener, err := net.Listen("tcp4", host+":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	go HandleCommands()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(bufio.NewReader(conn), bufio.NewWriter(conn))
		count++
	}
}

var commandChannel = make(chan MusicCommandInfo)
var cmdRespChannel = make(chan []string)

func handleConnection(reqReader *bufio.Reader, respWriter *bufio.Writer) {
	for {
		netData, err := reqReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("received: %s", netData)
		command := strings.TrimSpace(string(netData))
		if command == "/quit" {
			break
		}
		handleCommand(respWriter, command)
		respWriter.Flush()
	}
}

func handleCommand(respWriter *bufio.Writer, command string) {
	log.Printf("Handling command[%s]...", command)
	switch command {
	case "/list":
		commandChannel <- MusicCommandInfo{Command: List}
		resp := <-cmdRespChannel
		respWriter.WriteString(strings.Join(resp, "\n"))
	case "/play":
		commandChannel <- MusicCommandInfo{Command: Play}
	case "/pause":
		commandChannel <- MusicCommandInfo{Command: Pause}
	default:
		if strings.HasPrefix(command, string(Schedule)) {
			scheduleMusic(respWriter, command)
		} else {
			respWriter.WriteString("Unknown command\n")
		}
	}
}

func scheduleMusic(respWriter *bufio.Writer, command string) {
	tmpList := strings.Split(command[10:], ",")
	var toBeScheduled []string
	for _, name := range tmpList {
		toBeScheduled = append(toBeScheduled, filepath.Join(MusicTracks.Location, name))
	}
	log.Printf("music [%v] are schedule to play", toBeScheduled)
	commandChannel <- MusicCommandInfo{Command: Schedule, Data: toBeScheduled}
	respWriter.WriteString("Musics are scheduled to play!\n")
}
