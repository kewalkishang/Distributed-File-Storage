package p2p

//Install the package
//go get github.com/stretchr/testify
//go mod tidy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPTransport(listenAddr)
	assert.Equal(t, tr.listenAddress, listenAddr)

	//Server
	assert.Nil(t, tr.ListenAndAccept())
}
