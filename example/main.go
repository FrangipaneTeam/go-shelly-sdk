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

	deviceInfoResponse, err := s.API().Shelly().GetDeviceInfo(shelly.ShellyGetDeviceInfoRequest{
		Ident: false,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v", deviceInfoResponse)

	listMethods, err := s.API().Shelly().ListMethods(shelly.ShellyListMethodsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v", listMethods)

	s.API().Switch().GetConfig(shelly.SwitchGetConfigRequest{})

}
