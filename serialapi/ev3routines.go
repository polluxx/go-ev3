package serialapi

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"github.com/tarm/goserial"
	"io"
	"log"
)

// Lego brick constants
const (
	MotorPortA uint8 = 0x01
	MotorPortB = 0x02
	MotorPortC = 0x04
	MotorPortD = 0x08
)

const (
	SensorPort1 uint8 = iota
	SensorPort2
	SensorPort3
	SensorPort4
)

const (
	MaxMessageCount    = 0xFFFF
	CommandWithReply   = 0x00
	CommandWithNOReply = 0x80
	ReplayOk           = 0x02

	Data8   = 0x81
	Data16  = 0x82
	Data32  = 0x83
	DataStr = 0x84

	EV3MessageHeaderSize = 7
	EV3ReplayHeaderSize  = 5
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

// Functionality to wrap message entity to byte array conversion according to specification
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

// Functionality to wrap replay entity parsing
func getReplay(buf []byte) (*EV3Reply, error) {
	if len(buf) < EV3ReplayHeaderSize {
		err := errors.New("Replay buffer is too small")
		log.Fatal(err)
		return nil, err
	}

	reply := EV3Reply{
		replySize:    binary.LittleEndian.Uint16(buf[0:]),
		messageCount: binary.LittleEndian.Uint16(buf[2:]),
		replyType:    buf[4],
		byteCodes:    buf[5:], // TODO: Check this index is Ok in case of small reply
	}

	if reply.replyType != ReplayOk {
		err := errors.New("Replay reported as failed")
		log.Fatal(err)
		return nil, err
	}
	return &reply, nil
}

// Serial over BT wrapper: send
func (self *EV3) sendBytes(buf []byte) error {
	// Open port if not open yet
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

	log.Println(hex.EncodeToString(buf)) // DEBUG INFO
	if self.DebugOn == false {
		_, err := self.port.Write(buf)
		return err
	}
	return nil
}

// Serial over BT wrapper: receive
func (self *EV3) receiveBytes() ([]byte, error) {
	if self.DebugOn == false {
		// Open port if not open yet
		if self.port == nil {
			config := &serial.Config{Name: self.PortName, Baud: self.PortBaud}
			var err error
			self.port, err = serial.OpenPort(config)
			if err != nil {
				log.Fatal(err)
				return nil, err
			}
		}

		// Read size of following message
		buf := make([]byte, 2)
		n, err := self.port.Read(buf)
		if err != nil {
			log.Fatal(err)
			return nil, err
		} else if n != len(buf) {
			err = errors.New("Too few bytes read: expected replay header")
			log.Fatal(err)
			return nil, err
		}
		replySize := binary.LittleEndian.Uint16(buf)

		// Read message tail
		bufTail := make([]byte, replySize)
		n, err = self.port.Read(bufTail)
		if err != nil {
			log.Fatal(err)
			return nil, err
		} else if n != len(bufTail) {
			err = errors.New("Too few bytes read: expected replay tail")
			log.Fatal(err)
			return nil, err
		}

		// Stitch and return
		buf = append(buf, bufTail...)
		log.Println(hex.EncodeToString(buf)) // DEBUG INFO
		return buf, nil
	}
	return make([]byte, 0), nil
}

// Close connection to brick
func (self *EV3) Close() {
	if self.port != nil {
		self.port.Close()
	}
}

// Helper routine for variable and constants size allocation
func variablesReservation(globalSize uint8, localSize uint8) uint16 {
	return uint16(globalSize&0xFF) | uint16(((globalSize>>8)&0x3)|((localSize<<2)&0xFC))<<8
}

// Helper routine to generate variable global index
func getVarGlobalIndex(index int) []byte {
	var buf []byte
	if index <= 31 {
		buf = make([]byte, 1)
		buf[0] = uint8(index) | 0x60
	} else if index <= 255 {
		buf = make([]byte, 1)
		buf[0] = 0xE1
		buf[1] = uint8(index)
	} else if index <= 65535 {
		buf = make([]byte, 1)
		buf[0] = 0xE2
		buf[1] = uint8(index) & 0xFF
		buf[2] = uint8(index>>8) & 0xFF
	} else {
		buf = make([]byte, 1)
		buf[0] = 0xE3
		buf[1] = uint8(index) & 0xFF
		buf[2] = uint8(index>>8) & 0xFF
		buf[3] = uint8(index>>16) & 0xFF
		buf[4] = uint8(index>>24) & 0xFF
	}
	return buf
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
