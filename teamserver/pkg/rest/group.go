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

func getGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := d.AllGroups()
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(groups)
	}
}

func getGroup(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	group, err := d.GetGroupById(id)

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

	var group model.Group
	err = json.Unmarshal(reqBody, &group)
	if err != nil {
		fmt.Fprint(w, err)
	}
	group.UUID = uuid.New().String()

	_, err = d.InsertGroup(group)
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(group)
	}
}

func deleteGroup(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := d.RemoveGroupById(id)

	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		fmt.Fprint(w, "deleted")
	}
}

func updateGroup(w http.ResponseWriter, r *http.Request) {
	// Check if the group exists
	id := mux.Vars(r)["id"]
	group, err := d.GetGroupById(id)
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	}

	// Grab form data
	r.ParseMultipartForm(0)
	action := r.FormValue("action")
	implant := r.FormValue("implant")

	// Check if the implant is already present
	for _, val := range group.Implants {
		if val == implant {
			if action == "add" {
				fmt.Fprint(w, GENERAL_ERROR)
				return
			} else if action == "remove" {
				check, _ := d.RemoveUUIDFromGroup(group.UUID, implant)
				if !check {
					fmt.Fprint(w, GENERAL_ERROR)
				} else {
					fmt.Fprint(w, "removed")
				}
				return
			}
		}
	}

	// Fails because the implant to be removed is not present
	if action == "remove" {
		fmt.Fprint(w, GENERAL_ERROR)
		return
	}

	_, err = d.AddUUIDToGroup(group.UUID, implant)
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		fmt.Fprint(w, "added")
	}
}
