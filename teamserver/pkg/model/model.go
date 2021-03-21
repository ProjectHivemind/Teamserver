package model

type ImplantType struct {
	UUID           string `json:"uuid"`
	ImplantName    string `json:"implantname"`
	ImplantVersion string `json:"implantversion"`
}

type Implant struct {
	UUID             string   `json:"uuid"`
	UUIDImplantType  string   `json:"uuidimplanttype"`
	PrimaryIP        string   `json:"primaryip"`
	Hostname         string   `json:"hostname"`
	MAC              string   `json:"mac"`
	ImplantOS        string   `json:"implantos"`
	OtherIPs         []string `json:"otherips"`
	SupportedModules []string `json:"supportedmodules"`
}

type CallBack struct {
	UUIDImplant string `json:"uuidmplant"`
	FirstCall   string `json:"firstcall"`
	LastCall    string `json:"lastcall"`
}

type ImplantWithCallbacks struct {
	Implant  Implant  `json:"implant"`
	CallBack CallBack `json:"callback"`
}

type ModulesFuncs struct {
	UUID            string   `json:"uuid"`
	ModuleFuncName  string   `json:"modulefuncname"`
	ModuleFuncDesc  string   `json:"modulefuncdesc"`
	NumOfParameters int      `json:"numofparamenters"`
	ParameterTypes  []string `json:"parametertypes"`
	ParameterNames  []string `json:"parameternames"`
}

type Modules struct {
	ModuleName    string   `json:"modulename"`
	ModuleDesc    string   `json:"moduledesc"`
	ModuleFuncIds []string `json:"modulefuncnames"`
}

type ParamType struct {
	TypeName     string   `json:"typename"`
	IsCombo      bool     `json:"iscombo"`
	ComboOptions []string `json:"combooptions"`
}

type Groups struct {
	UUID      string   `json:"uuid"`
	GroupName string   `json:"groupname"`
	Implants  []string `json:"implants"`
}

type StagedActions struct {
	Id            string `json:"id"`
	UUIDofAction  string `json:"uuidofaction"`
	UUIDofImplant string `json:"uuidofimplant"`
	TimeStaged    string `json:"timestaged"`
}

type StoredActions struct {
	UUID        string   `json:"uuid"`
	ModuleToRun string   `json:"moduletorun"`
	ModuleFunc  string   `json:"modulefunc"`
	Arguments   []string `json:"arguments"`
}

type ExecutedActions struct {
	Id             string `json:"id"`
	UUIDofAction   string `json:"uuidofaction"`
	TimeSent       string `json:"timesent"`
	TimeRan        string `json:"timeran"`
	Successful     bool   `json:"successful"`
	ActionResponse string `json:"actionresponse"`
}

type Operators struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Permission int    `json:"permission"`
}
