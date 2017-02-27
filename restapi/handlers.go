package restapi

import (
	"net/http"
	"encoding/json"
)

func GetPortsStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	result := struct {}{}
	json.NewEncoder(w).Encode(result)
}

func GetAngle(w http.ResponseWriter, r *http.Request) {}
func SetMotorState(w http.ResponseWriter, r *http.Request) {}
func GetColor(w http.ResponseWriter, r *http.Request) {}
func GetLuminosity(w http.ResponseWriter, r *http.Request) {}
func GetDistance(w http.ResponseWriter, r *http.Request) {}
func GetSwitch(w http.ResponseWriter, r *http.Request) {}
