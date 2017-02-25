package serialapi

import (
	"errors"
	"log"
)

// Plays tone with specified params on the brick
func (self *EV3) PlaySound(volume uint8, frequency uint16, duration uint16) error {
	/*
		opSOUND Opcode sound related
		LC0(TONE) Command (TONE) encoded as single byte constant
		LC1(2) Sound-level 2 encoded as one constant byte to follow
		LC2(1000) Frequency 1000 Hz. encoded as two constant bytes to follow
		LC2(1000) Duration 1000 mS. encoded as two constant bytes to follow
		Example:
		0f0000008000009401810282e80382e803
	*/

	buf := make([]byte, 2)
	buf[0] = 0x94 // opSound
	buf[1] = 0x01 // type TONE, 3 params: 8, 16, 16
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

// Read devices
func (self *EV3) GetPortsStatus() (*EV3PortsStatus, error) {
	/*
		Opcode 0x99
		Arguments (Data8) CMD => Specific command parameter documented below
		Dispatch status Unchanged
		Description Read information about external sensor device
		CMD description is too huge to list it here, RTFM
		Example:
		45000100001000990500820000606199050082010062639905008202006465990500820300666799050082100068699905008211006a6b9905008212006c6d9905008213006e6f
		13000100027e007e007e007e007e007e007e007e00
	*/

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
	for i := 0; i < 8; i++ {
		portType = rep.byteCodes[i*2]
		portMode = rep.byteCodes[i*2+1]
		switch i {
		case 0:
			portsStatus.SensorPort1.Type = portType
			portsStatus.SensorPort1.Mode = portMode
		case 1:
			portsStatus.SensorPort2.Type = portType
			portsStatus.SensorPort2.Mode = portMode
		case 2:
			portsStatus.SensorPort3.Type = portType
			portsStatus.SensorPort3.Mode = portMode
		case 3:
			portsStatus.SensorPort4.Type = portType
			portsStatus.SensorPort4.Mode = portMode
		case 4:
			portsStatus.MotorPortA.Type = portType
			portsStatus.MotorPortA.Mode = portMode
		case 5:
			portsStatus.MotorPortB.Type = portType
			portsStatus.MotorPortB.Mode = portMode
		case 6:
			portsStatus.MotorPortC.Type = portType
			portsStatus.MotorPortC.Mode = portMode
		case 7:
			portsStatus.MotorPortD.Type = portType
			portsStatus.MotorPortD.Mode = portMode
		}
	}
	return &portsStatus, nil
}

// Start motor
// Read motor angle
// Read color
// Read luminosity
// Read distance
