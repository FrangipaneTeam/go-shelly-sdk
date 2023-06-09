// code generated by go generate - look at command.go.tmpl for source file
package shelly

/*
 > Command DevicePower.GetStatus
 > Get DevicePower status
*/

// DevicePowerGetStatusRequest is the request of GetStatus.
type DevicePowerGetStatusRequest struct {
	// Id of the DevicePower component instance
	Id float64 `json:"id"`
}

// DevicePowerGetStatusResponse is the response of GetStatus.
type DevicePowerGetStatusResponse struct {
	// Information about the battery charge
	Battery DevicePowerGetStatusResponseBattery `json:"battery"`
	// Whether external power source is connected
	Errors []string `json:"errors"`
	// Information about the external power source (only available if external power source is supported)
	External DevicePowerGetStatusResponseExternal `json:"external"`
	// Id of the DevicePower component instance
	Id float64 `json:"id"`
}

// DevicePowerGetStatusResponseBattery is the response of DevicePowerGetStatusResponseBattery.
type DevicePowerGetStatusResponseBattery struct {
	// Optional. Battery charge level in % (null if valid value could not be obtained)
	Percent float64 `json:"percent"`
	// Optional. Battery voltage in Volts (null if valid value could not be obtained)
	V float64 `json:"v"`
}

// DevicePowerGetStatusResponseExternal is the response of DevicePowerGetStatusResponseExternal.
type DevicePowerGetStatusResponseExternal struct {
	// Whether external power source is connected
	Present bool `json:"present"`
}

// readResponse reads the response into the given interface.
func (r *DevicePowerGetStatusResponse) readResponse(reader *responseReader) error {
	if reader.Response == nil {
		return ErrInvalidResponse
	}
	return reader.Read(r)
}

// GetStatus Get DevicePower status
func (c DevicePowerClient) GetStatus(args DevicePowerGetStatusRequest) (resp *DevicePowerGetStatusResponse, err error) {
	reader := NewResponseReader()

	if err = c.call("DevicePower.GetStatus", args, &reader.Response); err != nil {
		return
	}

	resp = &DevicePowerGetStatusResponse{}
	return resp, resp.readResponse(reader)
}

// Getbattery returns the battery value.
func (r *DevicePowerGetStatusResponse) GetBattery() DevicePowerGetStatusResponseBattery {
	return r.Battery
}

// Geterrors returns the errors value.
func (r *DevicePowerGetStatusResponse) GetErrors() []string {
	return r.Errors
}

// Getexternal returns the external value.
func (r *DevicePowerGetStatusResponse) GetExternal() DevicePowerGetStatusResponseExternal {
	return r.External
}

// Getid returns the id value.
func (r *DevicePowerGetStatusResponse) GetId() float64 {
	return r.Id
}

// GetPercent returns the percent value.
func (r *DevicePowerGetStatusResponseBattery) GetPercent() float64 {
	return r.Percent
}

// GetV returns the v value.
func (r *DevicePowerGetStatusResponseBattery) GetV() float64 {
	return r.V
}

// GetPresent returns the present value.
func (r *DevicePowerGetStatusResponseExternal) GetPresent() bool {
	return r.Present
}
