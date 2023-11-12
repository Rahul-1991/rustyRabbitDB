package main

import (
	"fmt"
	command "goredis/commands"
	parser "goredis/parser"
	"goredis/server"
	"goredis/store"
	"goredis/utility"
	"io"
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

	// Use a dynamic buffer to accumulate the data
	var buffer []byte
	tempBuffer := make([]byte, 1024)

	for {
		// Read a chunk of data from the client
		bytesRead, err := conn.Read(tempBuffer)
		if err != nil {
			if err == io.EOF {
				// Connection closed by the client
				break
			}
			fmt.Println("Error reading from client:", err)
			return
		}

		// Append the received chunk to the buffer
		buffer = append(buffer, tempBuffer[:bytesRead]...)

		// Check if the message is complete (ending with "\r\n")
		if len(buffer) >= 2 && buffer[len(buffer)-2] == '\r' && buffer[len(buffer)-1] == '\n' {
			// Process the complete message (excluding "\r\n")
			message := string(buffer[:len(buffer)-2])
			decodedStr := decodeRespString(message)

			// Echo the message back to the client
			conn.Write([]byte(fmt.Sprintf("+%s\r\n", decodedStr)))

			// Clear the buffer for the next message
			buffer = nil
		}
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
