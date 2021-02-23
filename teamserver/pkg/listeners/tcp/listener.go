package tcp

import (
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
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	msg := make([]byte, 9000)
	n, _ := conn.Read(msg)

	bytes, err := comms.HandleMessage(msg[:n])
	if err != nil {
		// fmt.Println(err)
		return
	}

	conn.Write(bytes)

}
