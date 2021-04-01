package comms

// Enum of PacketType
const (
	ComErrorEnum        = -1
	NoActionEnum        = 0
	ActionRequestEnum   = 1
	ActionEnum          = 2
	ActionResponseEnum  = 3
	RegisterRequestEnum = 4
	RegisterResponse    = 5
)

// Enum of Errors
const (
	NotRegisteredEnum       = -1
	UnknownModuleEnum       = -2
	UnknownModuleFuncEnum   = -3
	MissingRequiredDataEnum = -4
	ModuleTimeout           = -5
	DuplicateRegistration   = -6
	UnknownErrorEnum        = -7
)

// Packet is the data that is recieved from on implant
// NumLeft is how many more packets are coming in
// Data is one of the following structs
type Packet struct {
	Fingerprint string      `json:"fingerprint"`
	Implant     ImplantInfo `json:"implant_info"`
	PacketType  int         `json:"packet_type"`
	NumLeft     int         `json:"num_left"`
	Data        string      `json:"data"`
}

// ImplantInfo is the information about the implant being sent
type ImplantInfo struct {
	UUID      string `json:"uuid"`
	PrimaryIP string `json:"primary_ip"`
}

// Action is the action that needs to be completed
type Action struct {
	ActionID   string `json:"action_id"`
	Module     string `json:"module"`
	ModuleFunc string `json:"module_func"`
	Arguments  string `json:"arguments"`
}

// ActionResponse is the implant's response to an action
type ActionResponse struct {
	ActionID string `json:"action_id"`
	Response string `json:"response"`
}

// RegistrationRequest has all the info needed for an implant to register
type RegistrationRequest struct {
	IP               string       `json:"ip"`
	ImplantName      string       `json:"implant_name"`
	ImplantVersion   string       `json:"implant_version"`
	Hostname         string       `json:"hostname"`
	MAC              string       `json:"mac"`
	OtherIPs         []string     `json:"other_ips"`
	OS               string       `json:"os"`
	SupportedModules []ModuleInfo `json:"supported_modules"`
}

// RegistrationResponse gives the uuid to the bot
type RegistrationResponse struct {
	UUID string `json:"uuid"`
}

// ModuleInfo has the name of the module from the implant with funcs
type ModuleInfo struct {
	ModuleName  string           `json:"module_name"`
	ModuleDesc  string           `json:"module_desc"`
	ModuleFuncs []ModuleFuncInfo `json:"module_funcs"`
}

// ModuleFuncInfo has all the information about a module function
type ModuleFuncInfo struct {
	ModuleFuncName string   `json:"module_func_name"`
	ModuleFuncDesc string   `json:"module_func_desc"`
	ParamNum       int      `json:"param_num"`
	ParamNames     []string `json:"param_names"`
	ParamTypes     []string `json:"param_types"`
}

// ComError is an error for the action
type ComError struct {
	ActionID string `json:"action_id"`
	ErrorNum int    `json:"error_num"`
}
