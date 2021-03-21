package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/gorilla/mux"
)

func getSession(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	token := mux.Vars(r)["token"]
	session, err := d.GetSessionById(token)

	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(session)
	}
}

func insertSession(w http.ResponseWriter, r *http.Request) {
	var d crud.DatabaseModel
	d.Open()
	defer d.Close()

	exptime := r.FormValue("ExpTime")
	t, err := time.Parse(crud.TimeStamp, exptime)
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
		return
	}
	if t.After(time.Now()) {
		fmt.Fprint(w, "time is invalid")
		return
	}

	token := mux.Vars(r)["token"]
	session := model.Sessions{
		SessionToken: token,
		Username:     r.FormValue("username"),
		ExpTime:      exptime,
	}

	_, err = d.InsertSession(session)
	if err != nil {
		fmt.Fprint(w, GENERAL_ERROR)
	} else {
		json.NewEncoder(w).Encode(session)
	}
}
