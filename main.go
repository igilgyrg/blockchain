package main

import (
	"time"

	"github.com/igilgyrg/blockchain/network"
)

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

	srv := network.NewServer(opts)
	srv.Start()
}
