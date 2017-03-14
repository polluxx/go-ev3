package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"cmd/serialapi"
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
	resultStr := serialapi.ColorStr(result)
	writeResponse(w, resultStr, err)
}

// Get luminosity value for connected sensor
func GetLuminosity(w http.ResponseWriter, r *http.Request) {
	port, _ := strconv.Atoi(mux.Vars(r)["port-id"])
	result, err := ev3.GetLuminosity(uint8(port))
	writeResponse(w, result, err)
}

// Get distance value for connected sensor
func GetDistance(w http.ResponseWriter, r *http.Request) {
	port, _ := strconv.Atoi(mux.Vars(r)["port-id"])
	result, err := ev3.GetDistance(uint8(port))
	writeResponse(w, result, err)
}

// Get is clickable sensor clicked
func GetIsClicked(w http.ResponseWriter, r *http.Request) {
	port, _ := strconv.Atoi(mux.Vars(r)["port-id"])
	result, err := ev3.GetIsClicked(uint8(port))
	writeResponse(w, result, err)
}

// Get click count for connected sensor
func GetClickCount(w http.ResponseWriter, r *http.Request) {
	port, _ := strconv.Atoi(mux.Vars(r)["port-id"])
	result, err := ev3.GetClickCount(uint8(port))
	writeResponse(w, result, err)
}

// Get gyro angle for connected sensor
func GetGyroAngle(w http.ResponseWriter, r *http.Request) {
	port, _ := strconv.Atoi(mux.Vars(r)["port-id"])
	result, err := ev3.GetGyroAngle(uint8(port))
	writeResponse(w, result, err)
}

// Get gyro angle for connected sensor
func GetGyroGravity(w http.ResponseWriter, r *http.Request) {
	port, _ := strconv.Atoi(mux.Vars(r)["port-id"])
	result, err := ev3.GetGyroGravity(uint8(port))
	writeResponse(w, result, err)
}

// Get value for generic sensor
func GetSensorValue(w http.ResponseWriter, r *http.Request) {
	port, _ := strconv.Atoi(mux.Vars(r)["port-id"])
	result, err := ev3.GetSensorValue(uint8(port), 0xFF)
	writeResponse(w, result, err)
}

// Start motor
func MotorStart(w http.ResponseWriter, r *http.Request) {
	port, _ := strconv.Atoi(mux.Vars(r)["port-id"])
	err := ev3.MoveMotorStart(uint8(port))
	writeResponse(w, nil, err)
}

// Stop motor and (not) apply brake
func MotorStop(w http.ResponseWriter, r *http.Request) {
	port, _ := strconv.Atoi(mux.Vars(r)["port-id"])
	brake, _ := strconv.Atoi(mux.Vars(r)["brake"])
	err := ev3.MoveMotorStop(uint8(port), uint8(brake))
	writeResponse(w, nil, err)
}

// Set motor speed
func SetMotorSpeed(w http.ResponseWriter, r *http.Request) {
	port, _ := strconv.Atoi(mux.Vars(r)["port-id"])
	speed, _ := strconv.Atoi(mux.Vars(r)["speed"])
	err := ev3.MoveMotorSpeed(uint8(port), int8(speed))
	writeResponse(w, nil, err)
}

// Get motor current angle
func GetMotorState(w http.ResponseWriter, r *http.Request) {
	port, _ := strconv.Atoi(mux.Vars(r)["port-id"])
	result, err := ev3.GetMotorAngle(uint8(port))
	writeResponse(w, result, err)
}
