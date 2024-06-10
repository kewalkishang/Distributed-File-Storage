package p2p

import "errors"

var ErrInvalidHandshake = errors.New("Invalid Handshake")

type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error { return nil }
