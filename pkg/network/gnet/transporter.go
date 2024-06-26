package gnet

import (
	"context"
	"errors"

	"github.com/leslie-fei/ghttp/pkg/errs"
	"github.com/leslie-fei/ghttp/pkg/network"
	"github.com/panjf2000/gnet/v2"
	"github.com/panjf2000/gnet/v2/pkg/logging"
)

func NewTransporter(addr string, opts ...gnet.Option) network.Transporter {
	t := &transporter{addr: addr, opts: opts}
	return t
}

type Option func(t *transporter)

type transporter struct {
	addr    string
	opts    []gnet.Option
	engine  gnet.Engine
	handler *handler
}

func (t *transporter) Shutdown(ctx context.Context) error {
	return t.Close()
}

func (t *transporter) Close() error {
	if t.handler == nil {
		return nil
	}
	return t.handler.Stop(context.Background())
}

func (t *transporter) ListenAndServe(onData network.OnData) error {
	h := &handler{
		onData: onData,
	}
	t.handler = h
	return gnet.Run(h, t.addr, t.opts...)
}

type handler struct {
	gnet.BuiltinEventEngine
	onData network.OnData
	engine gnet.Engine
}

func (h *handler) Stop(ctx context.Context) error {
	return h.engine.Stop(ctx)
}

func (h *handler) OnBoot(eng gnet.Engine) (action gnet.Action) {
	h.engine = eng
	return
}

func (h *handler) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	return
}

func (h *handler) OnClose(c gnet.Conn, _ error) (action gnet.Action) {
	if c.Context() != nil {
		if r, ok := c.Context().(interface{ Release() }); ok {
			r.Release()
		}
	}
	return
}

func (h *handler) OnTraffic(c gnet.Conn) (action gnet.Action) {
	for c.InboundBuffered() > 0 {
		if err := h.onData(context.Background(), c); err != nil {
			if errors.Is(err, errs.ErrNeedMore) {
				return gnet.None
			}
			logging.Errorf("OnData error: %v", err)
			return gnet.Close
		}
	}
	return gnet.None
}
