package serialapi

// Brick port entity
type EV3Port struct {
	Type uint8
	Mode uint8
}

// Describes current brick ports type and mode
type EV3PortsStatus struct {
	SensorPort1 EV3Port
	SensorPort2 EV3Port
	SensorPort3 EV3Port
	SensorPort4 EV3Port

	MotorPortA EV3Port
	MotorPortB EV3Port
	MotorPortC EV3Port
	MotorPortD EV3Port
}
