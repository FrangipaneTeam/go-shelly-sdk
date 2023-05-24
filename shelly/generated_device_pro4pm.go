package shelly

type Pro4PM struct {
	client *Client
}

func (c *DevicesClient) Pro4PM() *Pro4PM {
	return &Pro4PM{client: c.client}
}

/*
   HTTP - [GET POST Request]
*/

type Pro4PMHTTPClient struct {
	client *HTTPClient
}

func (c *Pro4PM) HTTP() *Pro4PMHTTPClient {
	return &Pro4PMHTTPClient{client: c.client.API().HTTP()}
}

func (c *Pro4PMHTTPClient) GET(args HTTPGETRequest) (*HTTPGETResponse, error) {
	return c.client.GET(args)
}

func (c *Pro4PMHTTPClient) POST(args HTTPPOSTRequest) (*HTTPPOSTResponse, error) {
	return c.client.POST(args)
}

func (c *Pro4PMHTTPClient) Request(args HTTPRequestRequest) (*HTTPRequestResponse, error) {
	return c.client.Request(args)
}

/*
   Schedule - [Create Update List Delete DeleteAll]
*/

type Pro4PMScheduleClient struct {
	client *ScheduleClient
}

func (c *Pro4PM) Schedule() *Pro4PMScheduleClient {
	return &Pro4PMScheduleClient{client: c.client.API().Schedule()}
}

func (c *Pro4PMScheduleClient) Create(args ScheduleCreateRequest) (*ScheduleCreateResponse, error) {
	return c.client.Create(args)
}

func (c *Pro4PMScheduleClient) Update(args ScheduleUpdateRequest) (*ScheduleUpdateResponse, error) {
	return c.client.Update(args)
}

func (c *Pro4PMScheduleClient) List(args ScheduleListRequest) (*ScheduleListResponse, error) {
	return c.client.List(args)
}

func (c *Pro4PMScheduleClient) Delete(args ScheduleDeleteRequest) (*ScheduleDeleteResponse, error) {
	return c.client.Delete(args)
}

func (c *Pro4PMScheduleClient) DeleteAll(args ScheduleDeleteAllRequest) (*ScheduleDeleteAllResponse, error) {
	return c.client.DeleteAll(args)
}

/*
   Shelly - [ListMethods GetDeviceInfo]
*/

type Pro4PMShellyClient struct {
	client *ShellyClient
}

func (c *Pro4PM) Shelly() *Pro4PMShellyClient {
	return &Pro4PMShellyClient{client: c.client.API().Shelly()}
}

func (c *Pro4PMShellyClient) ListMethods(args ShellyListMethodsRequest) (*ShellyListMethodsResponse, error) {
	return c.client.ListMethods(args)
}

func (c *Pro4PMShellyClient) GetDeviceInfo(args ShellyGetDeviceInfoRequest) (*ShellyGetDeviceInfoResponse, error) {
	return c.client.GetDeviceInfo(args)
}

/*
   Switch - [SetConfig GetConfig GetStatus Toggle Set]
*/

type Pro4PMSwitchClient struct {
	client *SwitchClient
}

func (c *Pro4PM) Switch() *Pro4PMSwitchClient {
	return &Pro4PMSwitchClient{client: c.client.API().Switch()}
}

func (c *Pro4PMSwitchClient) SetConfig(args SwitchSetConfigRequest) (*SwitchSetConfigResponse, error) {
	return c.client.SetConfig(args)
}

func (c *Pro4PMSwitchClient) GetConfig(args SwitchGetConfigRequest) (*SwitchGetConfigResponse, error) {
	return c.client.GetConfig(args)
}

func (c *Pro4PMSwitchClient) GetStatus(args SwitchGetStatusRequest) (*SwitchGetStatusResponse, error) {
	return c.client.GetStatus(args)
}

func (c *Pro4PMSwitchClient) Toggle(args SwitchToggleRequest) (*SwitchToggleResponse, error) {
	return c.client.Toggle(args)
}

func (c *Pro4PMSwitchClient) Set(args SwitchSetRequest) (*SwitchSetResponse, error) {
	return c.client.Set(args)
}
