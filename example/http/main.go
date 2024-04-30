package main

import (
	"context"

	"github.com/leslie-fei/webapp/pkg/network"
	"github.com/leslie-fei/webapp/pkg/network/gnet"
	"github.com/leslie-fei/webapp/pkg/protocol"
)

var hello = []byte("HelloWorld!")

func main() {
	ts := gnet.NewTransporter("tcp://:8080", true)

	srv := &protocol.Server{
		Handler: func(ctx *protocol.RequestCtx) {
			_, _ = ctx.Write(hello)
		},
	}

	err := ts.ListenAndServe(func(ctx context.Context, conn interface{}) error {
		c := conn.(network.Conn)
		return srv.Serve(ctx, c)
	})

	if err != nil {
		panic(err)
	}
}
