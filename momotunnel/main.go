package main

import (
	"fmt"
	"os"
	"os/signal"
	// use cli pakcage in future
	"flag"
)

func main() {
	c := make(chan os.Signal, 5)
	signal.Notify(c, os.Interrupt)

	src := flag.String("src", ":5000", "source ip and port, such as :7000")
	dst := flag.String("dst", "localhost:80", "destination ip and port, such as popacai.com:80")
	// src = ":8000"
	// dst = "popacai.com:8000"
	flag.Parse()

	tunnel := MomoTunnel{*src, *dst, nil, nil}
	err := tunnel.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ctrl+c to exit")
	<-c
	tunnel.Stop()
}
