package comms

import (
	"fmt"
)

// HandleMessage takes the packet and decides what to do with it
func HandleMessage(packet Packet) ([]Packet, error) {
	switch packet.PacketType {
	case ActionRequestEnum:
		return ActionRequestHandler(packet)
	case ActionResponseEnum:
		return ActionResponseHandler(packet)
	case RegisterRequestEnum:
		return RegisterRequestHandler(packet)
	default:
		break
	}
	return nil, fmt.Errorf("no a valid PacketType")
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
