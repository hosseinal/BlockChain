package network

import (
	"fmt"
	"time"
)

type ServerOpts struct {
	Transport []Transport
}

type Server struct {
	ServerOpts
	rpcChan  chan RPC
	quitChan chan RPC
}

var serverInstance *Server

func NewServer(opts ServerOpts) *Server {
	// Singleton transport new instance

	serverInstance = &Server{
		ServerOpts: opts,
		rpcChan:    make(chan RPC, 1024),
		quitChan:   make(chan RPC),
	}

	return serverInstance
}

func (s *Server) Start() {
	// Start server logic here
	ticker := time.NewTicker(time.Second * 5)

	s.initTransport()

loop:
	for {
		select {
		case <-ticker.C:
			// Periodic tasks can be handled here
			fmt.Println(" Server is working")
		case rpc := <-s.rpcChan:
			// Handle incoming RPCs here
			fmt.Printf("Received RPC from %s with payload: %s\n", rpc.From, string(rpc.Payload))
		case <-s.quitChan:
			ticker.Stop()
			break loop
		}
	}
}

func (s *Server) initTransport() {
	// Initialize transport logic here
	for _, tr := range s.Transport {
		go func(tr Transport) {

			for rpc := range tr.Consume() {
				s.rpcChan <- rpc
			}

		}(tr)
	}
}

func (s *Server) AppendTransportLayerToServer(tr Transport) {

	serverInstance.Transport = append(serverInstance.Transport, tr)
	go func() {
		for rpc := range tr.Consume() {
			serverInstance.rpcChan <- rpc
		}
	}()
}

func (s *Server) Stop() {
	s.quitChan <- RPC{}
}
