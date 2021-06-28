// simplehttps package
package simplehttps

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/comms"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
)

var db crud.DatabaseModel

// StartListener start a tcp listening channel on that port
func StartListener(port, crtPath, keyPath string) {

	// Connect to the database for this listener
	db.Open()
	defer db.Close()

	http.HandleFunc("/simplehttps", handleConnection)

	fmt.Printf("Starting SimpleHTTPS listener on 0.0.0.0:%s\n", port)
	http.ListenAndServeTLS(":"+port, crtPath, keyPath, nil)
}

// handleConnection handles any of the HTTP connections
func handleConnection(w http.ResponseWriter, req *http.Request) {

	// If not a POST or GET return nothing
	if req.Method == http.MethodGet || req.Method == http.MethodPost {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return
		}
		respBytes, err := comms.HandleMessage(data, db)
		if err != nil {
			return
		}

		w.Write(respBytes)
	}

}
