package main

import (
	"make-blockchain/network"
	"time"
)

// server
// transport==>tcp/udp
// block
// tx
// keypair
func main() {

	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("hello world"))
			time.Sleep(1 * time.Second)
		}
	}()
	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}
	s := network.NewServer(opts)
	s.Start()
}
