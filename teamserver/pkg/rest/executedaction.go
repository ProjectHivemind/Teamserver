package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getExecutedActions(w http.ResponseWriter, r *http.Request) {
	actions, err := d.AllExecutedActions()
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(actions)
	}
}

func getExecutedActionsFrontend(w http.ResponseWriter, r *http.Request) {
	actions, err := d.AllExecutedActionsFrontend()
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(actions)
	}
}

func getExecutedAction(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	action, err := d.GetExecutedActionById(id)

	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(action)
	}
}
