package main

import (
	"fmt"
	command "goredis/commands"
	parser "goredis/parser"
	"goredis/server"
	"net"
	"strings"
)

func decodeRespString(reqStr string) string {
	if strings.HasPrefix(reqStr, "+") {
		return parser.DecodeSimpleString(reqStr)
	} else if strings.HasPrefix(reqStr, "*") {
		var commandArgList []string = parser.DecodeArrayString(reqStr)
		if commandArgList[0] == "GET" || commandArgList[0] == "get" {
			return command.Get(commandArgList)
		} else if commandArgList[0] == "SET" || commandArgList[0] == "set" {
			return command.Set(commandArgList)
		} else if commandArgList[0] == "PING" {
			return "PONG"
		} else if commandArgList[0] == "ECHO" {
			return commandArgList[1]
		} else if commandArgList[0] == "EXISTS" {
			return command.Exists(commandArgList)
		} else if commandArgList[0] == "INCR" {
			return command.AtomicIncrement(commandArgList)
		} else if commandArgList[0] == "DECR" {
			return command.AtomicDecrement(commandArgList)
		} else if commandArgList[0] == "DEL" {
			return command.Delete(commandArgList)
		} else if commandArgList[0] == "LPUSH" {
			return command.LeftPush(commandArgList)
		}
	}
	return ""
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Buffer to read data from the client
	buffer := make([]byte, 1024)

	for {
		// Read data from the client
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Connection closed.")
			return
		}

		message := string(buffer[:bytesRead])
		// fmt.Println("Received message", message)
		decodedStr := decodeRespString(message)
		// Echo the message back to the client
		conn.Write([]byte(fmt.Sprintf("+%s\r\n", decodedStr)))
	}
}

func main() {

	listener := server.CreateTCPConnection()

	for {
		// Accept incoming client connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle each client connection in a separate goroutine
		go handleClient(conn)
	}
}
