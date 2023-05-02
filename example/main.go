package main

import (
	"log"

	"github.com/FrangipaneTeam/go-shelly-sdk/shelly"
)

func main() {
	s, err := shelly.New("192.168.0.1")
	if err != nil {
		log.Fatal(err)
	}

	deviceInfoResponse, err := s.Shelly().GetDeviceInfo(shelly.ShellyGetDeviceInfoRequest{})
	if err != nil {
		s.Close()
		log.Fatal(err)
	}

	log.Printf("%#v", deviceInfoResponse)
	s.Close()
}
