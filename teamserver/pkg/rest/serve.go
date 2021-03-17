package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const GENERAL_ERROR = "there is an error"

func Start(port string) {
	router := mux.NewRouter()

	// Implant Functions
	router.Path("/implant").HandlerFunc(getImplants).Methods("GET")
	router.Path("/implant/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(getImplant).Methods("GET")
	router.Path("/implanttype").HandlerFunc(getImplantTypes).Methods("GET")
	router.Path("/implanttype/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(getImplantType).Methods("GET")
	router.Path("/callback/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(getCallBack).Methods("GET")

	// Group Funcs
	router.Path("/group").HandlerFunc(getGroups).Methods("GET")
	router.Path("/group/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(getGroup).Methods("GET")
	router.Path("/group").HandlerFunc(createGroup).Methods("POST")
	router.Path("/group/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(deleteGroup).Methods("DELETE")
	router.Path("/group/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(updateGroup).Methods("PUT")

	// Module Funcs

	// Stored Action Funcs

	// Staged Action Funcs

	// Executed Action Funcs

	fmt.Printf("RESTAPI on 0.0.0.0:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
