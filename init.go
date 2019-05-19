package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	addr = ":50051"
)

var (
	hostname, _ = os.Hostname()
	queueAddr   = os.Getenv("QUEUE_ADDR")
	queuePort   = os.Getenv("QUEUE_PORT")
	queueUser   = os.Getenv("QUEUE_USER")
	queuePass   = os.Getenv("QUEUE_PASS")
)

func init() {
	type Link struct {
		Network string
		Address string
	}

	links := []Link{
		Link{
			"tcp",
			fmt.Sprintf("%s:%s", queueAddr, queuePort),
		},
	}

	for _, link := range links {
		for {
			if conn, err := net.Dial(link.Network, link.Address); err == nil {
				conn.Close()
				break
			}
			log.Printf("Waiting for %s://%s", link.Network, link.Address)
		}
	}
}
