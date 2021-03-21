package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
	"github.com/gorilla/mux"
)

func getOperators(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	operators, err := d.AllOperators()

	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		// Blank out the passwords so it doesn't get sent
		for i := 0; i < len(operators); i++ {
			operators[i].Password = ""
		}

		json.NewEncoder(w).Encode(operators)
	}
}

func getOperator(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	username := mux.Vars(r)["username"]
	operator, err := d.GetOperatorByUsername(username)
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		// Blank out the password so it doesn't get sent
		operator.Password = ""

		json.NewEncoder(w).Encode(operator)
	}
}

func authOperator(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	username := r.FormValue("username")
	password := r.FormValue("password")
	check, err := d.CheckOperator(username, password)
	if err != nil {
		fmt.Fprint(w, "auth failed")
	} else {
		json.NewEncoder(w).Encode(check)
	}
}
