# Creating a Transport
Transports in Hivemind handle receiving data from an implant or C2 and send the bytes to the `comms.HandleMessage`. It will then need to be able to handle sending back to the receiver as well. The Hivemind specific action that needs to be taken is connection to the database. Check the template and example below.

All of the transports are located in `pkg/listeners/...`. This is where any of the listeners will be and anything you create should go in its own folder as to obtain its own package.

### Important Note
Transports only handle the bytes that are sent. They do not translate to anything else, but just extract the bytes out and send them to the `comms.HandleMessage`

## Template Transport
Place this in `pkg/listeners/<transport_name>/listener.go`
``` Golang
package pkg_name

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
   // Transport Logic Goes here
   // Note: should run handleConnections in a go routine so that it can take multiple connections
   //       This might change though depending on the type of connection your listener is. 
}

func handleConnection(conn net.Conn) {
    // Handles whatever connection type that you set up.
    // It then sends it to the HandleMessage function and will return what it needs to
    // send back the the client.
    bytes, err := comms.HandleMessage(msg[:n], db)
}
```

## Example TCP Transport
Located in `pkg/listeners/tcp`
``` Golang
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
```