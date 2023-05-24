package main

import (
	"log"

	"github.com/FrangipaneTeam/go-shelly-sdk/shelly"
)

func main() {
	s, err := shelly.New("192.168.0.161")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.Devices().Pro4PM().Switch().GetStatus(shelly.SwitchGetStatusRequest{
		Id: 0,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v", resp)
}
