package rcswitch

import (
	"fmt"
	"math"
)

func diff(a int64, b int64) int64 {
	return int64(math.Abs(float64(a) - float64(b)))
}

func getSyncLenghtInPulses(p Protocol) int64 {
	if p.Sync.Low > p.Sync.High {
		return p.Sync.Low
	}

	return p.Sync.High
}

func findProtocol(id int) (Protocol, error) {
	for _, p := range protocols {
		if p.ID == id {
			return p, nil
		}
	}

	return Protocol{}, fmt.Errorf("Unable to find protocol")
}
