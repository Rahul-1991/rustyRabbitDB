package main

import (
	"fmt"
	command "goredis/commands"
	parser "goredis/parser"
	"goredis/server"
	"goredis/store"
	"goredis/utility"
	"net"
	"strings"
)

func decodeRespString(reqStr string) string {
	if strings.HasPrefix(reqStr, "+") {
		return parser.DecodeSimpleString(reqStr)
	} else if strings.HasPrefix(reqStr, "*") {
		var commandArgList []string = parser.DecodeArrayString(reqStr)
		if strings.ToLower(commandArgList[0]) == "get" {
			return command.Get(commandArgList)
		} else if strings.ToLower(commandArgList[0]) == "set" {
			return command.Set(commandArgList)
		} else if strings.ToLower(commandArgList[0]) == "ping" {
			return "PONG"
		} else if strings.ToLower(commandArgList[0]) == "echo" {
			return commandArgList[1]
		} else if strings.ToLower(commandArgList[0]) == "exists" {
			return command.Exists(commandArgList)
		} else if strings.ToLower(commandArgList[0]) == "incr" {
			return command.AtomicIncrement(commandArgList)
		} else if strings.ToLower(commandArgList[0]) == "decr" {
			return command.AtomicDecrement(commandArgList)
		} else if strings.ToLower(commandArgList[0]) == "del" {
			return command.Delete(commandArgList)
		} else if strings.ToLower(commandArgList[0]) == "lpush" {
			return command.LeftPush(commandArgList)
		} else if strings.ToLower(commandArgList[0]) == "rpush" {
			return command.RightPush(commandArgList)
		} else if strings.ToLower(commandArgList[0]) == "save" {
			return command.SaveOnDisk()
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
		decodedStr := decodeRespString(message)
		// Echo the message back to the client
		conn.Write([]byte(fmt.Sprintf("+%s\r\n", decodedStr)))
	}
}

func main() {
	newHashMap, err := utility.ReadFromRDB("backup.rdb")
	if err == nil {
		store.HashMap = newHashMap
	}
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
