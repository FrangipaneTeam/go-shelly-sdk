package shelly

import (
	"github.com/FrangipaneTeam/go-shelly-sdk/internal/rpc"
)

type Client struct {
	rpc *rpc.RPC
}

// New creates a new client.
func New(ip string) (*Client, error) {
	_rpc, err := rpc.New(rpc.RequiredArgs{IP: ip})
	if err != nil {
		return nil, err
	}

	return &Client{rpc: _rpc}, nil
}

/*
   > APIClient
*/

type APIClient struct {
	client *Client
}

// API returns a client for the API.
func (c *Client) API() *APIClient {
	return &APIClient{client: c}
}

// call calls the given method with the given args and reply.
func (c *APIClient) call(method string, args, reply interface{}) error {
	// pretty.Print(args)
	return c.client.rpc.Call(method, args, reply)
}

type DevicesClient struct {
	client *Client
}

// Devices returns a client for the Devices.
func (c *Client) Devices() *DevicesClient {
	return &DevicesClient{client: c}
}
