package gnet

import (
	"context"
	"time"

	"github.com/leslie-fei/webapp/pkg/network"
	"github.com/panjf2000/gnet/v2"
	"github.com/panjf2000/gnet/v2/pkg/logging"
)

func NewTransporter(addr string, multicore bool) network.Transporter {
	return &transporter{addr: addr, multicore: multicore}
}

type transporter struct {
	addr      string
	multicore bool
	engine    gnet.Engine
	handler   *handler
}

func (t *transporter) Shutdown(ctx context.Context) error {
	return t.Close()
}

func (t *transporter) Close() error {
	if t.handler == nil {
		return nil
	}
	return t.handler.Close(context.Background())
}

func (t *transporter) ListenAndServe(onData network.OnData) error {
	h := &handler{
		onData: onData,
	}
	t.handler = h
	return gnet.Run(h, t.addr, gnet.WithMulticore(t.multicore))
}

type handler struct {
	gnet.BuiltinEventEngine
	onData network.OnData
	engine gnet.Engine
}

func (h *handler) Close(ctx context.Context) error {
	return h.engine.Stop(ctx)
}

func (h *handler) OnBoot(eng gnet.Engine) (action gnet.Action) {
	h.engine = eng
	return
}

func (h *handler) OnTraffic(c gnet.Conn) (action gnet.Action) {
	if err := h.onData(context.Background(), &conn{c}); err != nil {
		logging.Errorf("OnData error: %v", err)
		return gnet.Close
	}
	return gnet.None
}

type conn struct {
	gnet.Conn
}

func (c *conn) Skip(n int) error {
	_, err := c.Conn.Discard(n)
	return err
}

func (c *conn) Release() error {
	return nil
}

func (c *conn) Len() int {
	return c.Conn.InboundBuffered()
}

func (c *conn) ReadByte() (byte, error) {
	b, err := c.Next(1)
	if err != nil {
		return 0, err
	}
	return b[0], nil
}

func (c *conn) ReadBinary(n int) (p []byte, err error) {
	return c.Next(n)
}

func (c *conn) Malloc(n int) (buf []byte, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *conn) WriteBinary(b []byte) (n int, err error) {
	return c.Conn.Write(b)
}

func (c *conn) SetReadTimeout(t time.Duration) error {
	panic("implement me")
}

func (c *conn) SetWriteTimeout(t time.Duration) error {
	//TODO implement me
	panic("implement me")
}
