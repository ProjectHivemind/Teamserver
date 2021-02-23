package tcp

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/comms"
)

// StartListener start a tcp listening channel on that port
func StartListener(port string) {

	serverAddr := "0.0.0.0:" + port
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Starting TCP listener on %s\n", serverAddr)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	msg := make([]byte, 8096)
	n, _ := conn.Read(msg)

	allPackets, err := comms.HandleMessage(msg[:n])
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(allPackets) > 0 {
		for i := 0; i < len(allPackets); i++ {
			tmp, _ := json.Marshal(allPackets[i])
			conn.Write(tmp)
		}
	}

}
