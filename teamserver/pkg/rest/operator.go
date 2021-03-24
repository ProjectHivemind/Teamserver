package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getOperators(w http.ResponseWriter, r *http.Request) {
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
	username := r.FormValue("username")
	password := r.FormValue("password")
	check, err := d.CheckOperator(username, password)
	if err != nil {
		fmt.Fprint(w, "auth failed")
	} else {
		fmt.Fprint(w, check)
	}
}
