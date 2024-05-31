package main

import (
	"context"
	_ "net/http/pprof"

	"github.com/leslie-fei/ghttp/pkg/network"
	"github.com/leslie-fei/ghttp/pkg/network/gnet"
	"github.com/leslie-fei/ghttp/pkg/protocol"
)

var hello = []byte("HelloWorld!")

func main() {
	ts := gnet.NewTransporter("tcp://:8092", true)

	/*	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			panic(err)
		}
	}()*/

	srv := &protocol.Server{
		Handler: func(ctx *protocol.RequestCtx) {
			_, _ = ctx.Write(hello)
		},
	}

	/*go func() {
		time.Sleep(time.Second)
		conn, _ := net.Dial("tcp", "127.0.0.1:8088")
		conn.Write([]byte("GET / HTTP/1.1\r\nHost: 127.0.0.1:8088\r\nHeader1: Value1\r"))

		time.Sleep(time.Second * 5)
		conn.Write([]byte("\nHeader2: Value2\r\n\r\n"))
	}()*/

	err := ts.ListenAndServe(func(ctx context.Context, conn interface{}) error {
		c := conn.(network.Conn)
		return srv.Serve(ctx, c)
	})

	if err != nil {
		panic(err)
	}
}
