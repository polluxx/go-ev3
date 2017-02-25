package serialapi

// Plays tone with specified params on the brick
func (self *EV3) PlaySound(volume uint8, frequency uint16, duration uint16) error {
	/*
		opSOUND Opcode sound related
		LC0(TONE) Command (TONE) encoded as single byte constant
		LC1(2) Sound-level 2 encoded as one constant byte to follow
		LC2(1000) Frequency 1000 Hz. encoded as two constant bytes to follow
		LC2(1000) Duration 1000 mS. encoded as two constant bytes to follow
	*/

	buf := make([]byte, 2)
	buf[0] = 0x94 // opSound
	buf[1] = 0x01 // type TONE, 3 params: 8, 16, 16
	buf = append(buf, LC8(volume)...)
	buf = append(buf, LC16(frequency)...)
	buf = append(buf, LC16(duration)...)

	msg := EV3Message{
		messageCount: self.messageCount,
		commandType:  CommandWithNOReply,
		byteCodes:    buf,
	}
	return self.sendBytes(msg.getBytes())
}
