package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func getStoredActions(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	actions, err := d.AllStoredActions()
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(actions)
	}
}

func getStoredAction(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	id := mux.Vars(r)["id"]
	action, err := d.GetStoredActionById(id)

	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(action)
	}
}

func createStoredAction(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, err)
	}

	var action model.StoredActions
	err = json.Unmarshal(reqBody, &action)
	if err != nil {
		fmt.Fprint(w, err)
	}
	action.UUID = uuid.New().String()

	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	_, err = d.InsertStoredAction(action)
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(action)
	}
}

func deleteStoredAction(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	id := mux.Vars(r)["id"]
	_, err := d.DeleteStoredAction(id)

	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		fmt.Fprint(w, "deleted")
	}
}
