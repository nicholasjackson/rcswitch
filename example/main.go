package main

import (
	"flag"
	"log"
	"time"

	"github.com/nicholasjackson/rcswitch"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

var pin = flag.String("pin", "0", "Number of the pin the transmitter or receiver is attached to")
var scan = flag.Bool("scan", false, "Scan for remote codes")
var protocolID = flag.Int("protocol", 0, "Protocol id to use")
var onCode = flag.String("oncode", "000000000100000000010101", "Binary code for turning the switch on")
var offCode = flag.String("offcode", "000000000100000000010100", "Binary code for turning the switch off")

// Switch values should be 24 bits
// Protocol: 0
// Off: 000000000100000000010100
// On:  000000000100000000010101

func main() {
	flag.Parse()

	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Lookup a pin by its number:
	p := gpioreg.ByName(*pin)
	setPin(*scan, p)
	log.Printf("%s: %s\n", p, p.Function())

	sw := rcswitch.New(p)

	// scan for codes
	if *scan {
		log.Println("Starting code detection, press a button on your remote control")
		sw.Scan()
		return
	}

	// send codes
	sw.Send(*onCode, *protocolID)
	time.Sleep(1 * time.Second)
	sw.Send(*offCode, *protocolID)
}

// sets the pin to either input or output depending on if we are scanning
func setPin(scan bool, p gpio.PinIO) {
	if scan {
		// Set it as input, with an internal pull down resistor:
		if err := p.In(gpio.PullDown, gpio.BothEdges); err != nil {
			log.Fatal(err)
		}

		return
	}

	if err := p.Out(gpio.High); err != nil {
		log.Fatal(err)
	}
}
