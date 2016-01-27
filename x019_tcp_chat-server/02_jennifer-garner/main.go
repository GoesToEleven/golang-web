package main

import (
	"bufio"
	"log"
	"net"
	"fmt"
)

/*
Create a chat room server. 
A client can connect and send messages to the server. 
Those messages will be broadcast to any other currently connected clients.
*/
func send (chat chan string, conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		chat <- ln
	}
}

func server (chat chan string, conns chan net.Conn) {
	var clients []net.Conn

	for {
		select {
		case msg := <-chat:
			for _, c := range clients {
				fmt.Fprintf (c, "%s\n\r", msg) // the \r fixes the weird spacing
			}
		case c := <-conns:
			clients = append(clients, c)
		}
	}
}

func main () {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	chat := make(chan string)
	conns := make(chan net.Conn)
	go server (chat, conns)

	for {
		c, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		conns <- c
		go send(chat, c)
	}
}