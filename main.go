package main

import (
	"fmt"
	"go-ev3/serialapi"
	"time"
)

// Demo app main entry point
func main() {
	fmt.Println("EV3 API demo")
	time.Sleep(1 * time.Millisecond)

	ev3 := serialapi.EV3{
		PortName: "/dev/rfcomm5",
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
	fmt.Printf("%v", col)
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
