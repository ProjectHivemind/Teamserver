package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
	"github.com/gorilla/mux"
)

const GENERAL_ERROR = "there is an error"

var d crud.DatabaseModel

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
	router.Path("/api/stagedactionfrontend").HandlerFunc(getStagedActionsFrontend).Methods("GET")
	router.Path("/api/stagedaction/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(getStagedAction).Methods("GET")
	router.Path("/api/stagedaction").HandlerFunc(createStagedAction).Methods("POST")
	router.Path("/api/stagedaction/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(deleteStagedAction).Methods("DELETE")

	// Executed Action Funcs
	router.Path("/api/executedaction").HandlerFunc(getExecutedActions).Methods("GET")
	router.Path("/api/executedaction/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").HandlerFunc(getExecutedAction).Methods("GET")

	// Operator Funcs
	router.Path("/api/operator").HandlerFunc(getOperators).Methods("GET")
	router.Path("/api/operator/{username}").HandlerFunc(getOperator).Methods("GET")
	router.Path("/api/operator/auth").HandlerFunc(authOperator).Methods("POST")

	// Session Funcs
	router.Path("/api/session/{token}").HandlerFunc(getSession).Methods("GET")
	router.Path("/api/session/{token}").HandlerFunc(insertSession).Methods("POST")
	router.Path("/api/session/validate/{token}").HandlerFunc(validateSession).Methods("GET")
	router.Path("/api/session/{token}").HandlerFunc(deleteSession).Methods("DELETE")

	// Start Database Connection
	d.Open()
	defer d.Close()

	fmt.Printf("RESTAPI on 0.0.0.0:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
