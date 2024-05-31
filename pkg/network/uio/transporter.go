package uio

import (
	"context"
	"errors"
	"time"

	"github.com/leslie-fei/ghttp/pkg/errs"
	"github.com/leslie-fei/ghttp/pkg/network"
	"github.com/panjf2000/gnet/v2/pkg/logging"
	"github.com/urpc/uio"
)

func NewTransporter(addr string, numLoop int) network.Transporter {
	var events uio.Events
	events.Pollers = numLoop
	events.Addrs = []string{addr}
	return &transporter{events: &events}
}

type transporter struct {
	events *uio.Events
}

func (t *transporter) Close() error {
	return nil
}

func (t *transporter) Shutdown(ctx context.Context) error {
	return t.Close()
}

func (t *transporter) ListenAndServe(onData network.OnData) error {
	t.events.OnData = func(c uio.Conn) error {
		for c.InboundBuffered() > 0 {
			if err := onData(context.Background(), &uioConn{c}); err != nil {
				if errors.Is(err, errs.ErrNeedMore) {
					return nil
				}
				logging.Errorf("OnData error: %v", err)
				return err
			}
		}
		return nil
	}
	return t.events.Serve()
}

type uioConn struct {
	uio.Conn
}

func (u *uioConn) SetDeadline(t time.Time) error {
	//TODO implement me
	panic("implement me")
}

func (u *uioConn) SetReadDeadline(t time.Time) error {
	//TODO implement me
	panic("implement me")
}

func (u *uioConn) SetWriteDeadline(t time.Time) error {
	//TODO implement me
	panic("implement me")
}
