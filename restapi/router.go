package restapi

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/ports", GetPortsStatus).Methods("GET")

	router.HandleFunc("/motors/{port-id}/angle", GetAngle).Methods("GET")
	router.HandleFunc("/motors/{port-id}", SetMotorState).Methods("POST")

	router.HandleFunc("/sensors/{port-id}/color", GetColor).Methods("GET")
	router.HandleFunc("/sensors/{port-id}/luminosity", GetLuminosity).Methods("GET")
	router.HandleFunc("/sensors/{port-id}/distance", GetDistance).Methods("GET")
	router.HandleFunc("/sensors/{port-id}/switch", GetSwitch).Methods("GET")

	return router
}