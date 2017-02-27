package serialapi

import (
	"encoding/binary"
	"errors"
	"log"
	"math"
)

// Plays tone with specified params on the brick
func (self *EV3) PlaySound(volume uint8, frequency uint16, duration uint16) error {
	buf := make([]byte, 0)
	buf = append(buf, 0x94 /*opSOUND*/, 0x01 /*type TONE*/)
	buf = append(buf, LC8(volume)...)
	buf = append(buf, LC16(frequency)...)
	buf = append(buf, LC16(duration)...)

	msg := EV3Message{
		messageCount:         self.messageCount,
		commandType:          CommandWithNOReply,
		variablesReservation: 0x00,
		byteCodes:            buf,
	}
	return self.sendBytes(msg.getBytes())
}

// Stop motor on specific port(s)
func (self *EV3) MoveMotorStop(ports uint8, brake uint8) error {
	buf := make([]byte, 0)
	buf = append(buf, 0xA3 /*opOutput_Step_Speed*/, 0x00 /*module id*/, ports, brake)

	msg := EV3Message{
		messageCount:         self.messageCount,
		commandType:          CommandWithNOReply,
		variablesReservation: 0x00,
		byteCodes:            buf,
	}
	return self.sendBytes(msg.getBytes())
}

// Start motor on specific port(s)
func (self *EV3) MoveMotorStart(ports uint8) error {
	buf := make([]byte, 0)
	buf = append(buf, 0xA6 /*opOutput_Step_Speed*/, 0x00 /*module id*/, ports)

	msg := EV3Message{
		messageCount:         self.messageCount,
		commandType:          CommandWithNOReply,
		variablesReservation: 0x00,
		byteCodes:            buf,
	}
	return self.sendBytes(msg.getBytes())
}

// Set and maintain motor speed on specific port(s)
func (self *EV3) MoveMotorSpeed(ports uint8, speed int8) error {
	buf := make([]byte, 0)
	buf = append(buf, 0xA5 /*opOutput_Step_Speed*/, 0x00 /*module id*/, ports)
	buf = append(buf, LC8(uint8(speed))...)

	msg := EV3Message{
		messageCount:         self.messageCount,
		commandType:          CommandWithNOReply,
		variablesReservation: 0x00,
		byteCodes:            buf,
	}
	return self.sendBytes(msg.getBytes())
}

// Move motor for set angle on specific port(s)
func (self *EV3) MoveMotorAngle(ports uint8, speed int8, angle int32, brake uint8) error {
	buf := make([]byte, 0)
	buf = append(buf, 0xAE /*opOutput_Step_Speed*/, 0x00 /*module id*/, ports)
	buf = append(buf, LC8(uint8(speed))...)
	buf = append(buf, LC32(0 /*immediate start*/)...)
	buf = append(buf, LC32(uint32(angle))...)
	buf = append(buf, LC32(0 /*immediate stop*/)...)
	buf = append(buf, brake)

	msg := EV3Message{
		messageCount:         self.messageCount,
		commandType:          CommandWithNOReply,
		variablesReservation: 0x00,
		byteCodes:            buf,
	}
	return self.sendBytes(msg.getBytes())
}

// Move motor for set time on specific port(s)
func (self *EV3) MoveMotorTime(ports uint8, speed int8, timeMs int32, brake uint8) error {
	buf := make([]byte, 0)
	buf = append(buf, 0xAF /*opOutput_Time_Speed*/, 0x00 /*module id*/, ports)
	buf = append(buf, LC8(uint8(speed))...)
	buf = append(buf, LC32(0 /*immediate start*/)...)
	buf = append(buf, LC32(uint32(timeMs))...)
	buf = append(buf, LC32(0 /*immediate stop*/)...)
	buf = append(buf, brake)

	msg := EV3Message{
		messageCount:         self.messageCount,
		commandType:          CommandWithNOReply,
		variablesReservation: 0x00,
		byteCodes:            buf,
	}
	return self.sendBytes(msg.getBytes())
}

// Read devices
func (self *EV3) GetPortsStatus() (*EV3PortsStatus, error) {
	// Prepare message to check all 8 ports at once
	buf := make([]byte, 0)
	for i := 0; i < 8; i++ {
		buf = append(buf, 0x99, 0x05, 0x00)
		//buf = append(buf, LC8(0)...)
		if i < 4 {
			buf = append(buf, LC16(uint16(i))...)
		} else {
			buf = append(buf, LC16(uint16(i+12))...)
		}
		buf = append(buf, getVarGlobalIndex(i*2)...)   // write global index: type
		buf = append(buf, getVarGlobalIndex(i*2+1)...) // write global index: units
	}

	msg := EV3Message{
		messageCount:         self.messageCount,
		commandType:          CommandWithReply,
		variablesReservation: variablesReservation(16, 0),
		byteCodes:            buf,
	}
	err := self.sendBytes(msg.getBytes())

	// Receive response, check msg count & parse result
	if err != nil {
		return nil, err
	}
	buf, err = self.receiveBytes()
	if err != nil {
		return nil, err
	}
	rep, err := getReplay(buf)
	if err != nil {
		return nil, err
	}
	if rep.messageCount != msg.messageCount {
		err = errors.New("Received replay to another message")
		log.Fatal(err)
		return nil, err
	}
	if len(rep.byteCodes) != 16 {
		err = errors.New("Received replay contains not enough data")
		log.Fatal(err)
		return nil, err
	}

	// Parse response for all 8 ports
	portsStatus := EV3PortsStatus{}
	var portType, portMode uint8
	var portTypeStr string

	for i := 0; i < 8; i++ {
		portType = rep.byteCodes[i*2]
		portMode = rep.byteCodes[i*2+1]
		// Convert to awesome string format
		portTypeStr = DeviceTypeStr(portType)
		switch i {
		case 0:
			portsStatus.SensorPort1.Type = portTypeStr
			portsStatus.SensorPort1.Mode = portMode
		case 1:
			portsStatus.SensorPort2.Type = portTypeStr
			portsStatus.SensorPort2.Mode = portMode
		case 2:
			portsStatus.SensorPort3.Type = portTypeStr
			portsStatus.SensorPort3.Mode = portMode
		case 3:
			portsStatus.SensorPort4.Type = portTypeStr
			portsStatus.SensorPort4.Mode = portMode
		case 4:
			portsStatus.MotorPortA.Type = portTypeStr
			portsStatus.MotorPortA.Mode = portMode
		case 5:
			portsStatus.MotorPortB.Type = portTypeStr
			portsStatus.MotorPortB.Mode = portMode
		case 6:
			portsStatus.MotorPortC.Type = portTypeStr
			portsStatus.MotorPortC.Mode = portMode
		case 7:
			portsStatus.MotorPortD.Type = portTypeStr
			portsStatus.MotorPortD.Mode = portMode
		}
	}
	return &portsStatus, nil
}

// Read sensor value
func (self *EV3) GetSensorValue(port uint8, sensorMode uint8) (uint8, error) {
	buf := make([]byte, 0)
	buf = append(buf, 0x99 /*opInput_Device*/, 0x1D /*READY_SI*/, 0x00 /*LAYER*/, port, 0x00 /*TYPE*/, sensorMode, 0x01 /*Number of return values */)
	buf = append(buf, getVarGlobalIndex(0)...)

	msg := EV3Message{
		messageCount:         self.messageCount,
		commandType:          CommandWithReply,
		variablesReservation: variablesReservation(4, 0),
		byteCodes:            buf,
	}
	err := self.sendBytes(msg.getBytes())

	// Receive response, check msg count & parse result
	if err != nil {
		return 0, err
	}
	buf, err = self.receiveBytes()
	if err != nil {
		return 0, err
	}
	rep, err := getReplay(buf)
	if err != nil {
		return 0, err
	}
	if rep.messageCount != msg.messageCount {
		err = errors.New("Received replay to another message")
		log.Fatal(err)
		return 0, err
	}
	if len(rep.byteCodes) != 4 {
		err = errors.New("Received replay contains not enough data")
		log.Fatal(err)
		return 0, err
	}

	// Parse response
	intVal := binary.LittleEndian.Uint32(rep.byteCodes)
	floatVal := math.Float32frombits(intVal)
	return uint8(floatVal), nil
}

// Read light reflection
func (self *EV3) GetLightReflection(port uint8) (uint8, error) {
	return self.GetSensorValue(port, 0x00)
}

// Read luminosity
func (self *EV3) GetLuminosity(port uint8) (uint8, error) {
	return self.GetSensorValue(port, 0x01)
}

// Read color
func (self *EV3) GetColor(port uint8) (uint8, error) {
	return self.GetSensorValue(port, 0x02)
}

// Read is clicked
func (self *EV3) GetIsClicked(port uint8) (uint8, error) {
	return self.GetSensorValue(port, 0x00)
}

// Read click cound
func (self *EV3) GetClickCount(port uint8) (uint8, error) {
	return self.GetSensorValue(port, 0x01)
}

// Read distance
func (self *EV3) GetDistance(port uint8) (uint8, error) {
	return self.GetSensorValue(port, 0x00)
}

// Read gyro angle
func (self *EV3) GetGyroAngle(port uint8) (uint8, error) {
	return self.GetSensorValue(port, 0x00)
}

// Read gyro gravity
func (self *EV3) GetGyroGravity(port uint8) (uint8, error) {
	return self.GetSensorValue(port, 0x01)
}

// Get current motor rotot angle
func (self *EV3) GetMotorAngle(port uint8) (uint8, error) {
	// 16=PortA, 17=PortB, 18=PortC, 19=PortD
	// 0=angle, 1=rotation count, 2=power
	return self.GetSensorValue(uint8(0x0F+port), 0x00)
}
