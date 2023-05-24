package rpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/ybbus/jsonrpc/v3"
)

var ErrMissingRequiredArgs = errors.New("missing required args")

type RPC struct {
	rpc jsonrpc.RPCClient
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

	rpc := jsonrpc.NewClientWithOpts("http://"+rA.IP+"/rpc", &jsonrpc.RPCClientOpts{
		AllowUnknownFields: true,
	})

	return &RPC{ip: rA.IP, rpc: rpc}, nil
}

// Call calls the method on the server.
func (c *RPC) Call(method string, args, reply any) error {
	response, err := c.rpc.Call(context.Background(), method, args)
	if err != nil {
		return fmt.Errorf("could not call %s on %s: %w", method, c.ip, err)
	}

	err = response.GetObject(&reply) // expects a rpc-object result value like: {"id": 123, "name": "alex", "age": 33}
	if err != nil {
		return fmt.Errorf("could not read response: %w", err)
	}

	return nil
}
