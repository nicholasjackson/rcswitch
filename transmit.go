package rcswitch

import (
	"log"
	"time"

	"periph.io/x/periph/conn/gpio"
)

// Send the given data using the protocol defined by the protocol id
func (s *RCSwitch) Send(data string, protocolID int) {
	p, err := findProtocol(protocolID)
	if err != nil {
		log.Fatal("Unable to find protocol")
	}

	log.Printf("Transmitting: %s", data)
	for i := 0; i < 10; i++ {
		s.transmit(data, p)
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *RCSwitch) transmit(data string, p Protocol) {
	for _, d := range data {
		if d == '1' {
			s.transmitPulse(p.One, p)
		} else {
			s.transmitPulse(p.Zero, p)
		}
	}

	s.transmitPulse(p.Sync, p)
}

func (s *RCSwitch) transmitPulse(b Bit, p Protocol) {
	firstLogicLevel := gpio.High
	if p.InvertedSignal {
		firstLogicLevel = gpio.Low
	}

	secondLogicLevel := gpio.Low
	if p.InvertedSignal {
		secondLogicLevel = gpio.High
	}

	s.pin.Out(firstLogicLevel)
	time.Sleep(p.PulseLength * time.Duration(b.High))

	s.pin.Out(secondLogicLevel)
	time.Sleep(p.PulseLength * time.Duration(b.Low))
}
