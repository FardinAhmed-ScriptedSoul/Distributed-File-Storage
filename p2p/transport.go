package p2p

import "net"

// Peer represents a connected remote node. Embedding net.Conn allows the file
// server to read and write streamed file data through the peer.
type Peer interface {
	net.Conn
	Send([]byte) error
	CloseStream()
}

// Transport handles communication between nodes. TCP is the current
// implementation, but the interface leaves room for other transports.
type Transport interface {
	ListenAndAccept() error
	Addr() string
	Consume() <-chan RPC
	Dial(string) error
	Close() error
}
