package comms

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/google/uuid"
)

// RegisterRequestHandler handles the request to register an implant
func RegisterRequestHandler(packet Packet) ([]Packet, error) {
	var allPackets []Packet
	register := RegistrationRequest{}

	err := json.Unmarshal([]byte(packet.Data), &register)
	if err != nil {
		return nil, err
	}

	h := sha1.New()
	h.Write([]byte(register.ImplantName))
	h.Write([]byte(register.ImplantVersion))
	tmp := h.Sum(nil)
	id := fmt.Sprintf("%x", tmp)

	supportModulesStr, err := ModuleCheckHandler(register.SupportedModules)
	if err != nil {
		return nil, err
	}

	implantType := model.ImplantType{
		UUID:           id,
		ImplantName:    register.ImplantName,
		ImplantVersion: register.ImplantVersion,
	}

	implant := model.Implant{
		UUID:             uuid.New().String(),
		UUIDImplantType:  id,
		PrimaryIP:        register.IP,
		Hostname:         register.Hostname,
		MAC:              register.MAC,
		ImplantOS:        register.OS,
		OtherIPs:         register.OtherIPs,
		SupportedModules: supportModulesStr,
	}

	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	success := true
	_, err = d.GetImplantTypeById(implantType.UUID)
	if err != nil {
		check, err := d.InsertImplantType(implantType)
		if err != nil || check == false {
			success = false
		}
	}

	_, err = d.GetImplantById(implant.UUID)
	if err != nil {
		check, err := d.InsertImplant(implant)
		if err != nil || check == false {
			success = false
		}
	}

	if success {
		resp := RegistrationResponse{UUID: implant.UUID}
		bytes, _ := json.Marshal(resp)

		respPacket := Packet{
			Fingerprint: "fingerprint",
			Implant: ImplantInfo{
				UUID:      implant.UUID,
				PrimaryIP: implant.PrimaryIP,
			},
			PacketType: RegisterResponse,
			NumLeft:    0,
			Data:       string(bytes),
		}
		allPackets = append(allPackets, respPacket)
	} else if !success {
		commErr := ComError{
			ActionID: "-1",
			ErrorNum: ErrorHandler(err),
		}
		bytes, _ := json.Marshal(commErr)

		errPacket := Packet{
			Fingerprint: "fingerprint",
			Implant: ImplantInfo{
				UUID:      "",
				PrimaryIP: register.IP,
			},
			PacketType: ComErrorEnum,
			NumLeft:    0,
			Data:       string(bytes),
		}
		allPackets = append(allPackets, errPacket)
	}

	return allPackets, nil
}

// ModuleCheckHandler inserts any new modules and returns SupportedModule string slice
func ModuleCheckHandler(newModules []ModuleInfo) ([]string, error) {
	moduleStr := []string{}
	if len(newModules) == 0 {
		return nil, fmt.Errorf("no new modules")
	}

	// Open the database connection
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	// Check all of the module functions with the database
	for i := 0; i < len(newModules); i++ {
		name := newModules[i].ModuleName
		moduleStr = append(moduleStr, name)
		_, err := d.GetModuleByName(name)

		// If the module is not in the database add it
		if err != nil {

			var moduleFuncs []model.ModulesFuncs
			moduleFuncIds := []string{}
			for j := 0; j < len(newModules[i].ModuleFuncs); j++ {
				newFunc, _ := GenerateModuleFunc(newModules[i].ModuleFuncs[j])
				moduleFuncs = append(moduleFuncs, newFunc)
				moduleFuncIds = append(moduleFuncIds, newFunc.UUID)
			}

			// Insert the module after running checks
			module := model.Modules{
				ModuleName:    name,
				ModuleDesc:    newModules[i].ModuleDesc,
				ModuleFuncIds: moduleFuncIds,
			}
			check, err := d.InsertModule(module)
			if err != nil {
				return nil, fmt.Errorf("MissingRequiredData")
			}

			if check == true {
				for j := 0; j < len(moduleFuncs); j++ {
					_, err = d.InsertModuleFunc(moduleFuncs[i])
					if err != nil {
						return nil, fmt.Errorf("MissingRequiredData")
					}
				}
			}

		}

	}

	return moduleStr, nil
}

// GenerateModuleFunc converts ModuleFuncInfo into a ModuleFunc struct for the database
func GenerateModuleFunc(moduleFunc ModuleFuncInfo) (model.ModulesFuncs, error) {
	newModuleFunc := model.ModulesFuncs{
		UUID:            uuid.New().String(),
		ModuleFuncName:  moduleFunc.ModuleFuncName,
		ModuleFuncDesc:  moduleFunc.ModuleFuncDesc,
		NumOfParameters: moduleFunc.ParamNum,
		ParameterTypes:  moduleFunc.ParamTypes,
		ParameterNames:  moduleFunc.ParamNames,
	}

	return newModuleFunc, nil
}

func CreateErrorPacket(implant ImplantInfo, commErr ComError) (Packet, error) {

	packet := Packet{
		Fingerprint: "fingerprint",
		Implant: ImplantInfo{
			UUID:      "",
			PrimaryIP: register.IP,
		},
		PacketType: ComErrorEnum,
		NumLeft:    0,
		Data:       string(bytes),
	}

	return
}
