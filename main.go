package main

//telnet localhost 3000 - To connect to this server

import (
	"fmt"
	"log"

	"github.com/kewalkishang/Distributed-File-Storage/p2p"
)

func OnPeer(peer p2p.Peer) error {
	fmt.Printf("Doing something outside")
	return nil
}

//	func(p p2p.Peer) error {
//		return fmt.Errorf("failed the onPeer func")
//	}
func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
