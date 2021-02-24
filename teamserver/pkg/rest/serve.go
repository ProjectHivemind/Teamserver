package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start(port string) {
	router := mux.NewRouter()

	// Implant Functions
	router.HandleFunc("/api/implanttypes", getImplantTypes).Methods("GET")
	router.HandleFunc("/api/implanttype", getImplantType).Methods("GET")
	router.HandleFunc("/api/implants", getImplants).Methods("GET")
	router.HandleFunc("/api/implant", getImplant).Methods("GET")
	router.HandleFunc("/api/callback", getCallBack).Methods("GET")

	// Group Funcs

	// Module Funcs

	// Stored Action Funcs

	// Staged Action Funcs

	// Executed Action Funcs

	fmt.Printf("RESTAPI on 0.0.0.0:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
