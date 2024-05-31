package main

import (
	"context"
	"runtime"

	"github.com/leslie-fei/ghttp/pkg/network"
	"github.com/leslie-fei/ghttp/pkg/network/uio"
	"github.com/leslie-fei/ghttp/pkg/protocol"
)

var hello = []byte("HelloWorld!")

func main() {
	//ts := gnet.NewTransporter("tcp://:8092", gnet2.WithNumEventLoop(3))
	ts := uio.NewTransporter("tcp://:8091", runtime.NumCPU())

	srv := &protocol.Server{
		Handler: func(ctx *protocol.RequestCtx) {
			_, _ = ctx.Write(hello)
		},
		//Executor: func(fn func()) {
		//	go fn()
		//},
	}

	/*go func() {
		time.Sleep(time.Second)
		conn, _ := net.Dial("tcp", "127.0.0.1:8088")
		conn.Write([]byte("GET / HTTP/1.1\r\nHost: 127.0.0.1:8088\r\nHeader1: Value1\r"))

		time.Sleep(time.Second * 5)
		conn.Write([]byte("\nHeader2: Value2\r\n\r\n"))
	}()*/

	err := ts.ListenAndServe(func(ctx context.Context, conn network.Conn) error {
		return srv.Serve(ctx, conn)
	})

	if err != nil {
		panic(err)
	}
}
