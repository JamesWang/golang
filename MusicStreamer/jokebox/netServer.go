package jokebox

import (
	"bufio"
	"fmt"
	"net"
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
		csocket, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(csocket)
		count++
	}
}

var commandChannel = make(chan MusicCommandInfo)
var cmdRespChannel = make(chan []string)

func handleConnection(csocket net.Conn) {
	for {
		netData, err := bufio.NewReader(csocket).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("received: %s", netData)
		command := strings.TrimSpace(string(netData))
		if command == "/quit" {
			break
		}
		handleCommand(csocket, command)
	}
}

func handleCommand(csocket net.Conn, command string) {
	fmt.Println("inside handleCommand")
	switch command {
	case "/list":
		commandChannel <- MusicCommandInfo{Command: List}
		resp := <-cmdRespChannel
		fmt.Printf("received music: %v", resp)
		writer := bufio.NewWriter(csocket)
		writer.WriteString(strings.Join(resp, "\n"))
		writer.Flush()
	case "/play":
		commandChannel <- MusicCommandInfo{Command: Play}
	case "/pause":
		commandChannel <- MusicCommandInfo{Command: Pause}
	default:
		if strings.HasPrefix(command, string(Schedule)) {
			scheduleMusic(csocket, command)
		} else {
			writer := bufio.NewWriter(csocket)
			writer.WriteString("Unknown command\n")
			writer.Flush()
		}
	}
}

func scheduleMusic(csocket net.Conn, command string) {
	toBeScheduled := strings.Split(command[10:], ",")
	commandChannel <- MusicCommandInfo{Command: Schedule, Data: toBeScheduled}
	writer := bufio.NewWriter(csocket)
	writer.WriteString("Musics are scheduled to play!")
	writer.Flush()
}
