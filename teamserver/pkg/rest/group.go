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

func getGroups(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	groups, err := d.AllGroups()
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(groups)
	}
}

func getGroup(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	params := mux.Vars(r)
	group, err := d.GetGroupById(params["id"])

	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(group)
	}
}

func createGroup(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, err)
	}

	var group model.Groups
	err = json.Unmarshal(reqBody, &group)
	if err != nil {
		fmt.Fprint(w, err)
	}
	group.UUID = uuid.New().String()

	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	_, err = d.InsertGroup(group)
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(group)
	}
}

func deleteGroup(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	id := mux.Vars(r)["id"]
	err := d.RemoveGroupById(id)

	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		fmt.Fprint(w, "deleted")
	}
}

func addToGroup(w http.ResponseWriter, r *http.Request) {
	// var d crud.DatabaseModel
	// d.Open()
	// defer d.Close()

	// id := mux.Vars(r)["id"]
	// check, err := d.AddUUIDToGroup(id)

	// if err != nil {
	// 	fmt.Fprint(w, GENERAL_ERROR)
	// } else {
	// 	json.NewEncoder(w).Encode(check)
	// }
}

func removeFromGroup(w http.ResponseWriter, r *http.Request) {
	// var d crud.DatabaseModel
	// d.Open()
	// defer d.Close()

	// params := mux.Vars(r)
	// check, err := d.AddUUIDToGroup(params["id"])

	// if err != nil {
	// 	fmt.Fprint(w, GENERAL_ERROR)
	// } else {
	// 	json.NewEncoder(w).Encode(check)
	// }
}
