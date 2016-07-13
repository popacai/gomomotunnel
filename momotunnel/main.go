package main

import (
	"fmt"
	"os"
	"os/signal"
	// use cli pakcage in future
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: momotunnel src_host:src_port dst_host:dst_port")
		return
	}
	src := os.Args[1]
	dst := os.Args[2]
	// src = ":8000"
	// dst = "popacai.com:8000"

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	tunnel := MomoTunnel{src, dst, nil, nil}
	tunnel.Start()
	<-c
	tunnel.Stop()
}
