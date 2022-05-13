package sampler

import (
	"math"
	"time"
)

func CreateValueAtUnixTimestamp() float64 {

	x := float64(time.Now().Unix())
	return math.Sin(x)

}
