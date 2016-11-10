package main

import (
	"errors"
	"log"
	"net"
)

type MomoTunnel struct {
	src      string   // "e.g :8081"
	dst      string   // "e.g cseweb.ucsd.edu:80"
	signal   chan int //write anything into this channel will stop everything
	listener net.Listener
}

func (self *MomoTunnel) Start() error {
	var err error
	self.signal = make(chan int, 1)
	self.listener, err = net.Listen("tcp", self.src)
	if err != nil {
		return err
	}
	go self.Run()
	return nil
}

func (self *MomoTunnel) Stop() error {
	if self.signal == nil {
		err := errors.New("momotunnel is not running")
		return err
	}
	self.signal <- 1

	// Close the listener
	self.listener.Close()
	return nil
}

func (self *MomoTunnel) Join() error {
	if self.signal == nil {
		err := errors.New("momotunnel is not running")
		return err
	}
	signal := <-self.signal
	self.signal <- signal
	return nil
}

func (self *MomoTunnel) Run() {
	//This is a block IO
	for {
		conn, err := self.listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go self.HandleConnection(conn)
	}
	defer self.listener.Close()
}

func (self *MomoTunnel) HandleConnection(local_sock net.Conn) {
	remote_sock, err := net.Dial("tcp", self.dst)
	if err != nil {
		log.Println(err)
		return
	}
	momo1 := MomoForward{local_sock, remote_sock}
	momo2 := MomoForward{remote_sock, local_sock}
	go momo1.Dup()
	go momo2.Dup()
	self.Join()
	momo1.Close()
	momo2.Close()
}
