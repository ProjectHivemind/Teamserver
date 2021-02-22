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
	Fingerprint string      `json:"fingerPrint"`
	Implant     ImplantInfo `json:"implantInfo"`
	PacketType  int         `json:"packetType"`
	NumLeft     int         `json:"numLeft"`
	Data        string      `json:"data"`
}

// ImplantInfo is the information about the implant being sent
type ImplantInfo struct {
	UUID      string `json:"uuid"`
	PrimaryIP string `json:"primaryIP"`
}

// Action is the action that needs to be completed
type Action struct {
	ActionID   string   `json:"actionId"`
	Module     string   `json:"module"`
	ModuleFunc string   `json:"moduleFunc"`
	Arguments  []string `json:"arguments"`
}

// ActionResponse is the implant's response to an action
type ActionResponse struct {
	ActionID string `json:"actionId"`
	Response string `json:"response"`
}

// RegistrationRequest has all the info needed for an implant to register
type RegistrationRequest struct {
	IP               string       `json:"IP"`
	ImplantName      string       `json:"implantName"`
	ImplantVersion   string       `json:"implantVersion"`
	Hostname         string       `json:"hostname"`
	MAC              string       `json:"MAC"`
	OtherIPs         []string     `json:"otherips"`
	OS               string       `json:"OS"`
	SupportedModules []ModuleInfo `json:"supportedModules"`
}

// RegistrationResponse gives the uuid to the bot
type RegistrationResponse struct {
	UUID string `json:"uuid"`
}

// ModuleInfo has the name of the module from the implant with funcs
type ModuleInfo struct {
	ModuleName  string           `json:"moduleName"`
	ModuleDesc  string           `json:"moduleDesc"`
	ModuleFuncs []ModuleFuncInfo `json:"moduleFuncs"`
}

// ModuleFuncInfo has all the information about a module function
type ModuleFuncInfo struct {
	ModuleFuncName string   `json:"moduleFuncName"`
	ModuleFuncDesc string   `json:"moduleFuncDesc"`
	ParamNum       int      `json:"paramNum"`
	ParamNames     []string `json:"paramNames"`
	ParamTypes     []string `json:"paramTypes"`
}

// ComError is an error for the action
type ComError struct {
	ActionID string `json:"actionId"`
	ErrorNum int    `json:"errorNum"`
}
