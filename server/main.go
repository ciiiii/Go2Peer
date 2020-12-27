package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main()  {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		panic("port is empty")
	}
	port, _ := strconv.Atoi(portStr)
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: port})
	if err != nil {
		panic(err)
	}
	peers := make([]net.UDPAddr, 0, 2)
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			log.Printf("error during read: %v", err)
		}
		log.Printf("<%s> %s\n", remoteAddr.String(), data[:n])
		peers = append(peers, *remoteAddr)
		if len(peers) == 2 {
			_, _ = listener.WriteToUDP([]byte(peers[1].String()), &peers[0])
			_, _ = listener.WriteToUDP([]byte(peers[0].String()), &peers[1])
			time.Sleep(time.Second * 8)
			log.Println("exit")
		}
	}
}
