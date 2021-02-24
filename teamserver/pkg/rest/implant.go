package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
	"github.com/gorilla/mux"
)

func getImplantTypes(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	implantTypes, err := d.AllImplantTypes()
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		json.NewEncoder(w).Encode(implantTypes)
	}
}

func getImplantType(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	params := mux.Vars(r)
	implantType, err := d.GetImplantTypeById(params["id"])

	if err != nil {
		fmt.Fprint(w, err)
	} else {
		json.NewEncoder(w).Encode(implantType)
	}
}

func getImplants(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	implants, err := d.AllImplants()
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		json.NewEncoder(w).Encode(implants)
	}
}

func getImplant(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	params := mux.Vars(r)
	implant, err := d.GetImplantById(params["id"])

	if err != nil {
		fmt.Fprint(w, err)
	} else {
		json.NewEncoder(w).Encode(implant)
	}
}

func getCallBack(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	params := mux.Vars(r)
	callBack, err := d.GetCallBackById(params["id"])

	if err != nil {
		fmt.Fprint(w, err)
	} else {
		json.NewEncoder(w).Encode(callBack)
	}
}
