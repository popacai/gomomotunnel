package main

import (
	"io"
	"log"
	"net"
)

type MomoForward struct {
	src net.Conn
	dst net.Conn
}

func (self MomoForward) Dup() {
	for {
		n, err := io.Copy(self.src, self.dst)
		if err != nil || n == 0 {
			log.Println("connection error, close both sockets")
			break
		}
	}
	self.Close()
}

func (self MomoForward) Close() {
	self.src.Close()
	self.dst.Close()
}
