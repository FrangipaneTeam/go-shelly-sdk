package shelly

import "github.com/FrangipaneTeam/go-shelly-sdk/internal/rpc"

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

// Close closes the client.
func (c *Client) Close() error {
	return c.rpc.Close()
}

// * CLOUD

// Cloud is the client for the Cloud.
type CloudClient struct {
	client *Client
}

// Cloud returns a client for the Cloud.
func (c *Client) Cloud() *CloudClient {
	return &CloudClient{client: c}
}
