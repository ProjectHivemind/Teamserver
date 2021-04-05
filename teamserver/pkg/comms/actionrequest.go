package comms

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/plugins"
)

// ActionRequestHandler checks to see if there are any actions queued for the
// given implant.
func ActionRequestHandler(packet Packet) ([]Packet, error) {
	var allPackets []Packet

	// Check if implant has a uuid
	if packet.Implant.UUID == "" {
		return nil, fmt.Errorf("not registered")
	}

	// Update that the implant called back
	_, err := d.UpdateCallBackTime(packet.Implant.UUID, time.Now().Format(crud.TimeStamp))
	if err != nil {
		return nil, fmt.Errorf("not registered")
	}

	// -------- This is for the pwnboard and sawmill plugins ----------
	implant, _ := d.GetImplantById(packet.Implant.UUID)
	implantType, _ := d.GetImplantTypeById(implant.UUIDImplantType)
	plugins.UpdatepwnBoard(packet.Implant.PrimaryIP, implantType.ImplantName+"-"+implantType.ImplantVersion)
	// -----------------------------------------------------------------

	stagedActions, err := d.GetStagedActionByImplant(packet.Implant.UUID)
	if err != nil {
		return nil, err
	}

	// If there are not actions, send nothing
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

	for i := 0; i < len(stagedActions); i++ {

		action, err := generateAction(stagedActions[i])
		if err != nil {
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
		executed := model.ExecutedAction{
			Id:            stagedActions[i].Id,
			UUIDofImplant: stagedActions[i].UUIDofImplant,
			UUIDofAction:  stagedActions[i].UUIDofAction,
			TimeSent:      time.Now().Format(crud.TimeStamp),
		}
		d.InsertExecutedAction(executed)
	}

	return allPackets, nil
}

// generateAction given a staged action, it will return a sendable action Packet
func generateAction(stagedAction model.StagedAction) (Action, error) {
	var action Action

	// Get the StoredAction from the database
	storedAction, err := d.GetStoredActionById(stagedAction.UUIDofAction)
	if err != nil {
		return action, err
	}

	// Get the Module that is being called
	module, err := d.GetModuleByName(storedAction.ModuleToRun)
	if err != nil {
		return action, fmt.Errorf("unknown module")
	}

	// Checks the modulefunc is there
	// TODO: Decide if this should be taken out.
	for i := 0; i < len(module.ModuleFuncIds); i++ {
		_, err := d.GetModuleFuncById(module.ModuleFuncIds[i])
		if err != nil {
			return action, fmt.Errorf("unknown modulefunc")
		}
	}

	action = Action{
		ActionID:   stagedAction.Id,
		Module:     storedAction.ModuleToRun,
		ModuleFunc: storedAction.ModuleFunc,
		Arguments:  storedAction.Arguments,
	}

	return action, nil
}
