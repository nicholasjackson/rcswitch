package rcswitch

import (
	"time"

	"periph.io/x/periph/conn/gpio"
)

const maxChanges = 67
const separationLimit = 4200 // microseconds
const receiveTolerance = 60  // microseconds

// Bit represents the pulses for a bit
type Bit struct {
	High int64
	Low  int64
}

// Protocol represents the protocol for the RF controller
type Protocol struct {
	ID             int
	PulseLength    time.Duration
	Sync           Bit
	Zero           Bit
	One            Bit
	InvertedSignal bool
}

//go:generate moq -out switch_moq.go . Switch

// Switch defines the methods for interacting with a 433MHz switch
type Switch interface {
	Scan()
	Send(data string, protocolID int)
}

// RCSwitch is a structure for interacting with 433MHz remote switches
type RCSwitch struct {
	pin gpio.PinIO
}

// New returns a new RCSwitch using the given pin
func New(p gpio.PinIO) *RCSwitch {
	return &RCSwitch{p}
}
