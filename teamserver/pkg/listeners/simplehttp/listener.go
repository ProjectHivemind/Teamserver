// simplehttp package
package simplehttp

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/comms"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
)

var db crud.DatabaseModel

// StartListener start a tcp listening channel on that port
func StartListener(port, url string) {

	// Connect to the database for this listener
	db.Open()
	defer db.Close()

	http.HandleFunc(url, handleConnection)

	fmt.Printf("Starting SimpleHTTP listener on http://0.0.0.0:%s%s\n", port, url)
	http.ListenAndServe(":"+port, nil)
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
