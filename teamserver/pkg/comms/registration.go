package comms

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/google/uuid"
)

var registerModLock sync.Mutex = sync.Mutex{}

// RegisterRequestHandler handles the request to register an implant
func RegisterRequestHandler(packet Packet) ([]Packet, error) {
	var allPackets []Packet

	// Translate the data to Registration Request
	register := RegistrationRequest{}

	err := json.Unmarshal([]byte(packet.Data), &register)
	if err != nil {
		return nil, err
	}

	// Get the hash of the implant type
	h := sha1.New()
	h.Write([]byte(register.ImplantName))
	h.Write([]byte(register.ImplantVersion))
	tmp := h.Sum(nil)
	id := fmt.Sprintf("%x", tmp)

	// Check if implant is already registered
	check, tmpUUID, _ := checkDuplicateRegistration(id, register.IP)
	if check {
		// DuplicateRegistration Error
		errPacket, _ := CreateErrorPacket(
			ImplantInfo{UUID: tmpUUID, PrimaryIP: register.IP},
			ComError{ActionID: "", ErrorNum: DuplicateRegistration})

		allPackets = append(allPackets, errPacket)
		return allPackets, nil
	}

	// Checks and add Modules as needed
	supportModulesStr, err := moduleCheckHandler(register.SupportedModules)
	if err != nil {
		return nil, err
	}

	// Create the Implant Modules
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

	success, err := insertImplantInfo(implantType, implant)

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
		errPacket, _ := CreateErrorPacket(
			ImplantInfo{UUID: "", PrimaryIP: register.IP},
			ComError{ActionID: "-1", ErrorNum: ErrorHandler(err)})

		allPackets = append(allPackets, errPacket)
	}

	return allPackets, nil
}

// checkDuplicateRegistration checks to see if there is already an entry with that ImplantType and IP
func checkDuplicateRegistration(id string, ip string) (bool, string, error) {
	_, err := d.GetImplantTypeById(id)
	if err != nil {
		return false, "", err
	}

	implants, err := d.GetImplantByIp(ip)
	if err != nil {
		return false, "", err
	}

	for i := 0; i < len(implants); i++ {
		if implants[i].UUIDImplantType == id {
			return true, implants[i].UUID, nil
		}
	}

	return false, "", nil
}

// moduleCheckHandler inserts any new modules and returns SupportedModule string slice
func moduleCheckHandler(newModules []ModuleInfo) ([]string, error) {
	moduleStr := []string{}
	if len(newModules) == 0 {
		return nil, fmt.Errorf("no new modules")
	}

	// Check all of the module functions with the database
	for i := 0; i < len(newModules); i++ {
		name := newModules[i].ModuleName
		moduleStr = append(moduleStr, name)
		registerModLock.Lock()
		_, err := d.GetModuleByName(name)

		// If the module is not in the database add it
		if err != nil {

			var moduleFuncs []model.ModulesFunc
			moduleFuncIds := []string{}
			for j := 0; j < len(newModules[i].ModuleFuncs); j++ {
				newFunc, _ := generateModuleFunc(newModules[i].ModuleFuncs[j])
				moduleFuncs = append(moduleFuncs, newFunc)
				moduleFuncIds = append(moduleFuncIds, newFunc.UUID)
			}

			// Insert the module after running checks
			module := model.Module{
				ModuleName:    name,
				ModuleDesc:    newModules[i].ModuleDesc,
				ModuleFuncIds: moduleFuncIds,
			}
			check, err := d.InsertModule(module)
			if err != nil {
				registerModLock.Unlock()
				return nil, fmt.Errorf("MissingRequiredData")
			}

			if check == true {
				for j := 0; j < len(moduleFuncs); j++ {
					_, err = d.InsertModuleFunc(moduleFuncs[j])
					if err != nil {
						registerModLock.Unlock()
						return nil, fmt.Errorf("MissingRequiredData")
					}
				}
			}

		}
		registerModLock.Unlock()
	}

	return moduleStr, nil
}

// generateModuleFunc converts ModuleFuncInfo into a ModuleFunc struct for the database
func generateModuleFunc(moduleFunc ModuleFuncInfo) (model.ModulesFunc, error) {
	newModuleFunc := model.ModulesFunc{
		UUID:           uuid.New().String(),
		ModuleFuncName: moduleFunc.ModuleFuncName,
		ModuleFuncDesc: moduleFunc.ModuleFuncDesc,
		NumOfParams:    moduleFunc.ParamNum,
		ParamTypes:     moduleFunc.ParamTypes,
		ParamNames:     moduleFunc.ParamNames,
	}

	return newModuleFunc, nil
}

// insertImplantInfo insert implant info as needed
func insertImplantInfo(implantType model.ImplantType, implant model.Implant) (bool, error) {
	success := true
	_, err := d.GetImplantTypeById(implantType.UUID)
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

	callBack := model.CallBack{
		UUIDImplant: implant.UUID,
		FirstCall:   time.Now().Format(crud.TimeStamp),
		LastCall:    time.Now().Format(crud.TimeStamp),
	}
	_, err = d.InsertCallBack(callBack)
	if err != nil {
		success = false
	}

	return success, err
}
