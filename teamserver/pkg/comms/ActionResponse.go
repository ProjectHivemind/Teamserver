package comms

import (
	"encoding/json"
	"time"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
)

// ActionResponseHandler updates the database with any responses from the
// implant's action
func ActionResponseHandler(packet Packet) ([]Packet, error) {
	var allPackets []Packet
	actionResp := ActionResponse{}

	err := json.Unmarshal([]byte(packet.Data), &actionResp)
	if err != nil {
		return nil, err
	}

	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	_, err = d.GetExecutedActionById(actionResp.ActionID)
	if err != nil {
		commErr := ComError{
			ActionID: "-1",
			ErrorNum: ErrorHandler(err),
		}
		bytes, _ := json.Marshal(commErr)

		errPacket := Packet{
			Fingerprint: "fingerprint",
			Implant:     packet.Implant,
			PacketType:  ComErrorEnum,
			NumLeft:     0,
			Data:        string(bytes),
		}

		allPackets = append(allPackets, errPacket)
		return allPackets, err
	}

	d.UpdateExecutedActionResponse(actionResp.ActionID, actionResp.Response)
	d.UpdateExecutedActionSuccessful(actionResp.ActionID, true)
	d.UpdateExecutedActionTimeRan(actionResp.ActionID, time.Now().Format(crud.TimeStamp))

	return allPackets, nil
}
