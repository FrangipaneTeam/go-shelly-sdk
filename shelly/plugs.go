package shelly

type PlugS struct {
	client *Client
}

func (c *DevicesClient) PlugS() *PlugS {
	return &PlugS{client: c.client}
}

/*
 > Client Switch
*/

// Switch is the client for the Switch.
type PlugSSwitchClient struct {
	client *Client
}

// Switch returns a client for the Switch.
func (c *PlugS) Switch() *PlugSSwitchClient {
	return &PlugSSwitchClient{client: c.client}
}

// call calls the given method with the given args and reply.
func (c *PlugSSwitchClient) call(method string, args, reply interface{}) error {
	return c.client.rpc.Call(method, args, reply)
}

/*
 > Command Switch.SetConfig
 >
*/

// SwitchSetConfigRequest is the request of SetConfig.
type PlugSSwitchSetConfigRequest struct {
	// Optional. Configuration that the method takes.
	Config string `json:"config,omitempty"`
	// Id of the Switch component instance.
	Id float64 `json:"id"`
}

// SwitchSetConfigResponse is the response of SetConfig.
type PlugSSwitchSetConfigResponse struct {
}

// readResponse reads the response into the given interface.
func (r *PlugSSwitchSetConfigResponse) readResponse(reader *responseReader) error {
	if reader.Response == nil {
		return ErrInvalidResponse
	}
	return reader.Read(r)
}

// SetConfig
func (c PlugSSwitchClient) SetConfig(args SwitchSetConfigRequest) (resp *SwitchSetConfigResponse, err error) {
	reader := NewResponseReader()

	if err = c.call("Switch.SetConfig", args, &reader.Response); err != nil {
		return
	}

	resp = &SwitchSetConfigResponse{}
	return resp, resp.readResponse(reader)
}
