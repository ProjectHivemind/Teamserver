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
	router.Path("/module").HandlerFunc(getModules).Methods("GET")
	router.Path("/module/{id}").HandlerFunc(getModule).Methods("GET")
	router.Path("/modulefunc/{id}").HandlerFunc(getModuleFuncs).Methods("GET")

	// Stored Action Funcs
	router.Path("/storedaction").HandlerFunc(getStoredActions).Methods("GET")
	router.Path("/storedaction/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(getStoredAction).Methods("GET")
	router.Path("/storedaction").HandlerFunc(createStoredAction).Methods("POST")
	router.Path("/storedaction/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(deleteStoredAction).Methods("DELETE")

	// Staged Action Funcs
	router.Path("/stagedaction").HandlerFunc(getStagedActions).Methods("GET")
	router.Path("/stagedaction/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(getStagedAction).Methods("GET")
	router.Path("/stagedaction").HandlerFunc(createStagedAction).Methods("POST")
	router.Path("/stagedaction/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(deleteStagedAction).Methods("DELETE")

	// Executed Action Funcs
	router.Path("/stagedaction").HandlerFunc(getExecutedActions).Methods("GET")
	router.Path("/stagedaction/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(getExecutedAction).Methods("GET")

	fmt.Printf("RESTAPI on 0.0.0.0:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
