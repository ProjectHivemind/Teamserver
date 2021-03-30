package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func getStagedActions(w http.ResponseWriter, r *http.Request) {
	actions, err := d.AllStagedActions()
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(actions)
	}
}

func getStagedActionsFrontend(w http.ResponseWriter, r *http.Request) {
	actions, err := d.AllStagedActionsFrontend()
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(actions)
	}
}

func getStagedAction(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	action, err := d.GetStagedActionById(id)

	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(action)
	}
}

func createStagedAction(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, err)
	}

	var action model.StagedActions
	err = json.Unmarshal(reqBody, &action)
	if err != nil {
		fmt.Fprint(w, err)
	}
	action.Id = uuid.New().String()

	_, err = d.InsertStagedAction(action)
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(action)
	}
}

func deleteStagedAction(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := d.DeleteStoredAction(id)

	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		fmt.Fprint(w, "deleted")
	}
}
