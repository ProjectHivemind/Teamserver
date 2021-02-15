package model

import "time"

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

type ModulesFuncs struct {
	ModuleFuncName  string   `json:"modulefuncname"`
	NumOfParameters int      `json:"numofparamenters"`
	ParameterTypes  []string `json:"parametertypes"`
	ParameterNames  []string `json:"parameternames"`
}

type Modules struct {
	ModuleName      string   `json:"modulename"`
	ModuleFuncNames []string `json:"modulefuncnames"`
}

type Groups struct {
	UUID      string   `json:"uuid"`
	GroupName string   `json:"groupname"`
	Implants  []string `json:"implants"`
}

type StagedActions struct {
	Id            int       `json:"id"`
	UUIDofAction  string    `json:"uuidofaction"`
	UUIDofImplant string    `json:"uuidofimplant"`
	TimeStaged    time.Time `json:"timestaged"`
}

type StoredActions struct {
	UUID        string   `json:"uuid"`
	ModuleToRun string   `json:"moduletorun"`
	ModuleFunc  string   `json:"modulefunc"`
	Arguments   []string `json:"arguments"`
}

type ExecutedActions struct {
	Id             int    `json:"id"`
	UUIDofAction   string `json:"uuidofaction"`
	TimeRan        string `json:"timeran"`
	Successful     bool   `json:"successful"`
	ActionResponse string `json:"actionresponse"`
}

type Operators struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Permission int    `json:"permission"`
}
