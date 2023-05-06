// code generated by go generate - look at command.go.tmpl for source file
package shelly

/*
 > Command Cover.GetConfig
 >
*/

// CoverGetConfigRequest is the request of GetConfig.
type CoverGetConfigRequest struct {
	Id string `json:"id"` // Id of the Cover component instance.
}

// CoverGetConfigResponse is the response of GetConfig.
type CoverGetConfigResponse struct {
	CurrentLimit         int                                        `json:"current_limit"`         // Amperes, limit that must be exceeded to trigger an overcurrent error.
	Id                   string                                     `json:"id"`                    // Id of the Cover component instance.
	InMode               string                                     `json:"in_mode,omitempty"`     // One of single, dual or detached, only present if there is at least one input associated with the Cover instance. Single - Cover operation in both open and close directions is controlled via a single input. In this mode, only input_0 is used to open/close/stop the Cover. It doesn't matter if input_0 has in_type=switch or in_type=button, the behavior is the same: each switch toggle or button press cycles between open/stop/close/stop/... In single mode, input_1 is free to be used as a safety switch (e.g. end-of-motion limit switch, emergency-stop, etc.). dual - Cover operation is controlled via two inputs, one for open and one for close. In this mode, input_0 is used to open the Cover, input_1 is used to close the Cover.The exact behavior depends on the in_type of the inputs: if in_type = switch: toggle the switch to ON to move in the associated direction; toggle the switch to OFF to stop, if in_type = button: press the button to move in the associated direction; press the button again to stop. detached - Cover operation via the input/inputs is prohibited.
	InitialState         string                                     `json:"initial_state"`         // Defines Cover target state on power-on, one of open (Cover will fully open), closed (Cover will fully close) or stopped (Cover will not change its position).
	InvertDirections     bool                                       `json:"invert_directions"`     // Defines the motor rotation for open and close directions (changing this parameter requires a reboot). false - On open motor rotates clockwise, on close motor rotates counter-clockwise. true - On open motor rotates counter-clockwise, on close motor rotates clockwise.
	MaxtimeClose         int                                        `json:"maxtime_close"`         // Default timeout after which Cover will stop moving in close direction.
	MaxtimeOpen          int                                        `json:"maxtime_open"`          // Default timeout after which Cover will stop moving in open direction.
	Motor                CoverGetConfigResponseMotor                `json:"motor"`                 // configuration of the Cover motor. The exact contents depend on the type of motor used. The descriptions below are valid when an AC motor is used.
	Name                 string                                     `json:"name,omitempty"`        // Name of the cover instance.
	ObstructionDetection CoverGetConfigResponseObstructionDetection `json:"obstruction_detection"` // Defines the behavior of the obstruction detection safety feature.
	PowerLimit           int                                        `json:"power_limit"`           // Watts, limit that must be exceeded to trigger an overpower error.
	SafetySwitch         CoverGetConfigResponseSafetySwitch         `json:"safety_switch"`         // Defines the behavior of the safety switch feature, only present if there are two inputs associated with the Cover instance. The safety_switch feature will only work when in_mode=single
	SwapInputs           bool                                       `json:"swap_inputs"`           // Only present if there are two inputs associated with the Cover instance, defines whether the functions of the two inputs are swapped. The effect of swap_inputs is observable only when in_mode != detached. When swap_inputs is false: If in_mode = dual: input_0 is used to open, input_1 is used to close. If in_mode = single: input_0 is used to open/close/stop, input_1 is used as safety switch or is not used at all. When swap_inputs is true: If in_mode = dual: input_0 is used to close, input_1 is used to open. If in_mode = single: input_0 is used as safety switch or is not used at all, input_1 is used to open/close/stop.
	UndervoltageLimit    int                                        `json:"undervoltage_limit"`    // Volts, limit that must be exceeded to trigger an undervoltage error.
	VoltageLimit         int                                        `json:"voltage_limit"`         // Volts, limit that must be exceeded to trigger an undervoltage error.
}

// CoverGetConfigResponseMotor is the response of motor.
type CoverGetConfigResponseMotor struct {
	IdleConfirmPeriod int `json:"idle_confirm_period"` // Seconds, minimum period of time in idle state before state is confirmed.
	IdlePowerThr      int `json:"idle_power_thr"`      // Watts, threshold below which the motor is considered stopped.
}

// CoverGetConfigResponseObstructionDetection is the response of obstruction_detection.
type CoverGetConfigResponseObstructionDetection struct {
	Action    string `json:"action"`    // The recovery action which should be performed if the safety switch is engaged while moving in a monitored direction, one of: stop - Immediately stop Cover. reverse - Immediately stop Cover, then move in the opposite direction until a fully open or fully closed position is reached.
	Direction string `json:"direction"` // The direction of motion for which safety switch should be monitored, one of open, close, both
	Enable    bool   `json:"enable"`    // true when obstruction detection is enabled, false otherwise
	Holdoff   int    `json:"holdoff"`   // Seconds, time to wait after Cover starts moving before obstruction detection is activated (to avoid false detections because of the initial power consumption spike).
	PowerThr  int    `json:"power_thr"` // Watts, power consumption above this threshold should be interpreted as objects obstructing Cover movement. This property is editable at any time, but note that during the cover calibration procedure (Cover.Calibrate), power_thr will be automatically set to the peak power consumption + 15%, overwriting the current value. The automatic setup of power_thr during calibration will only start tracking power values when the holdoff time (see below) has elapsed.
}

// CoverGetConfigResponseSafetySwitch is the response of safety_switch.
type CoverGetConfigResponseSafetySwitch struct {
	Action      string `json:"action"`       // The recovery action which should be performed if the safety switch is engaged while moving in a monitored direction, one of: stop - Immediately stop Cover. reverse - Immediately stop Cover, then move in the opposite direction until a fully open or fully closed position is reached. pause - Immediately stop Cover, then either: wait for a command to move in an allowed direction (see below) or automatically continue movement in the same direction (i.e. the one that was interrupted) when the safety switch is disengaged
	AllowedMove string `json:"allowed_move"` // Allowed movement direction when the safety switch is engaged while moving in a monitored direction: null - null means Cover can't be moved in neither open nor close directions while the safety switch is engaged. reverse - the only other option is reverse, which means Cover can only be moved in the direction opposite to the one that was interrupted (for example, if the safety switch was hit while opening, Cover can only be commanded to close if the switch is not disengaged)
	Direction   string `json:"direction"`    // The direction of motion for which safety switch should be monitored, one of open, close, both
	Enable      bool   `json:"enable"`       // true when safety switch is enabled, false otherwise
}

// readResponse reads the response into the given interface.
func (r *CoverGetConfigResponse) readResponse(reader *responseReader) error { //nolint:dupl
	if reader.Response == nil {
		return ErrInvalidResponse
	}
	return reader.Read(r)
}

// GetConfig
func (c CoverClient) GetConfig(args CoverGetConfigRequest) (resp *CoverGetConfigResponse, err error) { //nolint:dupl
	reader := NewResponseReader()

	if err = c.client.rpc.Call("Cover.GetConfig", args, &reader.Response); err != nil {
		return
	}

	resp = &CoverGetConfigResponse{}
	return resp, resp.readResponse(reader)
}

// Getcurrent_limit returns the current_limit value.
func (r *CoverGetConfigResponse) GetCurrentLimit() int {
	return r.CurrentLimit
}

// Getid returns the id value.
func (r *CoverGetConfigResponse) GetId() string {
	return r.Id
}

// Getin_mode returns the in_mode value.
func (r *CoverGetConfigResponse) GetInMode() string {
	return r.InMode
}

// Getinitial_state returns the initial_state value.
func (r *CoverGetConfigResponse) GetInitialState() string {
	return r.InitialState
}

// Getinvert_directions returns the invert_directions value.
func (r *CoverGetConfigResponse) GetInvertDirections() bool {
	return r.InvertDirections
}

// Getmaxtime_close returns the maxtime_close value.
func (r *CoverGetConfigResponse) GetMaxtimeClose() int {
	return r.MaxtimeClose
}

// Getmaxtime_open returns the maxtime_open value.
func (r *CoverGetConfigResponse) GetMaxtimeOpen() int {
	return r.MaxtimeOpen
}

// Getmotor returns the motor value.
func (r *CoverGetConfigResponse) GetMotor() CoverGetConfigResponseMotor {
	return r.Motor
}

// Getname returns the name value.
func (r *CoverGetConfigResponse) GetName() string {
	return r.Name
}

// Getobstruction_detection returns the obstruction_detection value.
func (r *CoverGetConfigResponse) GetObstructionDetection() CoverGetConfigResponseObstructionDetection {
	return r.ObstructionDetection
}

// Getpower_limit returns the power_limit value.
func (r *CoverGetConfigResponse) GetPowerLimit() int {
	return r.PowerLimit
}

// Getsafety_switch returns the safety_switch value.
func (r *CoverGetConfigResponse) GetSafetySwitch() CoverGetConfigResponseSafetySwitch {
	return r.SafetySwitch
}

// Getswap_inputs returns the swap_inputs value.
func (r *CoverGetConfigResponse) GetSwapInputs() bool {
	return r.SwapInputs
}

// Getundervoltage_limit returns the undervoltage_limit value.
func (r *CoverGetConfigResponse) GetUndervoltageLimit() int {
	return r.UndervoltageLimit
}

// Getvoltage_limit returns the voltage_limit value.
func (r *CoverGetConfigResponse) GetVoltageLimit() int {
	return r.VoltageLimit
}

// Getidle_confirm_period returns the idle_confirm_period value.
func (r *CoverGetConfigResponseMotor) GetIdleConfirmPeriod() int {
	return r.IdleConfirmPeriod
}

// Getidle_power_thr returns the idle_power_thr value.
func (r *CoverGetConfigResponseMotor) GetIdlePowerThr() int {
	return r.IdlePowerThr
}

// Getaction returns the action value.
func (r *CoverGetConfigResponseObstructionDetection) GetAction() string {
	return r.Action
}

// Getdirection returns the direction value.
func (r *CoverGetConfigResponseObstructionDetection) GetDirection() string {
	return r.Direction
}

// Getenable returns the enable value.
func (r *CoverGetConfigResponseObstructionDetection) GetEnable() bool {
	return r.Enable
}

// Getholdoff returns the holdoff value.
func (r *CoverGetConfigResponseObstructionDetection) GetHoldoff() int {
	return r.Holdoff
}

// Getpower_thr returns the power_thr value.
func (r *CoverGetConfigResponseObstructionDetection) GetPowerThr() int {
	return r.PowerThr
}

// Getaction returns the action value.
func (r *CoverGetConfigResponseSafetySwitch) GetAction() string {
	return r.Action
}

// Getallowed_move returns the allowed_move value.
func (r *CoverGetConfigResponseSafetySwitch) GetAllowedMove() string {
	return r.AllowedMove
}

// Getdirection returns the direction value.
func (r *CoverGetConfigResponseSafetySwitch) GetDirection() string {
	return r.Direction
}

// Getenable returns the enable value.
func (r *CoverGetConfigResponseSafetySwitch) GetEnable() bool {
	return r.Enable
}

/*
 > Command Cover.SetConfig
 >
*/

// CoverSetConfigRequest is the request of SetConfig.
type CoverSetConfigRequest struct {
	Config CoverSetConfigRequestConfig `json:"config"` // The configuration to apply to the Cover instance.
	Id     string                      `json:"id"`     // The ID of the Cover instance to configure.
}

// CoverSetConfigRequestConfigMotor is the request of ConfigMotor.
type CoverSetConfigRequestConfigMotor struct {
	IdleConfirmPeriod int `json:"idle_confirm_period"` // Seconds, minimum period of time in idle state before state is confirmed.
	IdlePowerThr      int `json:"idle_power_thr"`      // Watts, threshold below which the motor is considered stopped.
}

// CoverSetConfigRequestConfigObstructionDetection is the request of ConfigObstructionDetection.
type CoverSetConfigRequestConfigObstructionDetection struct {
	Action    string `json:"action,omitempty"` // The recovery action which should be performed if the safety switch is engaged while moving in a monitored direction, one of: stop - Immediately stop Cover. reverse - Immediately stop Cover, then move in the opposite direction until a fully open or fully closed position is reached.
	Direction string `json:"direction"`        // The direction of motion for which safety switch should be monitored, one of open, close, both
	Enable    bool   `json:"enable"`           // true when obstruction detection is enabled, false otherwise
	Holdoff   int    `json:"holdoff"`          // Seconds, time to wait after Cover starts moving before obstruction detection is activated (to avoid false detections because of the initial power consumption spike).
	PowerThr  int    `json:"power_thr"`        // Watts, power consumption above this threshold should be interpreted as objects obstructing Cover movement. This property is editable at any time, but note that during the cover calibration procedure (Cover.Calibrate), power_thr will be automatically set to the peak power consumption + 15%, overwriting the current value. The automatic setup of power_thr during calibration will only start tracking power values when the holdoff time (see below) has elapsed.
}

// CoverSetConfigRequestConfigSafetySwitch is the request of ConfigSafetySwitch.
type CoverSetConfigRequestConfigSafetySwitch struct {
	Action      string `json:"action,omitempty"`       // The recovery action which should be performed if the safety switch is engaged while moving in a monitored direction, one of: stop - Immediately stop Cover. reverse - Immediately stop Cover, then move in the opposite direction until a fully open or fully closed position is reached. pause - Immediately stop Cover, then either: wait for a command to move in an allowed direction (see below) or automatically continue movement in the same direction (i.e. the one that was interrupted) when the safety switch is disengaged
	AllowedMove string `json:"allowed_move,omitempty"` // Allowed movement direction when the safety switch is engaged while moving in a monitored direction: null - null means Cover can't be moved in neither open nor close directions while the safety switch is engaged. reverse - the only other option is reverse, which means Cover can only be moved in the direction opposite to the one that was interrupted (for example, if the safety switch was hit while opening, Cover can only be commanded to close if the switch is not disengaged)
	Direction   string `json:"direction"`              // The direction of motion for which safety switch should be monitored, one of open, close, both
	Enable      bool   `json:"enable"`                 // true when safety switch is enabled, false otherwise
}

// CoverSetConfigRequestConfig is the request of config.
type CoverSetConfigRequestConfig struct {
	CurrentLimit         int                                             `json:"current_limit"`         // Amperes, limit that must be exceeded to trigger an overcurrent error.
	Id                   string                                          `json:"id"`                    // Id of the Cover component instance.
	InMode               string                                          `json:"in_mode,omitempty"`     // One of single, dual or detached, only present if there is at least one input associated with the Cover instance. Single - Cover operation in both open and close directions is controlled via a single input. In this mode, only input_0 is used to open/close/stop the Cover. It doesn't matter if input_0 has in_type=switch or in_type=button, the behavior is the same: each switch toggle or button press cycles between open/stop/close/stop/... In single mode, input_1 is free to be used as a safety switch (e.g. end-of-motion limit switch, emergency-stop, etc.). dual - Cover operation is controlled via two inputs, one for open and one for close. In this mode, input_0 is used to open the Cover, input_1 is used to close the Cover.The exact behavior depends on the in_type of the inputs: if in_type = switch: toggle the switch to ON to move in the associated direction; toggle the switch to OFF to stop, if in_type = button: press the button to move in the associated direction; press the button again to stop. detached - Cover operation via the input/inputs is prohibited.
	InitialState         string                                          `json:"initial_state"`         // Defines Cover target state on power-on, one of open (Cover will fully open), closed (Cover will fully close) or stopped (Cover will not change its position).
	InvertDirections     bool                                            `json:"invert_directions"`     // Defines the motor rotation for open and close directions (changing this parameter requires a reboot). false - On open motor rotates clockwise, on close motor rotates counter-clockwise. true - On open motor rotates counter-clockwise, on close motor rotates clockwise.
	MaxtimeClose         int                                             `json:"maxtime_close"`         // Default timeout after which Cover will stop moving in close direction.
	MaxtimeOpen          int                                             `json:"maxtime_open"`          // Default timeout after which Cover will stop moving in open direction.
	Motor                CoverSetConfigRequestConfigMotor                `json:"motor"`                 // configuration of the Cover motor. The exact contents depend on the type of motor used. The descriptions below are valid when an AC motor is used.
	Name                 string                                          `json:"name,omitempty"`        // Name of the cover instance.
	ObstructionDetection CoverSetConfigRequestConfigObstructionDetection `json:"obstruction_detection"` // Defines the behavior of the obstruction detection safety feature.
	PowerLimit           int                                             `json:"power_limit"`           // Watts, limit that must be exceeded to trigger an overpower error.
	SafetySwitch         CoverSetConfigRequestConfigSafetySwitch         `json:"safety_switch"`         // Defines the behavior of the safety switch feature, only present if there are two inputs associated with the Cover instance. The safety_switch feature will only work when in_mode=single
	SwapInputs           bool                                            `json:"swap_inputs"`           // Only present if there are two inputs associated with the Cover instance, defines whether the functions of the two inputs are swapped. The effect of swap_inputs is observable only when in_mode != detached. When swap_inputs is false: If in_mode = dual: input_0 is used to open, input_1 is used to close. If in_mode = single: input_0 is used to open/close/stop, input_1 is used as safety switch or is not used at all. When swap_inputs is true: If in_mode = dual: input_0 is used to close, input_1 is used to open. If in_mode = single: input_0 is used as safety switch or is not used at all, input_1 is used to open/close/stop.
	UndervoltageLimit    int                                             `json:"undervoltage_limit"`    // Volts, limit that must be exceeded to trigger an undervoltage error.
	VoltageLimit         int                                             `json:"voltage_limit"`         // Volts, limit that must be exceeded to trigger an undervoltage error.
}

// CoverSetConfigResponse is the response of SetConfig.
type CoverSetConfigResponse struct {
}

// readResponse reads the response into the given interface.
func (r *CoverSetConfigResponse) readResponse(reader *responseReader) error { //nolint:dupl
	if reader.Response == nil {
		return ErrInvalidResponse
	}
	return reader.Read(r)
}

// SetConfig
func (c CoverClient) SetConfig(args CoverSetConfigRequest) (resp *CoverSetConfigResponse, err error) { //nolint:dupl
	reader := NewResponseReader()

	if err = c.client.rpc.Call("Cover.SetConfig", args, &reader.Response); err != nil {
		return
	}

	resp = &CoverSetConfigResponse{}
	return resp, resp.readResponse(reader)
}

/*
 > Command Cover.GetStatus
 > Get Cover status
*/

// CoverGetStatusRequest is the request of GetStatus.
type CoverGetStatusRequest struct {
	Id string `json:"id"` // The numeric ID of the Cover component instance
}

// CoverGetStatusResponse is the response of GetStatus.
type CoverGetStatusResponse struct {
	Aenergy       CoverGetStatusResponseAenergy     `json:"aenergy,omitempty"` // Optional. Information about the active energy counter (shown if applicable)
	Apower        int                               `json:"apower"`            // Active power in Watts
	Current       int                               `json:"current"`           // Current in Amperes
	CurrentPos    int                               `json:"current_pos"`       // Only present if Cover is calibrated. Represents current position in percent from 0 (fully closed) to 100 (fully open); null if position is unknown
	Errors        []string                          `json:"errors"`            // Only present if an error condition has occurred
	Id            int                               `json:"id"`                // The numeric ID of the Cover component instance
	MoveStartedAt int                               `json:"move_started_at"`   // Only present if Cover is actively moving in either open or close directions. Represents the time at which the movement has begun
	MoveTimeout   int                               `json:"move_timeout"`      // Seconds, only present if Cover is actively moving in either open or close directions. Cover will automatically stop after the timeout expires
	Pf            int                               `json:"pf"`                // Power factor
	PosControl    bool                              `json:"pos_control"`       // False if Cover is not calibrated and only discrete open/close is possible; true if Cover is calibrated and can be commanded to go to arbitrary positions between fully open and fully closed
	Source        string                            `json:"source"`            // Source of the last command
	State         string                            `json:"state"`             // One of open (Cover is fully open), closed (Cover is fully closed), opening (Cover is actively opening), closing (Cover is actively closing), stopped (Cover is not moving, and is neither fully open nor fully closed, or the open/close state is unknown), calibrating (Cover is performing a calibration procedure)
	TargetPos     int                               `json:"target_pos"`        // Only present if Cover is calibrated and is actively moving to a requested position in either open or close directions. Represents the target position in percent from 0 (fully closed) to 100 (fully open); null if target position has been reached or the movement was cancelled
	Temperature   CoverGetStatusResponseTemperature `json:"temperature"`       // Temperature sensor information, only present if a temperature monitor is associated with the Cover instance
	Voltage       int                               `json:"voltage"`           // Voltage in Volts
}

// CoverGetStatusResponseAenergy is the response of aenergy.
type CoverGetStatusResponseAenergy struct {
	ByMinute []int `json:"by_minute"` // Energy consumption by minute (in Milliwatt-hours) for the last three minutes (the lower the index of the element in the array, the closer to the current moment the minute)
	MinuteTs int   `json:"minute_ts"` // Unix timestamp of the first second of the last minute (in UTC)
	Total    int   `json:"total"`     // Total energy consumed in Watt-hours
}

// CoverGetStatusResponseTemperature is the response of temperature.
type CoverGetStatusResponseTemperature struct {
	TC int `json:"tc"` // Temperature in Celsius (null if temperature is out of the measurement range)
	TF int `json:"tf"` // Temperature in Fahrenheit (null if temperature is out of the measurement range)
}

// readResponse reads the response into the given interface.
func (r *CoverGetStatusResponse) readResponse(reader *responseReader) error { //nolint:dupl
	if reader.Response == nil {
		return ErrInvalidResponse
	}
	return reader.Read(r)
}

// GetStatus Get Cover status
func (c CoverClient) GetStatus(args CoverGetStatusRequest) (resp *CoverGetStatusResponse, err error) { //nolint:dupl
	reader := NewResponseReader()

	if err = c.client.rpc.Call("Cover.GetStatus", args, &reader.Response); err != nil {
		return
	}

	resp = &CoverGetStatusResponse{}
	return resp, resp.readResponse(reader)
}

// Getaenergy returns the aenergy value.
func (r *CoverGetStatusResponse) GetAenergy() CoverGetStatusResponseAenergy {
	return r.Aenergy
}

// Getapower returns the apower value.
func (r *CoverGetStatusResponse) GetApower() int {
	return r.Apower
}

// Getcurrent returns the current value.
func (r *CoverGetStatusResponse) GetCurrent() int {
	return r.Current
}

// Getcurrent_pos returns the current_pos value.
func (r *CoverGetStatusResponse) GetCurrentPos() int {
	return r.CurrentPos
}

// Geterrors returns the errors value.
func (r *CoverGetStatusResponse) GetErrors() []string {
	return r.Errors
}

// Getid returns the id value.
func (r *CoverGetStatusResponse) GetId() int {
	return r.Id
}

// Getmove_started_at returns the move_started_at value.
func (r *CoverGetStatusResponse) GetMoveStartedAt() int {
	return r.MoveStartedAt
}

// Getmove_timeout returns the move_timeout value.
func (r *CoverGetStatusResponse) GetMoveTimeout() int {
	return r.MoveTimeout
}

// Getpf returns the pf value.
func (r *CoverGetStatusResponse) GetPf() int {
	return r.Pf
}

// Getpos_control returns the pos_control value.
func (r *CoverGetStatusResponse) GetPosControl() bool {
	return r.PosControl
}

// Getsource returns the source value.
func (r *CoverGetStatusResponse) GetSource() string {
	return r.Source
}

// Getstate returns the state value.
func (r *CoverGetStatusResponse) GetState() string {
	return r.State
}

// Gettarget_pos returns the target_pos value.
func (r *CoverGetStatusResponse) GetTargetPos() int {
	return r.TargetPos
}

// Gettemperature returns the temperature value.
func (r *CoverGetStatusResponse) GetTemperature() CoverGetStatusResponseTemperature {
	return r.Temperature
}

// Getvoltage returns the voltage value.
func (r *CoverGetStatusResponse) GetVoltage() int {
	return r.Voltage
}

// Getby_minute returns the by_minute value.
func (r *CoverGetStatusResponseAenergy) GetByMinute() []int {
	return r.ByMinute
}

// Getminute_ts returns the minute_ts value.
func (r *CoverGetStatusResponseAenergy) GetMinuteTs() int {
	return r.MinuteTs
}

// Gettotal returns the total value.
func (r *CoverGetStatusResponseAenergy) GetTotal() int {
	return r.Total
}

// GettC returns the tC value.
func (r *CoverGetStatusResponseTemperature) GetTC() int {
	return r.TC
}

// GettF returns the tF value.
func (r *CoverGetStatusResponseTemperature) GetTF() int {
	return r.TF
}

/*
 > Command Cover.Open
 > Open Cover
*/

// CoverOpenRequest is the request of Open.
type CoverOpenRequest struct {
	Duration int    `json:"duration,omitempty"` // Optional. If duration is not provided, Cover will fully open, unless it times out because of maxtime_open first. If duration (seconds) is provided, Cover will move in open direction for the specified time. duration must be in range [0.1..maxtime_open]
	Id       string `json:"id"`                 // The numeric ID of the Cover component instance
}

// CoverOpenResponse is the response of Open.
type CoverOpenResponse struct {
}

// readResponse reads the response into the given interface.
func (r *CoverOpenResponse) readResponse(reader *responseReader) error { //nolint:dupl
	if reader.Response == nil {
		return ErrInvalidResponse
	}
	return reader.Read(r)
}

// Open Open Cover
func (c CoverClient) Open(args CoverOpenRequest) (resp *CoverOpenResponse, err error) { //nolint:dupl
	reader := NewResponseReader()

	if err = c.client.rpc.Call("Cover.Open", args, &reader.Response); err != nil {
		return
	}

	resp = &CoverOpenResponse{}
	return resp, resp.readResponse(reader)
}

/*
 > Command Cover.Close
 > Close Cover
*/

// CoverCloseRequest is the request of Close.
type CoverCloseRequest struct {
	Id string `json:"id"` // The numeric ID of the Cover component instance
}

// CoverCloseResponse is the response of Close.
type CoverCloseResponse struct {
}

// readResponse reads the response into the given interface.
func (r *CoverCloseResponse) readResponse(reader *responseReader) error { //nolint:dupl
	if reader.Response == nil {
		return ErrInvalidResponse
	}
	return reader.Read(r)
}

// Close Close Cover
func (c CoverClient) Close(args CoverCloseRequest) (resp *CoverCloseResponse, err error) { //nolint:dupl
	reader := NewResponseReader()

	if err = c.client.rpc.Call("Cover.Close", args, &reader.Response); err != nil {
		return
	}

	resp = &CoverCloseResponse{}
	return resp, resp.readResponse(reader)
}

/*
 > Command Cover.GoToPosition
 > Go to position
*/

// CoverGoToPositionRequest is the request of GoToPosition.
type CoverGoToPositionRequest struct {
	Id  string `json:"id"`  // The numeric ID of the Cover component instance
	Pos int    `json:"pos"` // Required and mutually exclusive (at least one of them pos/rel be provided, but not both at the same time). pos represents target position in %, allowed range [0..100]. If rel is provided, Cover will move to a target_position = current_position + rel. If the value of rel is so big that it results in overshoot (i.e. target_position is beyond fully open / fully closed), target_position will be silently capped to fully open / fully closed
	Rel int    `json:"rel"` // Required and mutually exclusive (at least one of them pos/rel be provided, but not both at the same time). rel represents relative move in %, allowed range [-100..100]
}

// CoverGoToPositionResponse is the response of GoToPosition.
type CoverGoToPositionResponse struct {
}

// readResponse reads the response into the given interface.
func (r *CoverGoToPositionResponse) readResponse(reader *responseReader) error { //nolint:dupl
	if reader.Response == nil {
		return ErrInvalidResponse
	}
	return reader.Read(r)
}

// GoToPosition Go to position
func (c CoverClient) GoToPosition(args CoverGoToPositionRequest) (resp *CoverGoToPositionResponse, err error) { //nolint:dupl
	reader := NewResponseReader()

	if err = c.client.rpc.Call("Cover.GoToPosition", args, &reader.Response); err != nil {
		return
	}

	resp = &CoverGoToPositionResponse{}
	return resp, resp.readResponse(reader)
}
