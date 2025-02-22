package tcp

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/comms"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
)

var db crud.DatabaseModel

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

	// Connect to the database for this listener
	db.Open()
	defer db.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

// handleConnection is used to handle the TCP connection.
func handleConnection(conn net.Conn) {
	defer conn.Close()
	msg := make([]byte, 5000) // Needs to be able to accept large registrations, may need to be bigger or done differently
	reader := bufio.NewReader(conn)
	n, err := io.ReadFull(reader, msg)
	if err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF {
			fmt.Println("Read error:", err)
		}
	}

	bytes, err := comms.HandleMessage(msg[:n], db)
	if err != nil {
		return
	}

	conn.Write(bytes)

}
