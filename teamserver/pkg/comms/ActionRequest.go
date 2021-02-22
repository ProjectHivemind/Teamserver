package comms

import (
	"encoding/json"
	"fmt"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
)

// ActionRequestHandler checks to see if there are any actions queued for the
// given implant.
func ActionRequestHandler(packet Packet) ([]Packet, error) {
	var allPackets []Packet
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	stagedActions, err := d.GetStagedActionByImplant(packet.Implant.UUID)
	if err != nil {
		return nil, fmt.Errorf("no actions")
	}

	packetCtr := len(stagedActions)
	for i := 0; i < len(stagedActions); i++ {
		if packetCtr == 0 {
			break
		}

		action, err := GenerateAction(stagedActions[i])
		if err != nil {
			packetCtr--
			continue
		}

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

		bytes, _ := json.Marshal(action)

		actionPacket := Packet{
			Fingerprint: "fingerprint",
			Implant:     packet.Implant,
			PacketType:  ActionEnum,
			NumLeft:     packetCtr,
			Data:        string(bytes),
		}

		packetCtr--
		allPackets = append(allPackets, actionPacket)
	}

	return allPackets, nil
}

// GenerateAction given a staged action, it will return a sendable action Packet
func GenerateAction(stagedAction model.StagedActions) (Action, error) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	var action Action
	args := make(map[string]string, 0)

	storedAction, err := d.GetStoredActionById(stagedAction.UUIDofAction)
	if err != nil {
		return action, err
	}

	// TODO: ADD ERROR CHECKING
	module, _ := d.GetModuleByName(storedAction.ModuleFunc)
	var argStr []string
	for i := 0; i < len(module.ModuleFuncIds); i++ {
		moduleFunc, _ := d.GetModuleFuncById(module.ModuleFuncIds[i])
		if moduleFunc.ModuleFuncName == storedAction.ModuleFunc {
			argStr = moduleFunc.ParameterNames
			break
		}
	}

	for i := 0; i < len(argStr); i++ {
		args[argStr[i]] = storedAction.Arguments[i]
	}
	// Json Marshal the arguments
	bytes, _ := json.Marshal(args)

	action = Action{
		ActionID:   stagedAction.Id,
		Module:     storedAction.ModuleToRun,
		ModuleFunc: storedAction.ModuleFunc,
		Arguments:  string(bytes),
	}

	return action, nil
}
