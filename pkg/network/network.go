package network

import (
	"net"
)

type Conn interface {
	net.Conn
	SetContext(ctx any)
	Context() any
	InboundBuffered() int
	RemoteAddr() net.Addr
}
