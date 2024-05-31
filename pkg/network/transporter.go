package network

import "context"

type Transporter interface {
	// Close the transporter immediately
	Close() error

	// Graceful shutdown the transporter
	Shutdown(ctx context.Context) error

	// Start listen and ready to accept connection
	ListenAndServe(onData OnData) error
}

// Callback when data is ready on the connection
type OnData func(ctx context.Context, conn Conn) error
