package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Response helper functionality
func writeResponse(w http.ResponseWriter, v interface{}, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		if v != nil {
			json.NewEncoder(w).Encode(v)
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}
}

// Get port status
func GetPortsStatus(w http.ResponseWriter, r *http.Request) {
	result, err := ev3.GetPortsStatus()
	writeResponse(w, result, err)
}

// Plays sound on brick side
func PlaySound(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volume, _ := strconv.Atoi(vars["volume"])
	frequency, _ := strconv.Atoi(vars["frequency"])
	duration, _ := strconv.Atoi(vars["duration"])

	err := ev3.PlaySound(uint8(volume), uint16(frequency), uint16(duration))
	writeResponse(w, nil, err)
}

// Get color value for connected sensor
func GetColor(w http.ResponseWriter, r *http.Request) {
	port, _ := strconv.Atoi(mux.Vars(r)["port-id"])
	result, err := ev3.GetColor(uint8(port))
	writeResponse(w, result, err)
}

func MotorStart(w http.ResponseWriter, r *http.Request)    {}
func MotorStop(w http.ResponseWriter, r *http.Request)     {}
func SetMotorSpeed(w http.ResponseWriter, r *http.Request) {}
func GetMotorState(w http.ResponseWriter, r *http.Request) {}

func GetLuminosity(w http.ResponseWriter, r *http.Request)  {}
func GetDistance(w http.ResponseWriter, r *http.Request)    {}
func GetSwitch(w http.ResponseWriter, r *http.Request)      {}
func GetSwitchCount(w http.ResponseWriter, r *http.Request) {}
