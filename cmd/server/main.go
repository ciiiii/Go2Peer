package main

import (
	"os"
	"strconv"

	"github.com/ciiiii/Go2Peer/server"
)

func main()  {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		panic("port is empty")
	}
	port, _ := strconv.Atoi(portStr)
	server.StartServer(port)
}
