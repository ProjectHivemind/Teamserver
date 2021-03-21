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
	router.Path("/api/implant").HandlerFunc(getImplants).Methods("GET")
	router.Path("/api/implant/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(getImplant).Methods("GET")
	router.Path("/api/implanttype").HandlerFunc(getImplantTypes).Methods("GET")
	router.Path("/api/implanttype/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(getImplantType).Methods("GET")
	router.Path("/api/callback/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(getCallBack).Methods("GET")
	router.Path("/api/implantswithcallbacks").HandlerFunc(getImplantsWithCallbacks).Methods("GET")

	// Group Funcs
	router.Path("/api/group").HandlerFunc(getGroups).Methods("GET")
	router.Path("/api/group/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(getGroup).Methods("GET")
	router.Path("/api/group").HandlerFunc(createGroup).Methods("POST")
	router.Path("/api/group/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(deleteGroup).Methods("DELETE")
	router.Path("/api/group/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(updateGroup).Methods("PUT")

	// Module Funcs
	router.Path("/api/module").HandlerFunc(getModules).Methods("GET")
	router.Path("/api/module/{id}").HandlerFunc(getModule).Methods("GET")
	router.Path("/api/modulefunc/{id}").HandlerFunc(getModuleFuncs).Methods("GET")

	// Stored Action Funcs
	router.Path("/api/storedaction").HandlerFunc(getStoredActions).Methods("GET")
	router.Path("/api/storedaction/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(getStoredAction).Methods("GET")
	router.Path("/api/storedaction").HandlerFunc(createStoredAction).Methods("POST")
	router.Path("/api/storedaction/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(deleteStoredAction).Methods("DELETE")

	// Staged Action Funcs
	router.Path("/api/stagedaction").HandlerFunc(getStagedActions).Methods("GET")
	router.Path("/api/stagedaction/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(getStagedAction).Methods("GET")
	router.Path("/api/stagedaction").HandlerFunc(createStagedAction).Methods("POST")
	router.Path("/api/stagedaction/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(deleteStagedAction).Methods("DELETE")

	// Executed Action Funcs
	router.Path("/api/stagedaction").HandlerFunc(getExecutedActions).Methods("GET")
	router.Path("/api/stagedaction/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(getExecutedAction).Methods("GET")

	fmt.Printf("RESTAPI on 0.0.0.0:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
