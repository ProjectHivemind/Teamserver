package model

type ImplantType struct {
	UUID           string `json:"uuid"`
	ImplantName    string `json:"implant_name"`
	ImplantVersion string `json:"implant_version"`
}

type Implant struct {
	UUID             string   `json:"uuid"`
	UUIDImplantType  string   `json:"uuid_implant_type"`
	PrimaryIP        string   `json:"primary_ip"`
	Hostname         string   `json:"hostname"`
	MAC              string   `json:"mac"`
	ImplantOS        string   `json:"implant_os"`
	OtherIPs         []string `json:"other_ips"`
	SupportedModules []string `json:"supported_modules"`
}

type CallBack struct {
	UUIDImplant string `json:"uuid_implant"`
	FirstCall   string `json:"first_call"`
	LastCall    string `json:"last_call"`
}

type ImplantWithCallback struct {
	Implant     Implant     `json:"implant"`
	ImplantType ImplantType `json:"implant_type"`
	CallBack    CallBack    `json:"callback"`
}

type ModulesFunc struct {
	UUID           string   `json:"uuid"`
	ModuleFuncName string   `json:"module_func_name"`
	ModuleFuncDesc string   `json:"module_func_desc"`
	NumOfParams    int      `json:"num_of_params"`
	ParamTypes     []string `json:"param_types"`
	ParamNames     []string `json:"param_names"`
}

type Module struct {
	ModuleName    string   `json:"module_name"`
	ModuleDesc    string   `json:"module_desc"`
	ModuleFuncIds []string `json:"module_func_names"`
}

type ParamType struct {
	TypeName     string   `json:"type_name"`
	IsCombo      bool     `json:"is_combo"`
	ComboOptions []string `json:"combo_options"`
}

type Group struct {
	UUID      string   `json:"uuid"`
	GroupName string   `json:"group_name"`
	Implants  []string `json:"implants"`
}

type StagedAction struct {
	Id            string `json:"id"`
	UUIDofAction  string `json:"uuid_of_action"`
	UUIDofImplant string `json:"uuid_of_implant"`
	TimeStaged    string `json:"time_staged"`
}

type StagedActionsFrontend struct {
	Implant      Implant      `json:"implant"`
	StagedAction StagedAction `json:"staged_action"`
	StoredAction StoredAction `json:"stored_action"`
}

type StoredAction struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	ModuleToRun string `json:"module_to_run"`
	ModuleFunc  string `json:"module_func"`
	Arguments   string `json:"arguments"`
}

type ExecutedAction struct {
	Id             string `json:"id"`
	UUIDofImplant  string `json:"uuid_of_implant"`
	UUIDofAction   string `json:"uuid_of_action"`
	TimeSent       string `json:"time_sent"`
	TimeRan        string `json:"time_ran"`
	Successful     bool   `json:"successful"`
	ActionResponse string `json:"action_response"`
}

type ExecutedActionsFrontend struct {
	Implant        Implant        `json:"implant"`
	ExecutedAction ExecutedAction `json:"executed_action"`
	StoredAction   StoredAction   `json:"stored_action"`
}

type Operator struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Permission int    `json:"permission"`
}

type Session struct {
	SessionToken string `json:"session_token"`
	Username     string `json:"username"`
	ExpTime      string `json:"exptime"`
}
