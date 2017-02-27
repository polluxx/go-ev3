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

type Color struct {
	Value uint8
}

func (self *Color) String() string {
	switch self.Value {
	case 1:
		return "Black"
	case 2:
		return "Blue"
	case 3:
		return "Green"
	case 4:
		return "Yellow"
	case 5:
		return "Red"
	case 6:
		return "White"
	case 7:
		return "Brown"
	}
	return "None"	// equal to 0
}