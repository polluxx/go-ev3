package main

import (
	"go-ev3/serialapi"
	"log"
	"net/http"
	"fmt"
)

var ev3 serialapi.EV3

// Demo app main entry point
func main() {
	// Run REST API
	// TODO: Read values from config/input params
	initRestApi(8081, "/dev/rfcomm7", 9600)

	// Execute DEMO code
	//runDemo()
}

// Initializes EV3 interface REST API
func initRestApi(httpPort int, serialPort string, serialBaud int) {
	ev3 = serialapi.EV3{
		PortName: serialPort,
		PortBaud: serialBaud,
		DebugOn:  false,
	}

	router := NewRouter()
	log.Println("Go EV3 REST API server initialized")
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", httpPort), router))
}

// Execute demo code
func runDemo()  {
	ev3 := serialapi.EV3{
		PortName: "/dev/rfcomm6",
		PortBaud: 9600,
		//DebugOn:  true,
	}
	ev3.PlaySound(2, 1000, 200)

	/*
	portsStatus, _ := ev3.GetPortsStatus()
	fmt.Printf("%v", col)
	*/

	/*
	col, _ := ev3.GetColor(serialapi.SensorPort2)
	fmt.Printf("%v\n", col)
	*/

	//ev3.MoveMotorAngle(serialapi.MotorPortD, 50, 90)
	/*
	ev3.MoveMotorSpeed(serialapi.MotorPortD, 0)
	ev3.MoveMotorStart(serialapi.MotorPortD)
	ev3.MoveMotorSpeed(serialapi.MotorPortD, 100)
	time.Sleep(1 * time.Second)
	ev3.MoveMotorSpeed(serialapi.MotorPortD, 30)
	time.Sleep(1 * time.Second)
	ev3.MoveMotorSpeed(serialapi.MotorPortD, 0)
	ev3.MoveMotorStop(serialapi.MotorPortD, 0)
	*/
}