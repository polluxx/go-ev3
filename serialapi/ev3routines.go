package serialapi

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/tarm/goserial"
	"io"
	"log"
)

// Lego brick constants
const (
	MotorPortA int = iota
	MotorPortB
	MotorPortC
	MotorPortD

	SensorPort1
	SensorPort2
	SensorPort3
	SensorPort4

	MaxMessageCount    = 0xFFFF
	CommandWithReply   = 0x00
	CommandWithNOReply = 0x80

	Data8   = 0x81
	Data16  = 0x82
	Data32  = 0x83
	DataStr = 0x84

	EV3MessageHeaderSize  = 7
	EV3ResponseHeaderSize = 5
)

// Lego brick API interface
type EV3 struct {
	messageCount uint16
	port         io.ReadWriteCloser
	PortName     string
	PortBaud     int
	DebugOn      bool
}

// Brick message entity
type EV3Message struct {
	commandSize          uint16
	messageCount         uint16
	commandType          uint8
	variablesReservation uint16
	byteCodes            []byte
}

// Brick response entity
type EV3Reply struct {
	replySize    uint16
	messageCount uint16
	replyType    uint8
	byteCodes    []byte
}

// Functionality to wrap message entity to byte array according to specification
func (self *EV3Message) getBytes() []byte {
	self.commandSize = uint16(len(self.byteCodes)) + EV3MessageHeaderSize - 2 // 2 means commandSize = uint16 that should not be counted
	buf := make([]byte, EV3MessageHeaderSize)
	binary.LittleEndian.PutUint16(buf[0:], self.commandSize)
	binary.LittleEndian.PutUint16(buf[2:], self.messageCount)
	buf[4] = self.commandType
	binary.LittleEndian.PutUint16(buf[5:], self.variablesReservation)
	buf = append(buf, self.byteCodes...)
	return buf
}

// Serial over BT wrapper
func (self *EV3) sendBytes(buf []byte) error {
	if (self.port == nil) && (self.DebugOn == false) {
		config := &serial.Config{Name: self.PortName, Baud: self.PortBaud}
		var err error
		self.port, err = serial.OpenPort(config)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	self.messageCount = self.messageCount + 1
	if self.messageCount > MaxMessageCount {
		self.messageCount = 0
	}

	/*
		buf := make([]byte, 40)
		n, err := s.Read(buf)
		fmt.Println("Data read")

		if err != nil {
			fmt.Println(err)
		}
		s.Close()
	*/

	if self.DebugOn {
		log.Println(hex.EncodeToString(buf)) // DEBUG INFO
		return nil
	} else {
		_, err := self.port.Write(buf)
		return err
	}
}

// Helper routines for variable and constants wrapping according to protocol specification
func LC8(val uint8) []byte {
	buf := make([]byte, 2)
	buf[0] = Data8
	buf[1] = val
	return buf
}

func LC16(val uint16) []byte {
	buf := make([]byte, 3)
	buf[0] = Data16
	binary.LittleEndian.PutUint16(buf[1:], val)
	return buf
}

func LC32(val uint32) []byte {
	buf := make([]byte, 5)
	buf[0] = Data32
	binary.LittleEndian.PutUint32(buf[1:], val)
	return buf
}
