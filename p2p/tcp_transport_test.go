package p2p

import (
	"testing"
)

func TestTCPTransport(t *testing.T) {
	opts := TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       DefaultDecoder{},
	}
	tr := NewTCPTransport(opts)
	if tr.ListenAddr != ":3000" {
		t.Fatalf("expected listen address %q, got %q", ":3000", tr.ListenAddr)
	}

	if err := tr.ListenAndAccept(); err != nil {
		t.Fatalf("expected transport to listen: %v", err)
	}
	t.Cleanup(func() { tr.Close() })
}
