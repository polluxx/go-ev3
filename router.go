package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/ports", GetPortsStatus).Methods("GET")

	router.HandleFunc("/sound/{volume:[0-9]{1,3}}/frequency/{frequency:[0-9]{1,5}}/duration/{duration:[0-9]{1,5}}",
		PlaySound).Methods("POST")

	router.HandleFunc("/motors/{port-id:[0-9]{1,3}}/start", MotorStart).Methods("POST")
	router.HandleFunc("/motors/{port-id:[0-9]{1,3}}/stop/{brake:[0-1]}", MotorStop).Methods("POST")
	router.HandleFunc("/motors/{port-id:[0-9]{1,3}}/speed/{speed}", SetMotorSpeed).Methods("POST")
	router.HandleFunc("/motors/{port-id:[0-9]{1,3}}", GetMotorState).Methods("GET")

	router.HandleFunc("/sensors/{port-id:[0-3]}/color", GetColor).Methods("GET")
	router.HandleFunc("/sensors/{port-id:[0-3]}/luminosity", GetLuminosity).Methods("GET")
	router.HandleFunc("/sensors/{port-id:[0-3]}/distance", GetDistance).Methods("GET")
	router.HandleFunc("/sensors/{port-id:[0-3]}/click", GetIsClicked).Methods("GET")
	router.HandleFunc("/sensors/{port-id:[0-3]}/clickcount", GetClickCount).Methods("GET")
	router.HandleFunc("/sensors/{port-id:[0-3]}/gyro/angle", GetGyroAngle).Methods("GET")
	router.HandleFunc("/sensors/{port-id:[0-3]}/gyro/gravity", GetGyroGravity).Methods("GET")
	router.HandleFunc("/sensors/{port-id:[0-3]}", GetSensorValue).Methods("GET")

	return router
}
