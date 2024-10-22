package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getImplantTypes(w http.ResponseWriter, r *http.Request) {
	implantTypes, err := d.AllImplantTypes()
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(implantTypes)
	}
}

func getImplantType(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	implantType, err := d.GetImplantTypeById(id)

	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(implantType)
	}
}

func getImplants(w http.ResponseWriter, r *http.Request) {
	implants, err := d.AllImplants()
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(implants)
	}
}

func getImplant(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	implant, err := d.GetImplantById(id)

	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(implant)
	}
}

func getCallBack(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	callBack, err := d.GetCallBackById(id)

	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(callBack)
	}
}

func getImplantsWithCallbacks(w http.ResponseWriter, r *http.Request) {
	callBack, err := d.GetImplantsWithCallbacks()

	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(callBack)
	}
}
