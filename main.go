package main

import (
	"time"

	network "github.com/hosseinal/BlockChain/Network"
)

func main() {
	trlocal := network.NewLocalTransport("local1")
	trlocal2 := network.NewLocalTransport("local2")

	trlocal.Connect(trlocal2)
	trlocal2.Connect(trlocal)

	go func() {
		for {
			trlocal2.SendMessage(trlocal.Addr(), []byte("Hello from local2"))
			time.Sleep(time.Second * 5)
		}
	}()

	opts := network.ServerOpts{
		Transport: []network.Transport{trlocal},
	}

	server := network.NewServer(opts)
	server.Start()

}
