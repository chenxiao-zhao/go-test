package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var mapUsers map[string]int

func main() {
	var listener net.Listener
	var error error
	var conn net.Conn
	mapUsers = make(map[string]int)
	fmt.Println("Starting the serve ...")
	listener, error = net.Listen("tcp", "localhost:50002")
	checkError(error)
	for {
		conn, error = listener.Accept()
		checkError(error)
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	var buf []byte
	var error error
	for {
		buf = make([]byte, 512)
		_, error = conn.Read(buf)
		checkError(error)
		input := string(buf)
		if strings.Contains(input, ": SH") {
			fmt.Println("Server shutting down.")
			os.Exit(0)
		}
		if strings.Contains(input, ":WHO") {
			DisplayList()
		}
		ix := strings.Index(input, "says")
		clName := input[0 : ix-1]
		mapUsers[string(clName)] = 1
		fmt.Printf("Received data: --%v--", string(buf))
	}
}

func checkError(error error) {
	if error != nil {
		panic("Error: " + error.Error())
	}
}

func DisplayList() {
	fmt.Println("------------------------------------")
	fmt.Println("This is the client list: 1=active, 0=inactive")
	for key, value := range mapUsers {
		fmt.Printf("User %s is %d\n", key, value)
	}
	fmt.Println("-------------------------------------")
}
