package rcswitch

import (
	"log"
	"time"
)

// Scan scans for button presses from the remote
func (r *RCSwitch) Scan() {
	lastTime := time.Now()
	changeCount := 0
	repeatCount := 0
	changes := make([]int64, maxChanges)

	// Wait for edges as detected by the hardware:
	for r.pin.WaitForEdge(-1) {

		dur := time.Now().Sub(lastTime).Nanoseconds() / 1000 // duration in microseconds

		if dur >= separationLimit {
			// A long stretch without signal level change, this could be the gap
			// between two transmissions.
			if diff(dur, changes[0]) < 200 {
				// This long signal is close in length to the long signal which
				// started the previously recorded timings.  This suggests that it may
				// indeed be a gap between two transmissions.  We assume that the sender
				// will send the signal multiple times with roughly the same gap
				repeatCount++
				if repeatCount == 2 {

					for _, p := range protocols {
						if processChange(changes, changeCount, p) {
							// receive successfull
							break
						}
					}

					repeatCount = 0
				}
			}

			changeCount = 0
		}

		if changeCount >= maxChanges {
			changeCount = 0
			repeatCount = 0
		}

		changes[changeCount] = dur
		changeCount++
		lastTime = time.Now()
	}
}

func processChange(c []int64, changeCount int, p Protocol) bool {
	code := 0
	syncLengthInPulses := getSyncLenghtInPulses(p)
	delay := c[0] / syncLengthInPulses
	delayTolerance := delay * receiveTolerance / 100

	firstDataTiming := 1
	if p.InvertedSignal {
		firstDataTiming = 2
	}

	for i := firstDataTiming; i < changeCount-1; i += 2 {
		code <<= 1

		if diff(c[i], (delay*p.Zero.High)) < delayTolerance &&
			diff(c[i+1], (delay*p.Zero.Low)) < delayTolerance {

		} else if diff(c[i], (delay*p.One.High)) < delayTolerance &&
			diff(c[i+1], (delay*p.One.Low)) < delayTolerance {
			code |= 1
		} else {
			return false
		}
	}

	if changeCount > 7 {
		log.Printf("decimal: %d\n", code)
		log.Printf("binary: %024b\n", code)
		log.Printf("bitlength: %d", (changeCount-1)/2)
		log.Printf("protocol: %d", p.ID)
		log.Println("")
		return true
	}

	return false
}
