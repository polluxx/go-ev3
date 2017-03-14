package serialapi

// Brick port entity
type EV3Port struct {
	// Id
	Type string
	Mode uint8 `json:"-"`
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

// Returns human readable DeviceType
func DeviceTypeStr(value uint8) string {
	switch value {
	case 7:
		return "Big motor"
	case 8:
		return "Medium motor"
	case 16:
		return "Touch sensor"
	case 29:
		return "Color sensor"
	case 30:
		return "Ultrasonic distance sensor"
	case 32:
		return "Gyro sensor"
	case 33:
		return "Optical distance sensor"
	case 0x7e:
		return "Empty"
	}
	return "Unknown"
}

// Returns human readable Color
func ColorStr(value uint8) string {
	switch value {
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
	return "None" // equal to 0
}
