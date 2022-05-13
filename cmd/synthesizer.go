package main

import (
	"fmt"
	"time"

	"github.com/IcebergLettuce/DataSynthesizer/internal/sampler"
)

func main() {

	generate := true

	for generate {

		value := sampler.CreateValueAtUnixTimestamp()
		fmt.Printf("generated value %f\n", value)
		time.Sleep(time.Second)
	}

}
