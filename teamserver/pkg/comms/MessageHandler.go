package comms

import (
	"encoding/json"
	"fmt"
)

// HandleMessage takes the packet and decides what to do with it
func HandleMessage(packetBytes []byte) ([]Packet, error) {
	var packet Packet
	if err := json.Unmarshal(packetBytes, &packet); err != nil {
		fmt.Print(err)
		return nil, fmt.Errorf("error parsing packet")
	}

	switch packet.PacketType {
	case ActionRequestEnum:
		return ActionRequestHandler(packet)
	case ActionResponseEnum:
		return ActionResponseHandler(packet)
	case RegisterRequestEnum:
		return RegisterRequestHandler(packet)
	default:
		return nil, fmt.Errorf("no a valid PacketType")
	}

	// get up to send bytes back by looping through the packets

}

// Creates an error packet with the implant info and commerror struct
func CreateErrorPacket(implant ImplantInfo, commErr ComError) (Packet, error) {

	bytes, _ := json.Marshal(commErr)

	packet := Packet{
		Fingerprint: "fingerprint",
		Implant:     implant,
		PacketType:  ComErrorEnum,
		NumLeft:     0,
		Data:        string(bytes),
	}

	return packet, nil
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
