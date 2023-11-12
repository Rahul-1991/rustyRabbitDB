package server

import (
	"fmt"
	"net"
	"os"
)

func CreateTCPConnection() net.Listener {
	address := "0.0.0.0:6379"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	//defer listener.Close()

	fmt.Println("Server is listening on", address)
	return listener
}
