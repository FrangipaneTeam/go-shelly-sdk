// code generated by go generate - look at command.go.tmpl for source file
package shelly

/*
 > Client Shelly
*/

// Shelly is the client for the Shelly.
type ShellyClient struct {
	client *Client
}

// Shelly returns a client for the Shelly.
func (c *Client) Shelly() *ShellyClient {
	return &ShellyClient{client: c}
}

/*
 > Client Switch
*/

// Switch is the client for the Switch.
type SwitchClient struct {
	client *Client
}

// Switch returns a client for the Switch.
func (c *Client) Switch() *SwitchClient {
	return &SwitchClient{client: c}
}

/*
 > Client Light
*/

// Light is the client for the Light.
type LightClient struct {
	client *Client
}

// Light returns a client for the Light.
func (c *Client) Light() *LightClient {
	return &LightClient{client: c}
}

/*
 > Client Cover
*/

// Cover is the client for the Cover.
type CoverClient struct {
	client *Client
}

// Cover returns a client for the Cover.
func (c *Client) Cover() *CoverClient {
	return &CoverClient{client: c}
}

/*
 > Client DevicePower
*/

// DevicePower is the client for the DevicePower.
type DevicePowerClient struct {
	client *Client
}

// DevicePower returns a client for the DevicePower.
func (c *Client) DevicePower() *DevicePowerClient {
	return &DevicePowerClient{client: c}
}
