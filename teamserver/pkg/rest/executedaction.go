package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
	"github.com/gorilla/mux"
)

func getExecutedActions(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	actions, err := d.AllExecutedActions()
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(actions)
	}
}

func getExecutedAction(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	id := mux.Vars(r)["id"]
	action, err := d.GetExecutedActionById(id)

	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(action)
	}
}
