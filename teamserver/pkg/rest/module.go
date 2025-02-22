package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/gorilla/mux"
)

func getModules(w http.ResponseWriter, r *http.Request) {
	modules, err := d.AllModules()
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(modules)
	}
}

func getModule(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	module, err := d.GetModuleByName(id)

	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(module)
	}
}

func getModuleFuncs(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	module, err := d.GetModuleByName(id)

	var moduleFuncs []model.ModulesFunc
	for _, val := range module.ModuleFuncIds {
		moduleFunc, err := d.GetModuleFuncById(val)
		if err == nil {
			moduleFuncs = append(moduleFuncs, moduleFunc)
		}
	}

	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(moduleFuncs)
	}
}
