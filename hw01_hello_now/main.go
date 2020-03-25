package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	fmt.Printf("current time: %v\n", time.Now())

	ntpTime, err := ntp.Time("time.apple.com")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("exact time: %v\n", ntpTime)
}
