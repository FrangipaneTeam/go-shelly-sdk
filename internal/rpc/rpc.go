package rpc

import (
	"errors"
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"

	"golang.org/x/net/websocket"
)

var ErrMissingRequiredArgs = errors.New("missing required args")

type RPC struct {
	ws  *websocket.Conn
	rpc *rpc.Client
	ip  string
}

type RequiredArgs struct {
	IP string
}

type OptionalArgs struct {
	origin string
}
type Args func(*OptionalArgs)

// WithCustomOrigin sets the origin for the websocket connection.
func WithCustomOrigin(origin string) Args {
	return func(o *OptionalArgs) {
		o.origin = origin
	}
}

// New creates a new client.
func New(rA RequiredArgs, oA ...Args) (*RPC, error) {
	var o OptionalArgs
	for _, opt := range oA {
		opt(&o)
	}

	// Check if all required args are set
	if rA.IP == "" {
		return nil, ErrMissingRequiredArgs
	}

	if o.origin == "" {
		o.origin = "http://localhost"
	}

	ws, err := websocket.Dial("ws://"+rA.IP+"/rpc", "", o.origin)
	if err != nil {
		return nil, fmt.Errorf("could not connect to %s: %w", rA.IP, err)
	}

	rpc := jsonrpc.NewClient(ws)

	return &RPC{ip: rA.IP, rpc: rpc, ws: ws}, nil
}

// Call calls the method on the server.
func (c *RPC) Call(method string, args, reply interface{}) error {
	if err := c.rpc.Call(method, args, reply); err != nil {
		return fmt.Errorf("could not call %s on %s: %w", method, c.ip, err)
	}

	return nil
}

// Close closes the connection.
func (c *RPC) Close() error {
	if err := c.rpc.Close(); err != nil {
		return err
	}

	return c.ws.Close()
}
