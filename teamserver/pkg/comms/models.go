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
// Data is one of the following structs
type Packet struct {
	Fingerprint string      `json:"fingerprint"`
	Implant     ImplantInfo `json:"implantinfo"`
	PacketType  int         `json:"packetType"`
	Size        int         `json:"size"`
	Data        string      `json:"data"`
}

// ImplantInfo is the information about the implant being sent
type ImplantInfo struct {
	UUID      string `json:"uuid"`
	PrimaryIP string `json:"primaryip"`
}

// Action is the action that needs to be completed
type Action struct {
	ActionID   int      `json:"actionid"`
	Module     string   `json:"module"`
	ModuleFunc string   `json:"modulefunc"`
	Arguments  []string `json:"arguments"`
}

// ActionResponse is the bots response to an action
type ActionResponse struct {
	ActionID int    `json:"actionid"`
	Response string `json:"response"`
}

// ComError is an error for the action
type ComError struct {
	ActionID int `json:"actionid"`
	ErrorNum int `json:"errornum"`
}
