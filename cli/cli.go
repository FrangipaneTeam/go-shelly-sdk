package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/FrangipaneTeam/go-shelly-sdk/shelly"
	"github.com/kr/pretty"
	"gopkg.in/yaml.v3"
)

type DevicesFile struct {
	Devices map[string]device `yaml:"devices"`
}

type device struct {
	Name    string              `yaml:"name"`
	Methods map[string][]string `yaml:"methods"`
}

func main() {
	address := flag.String("address", "", "Address to connect to. (Required)")
	appendToDeviceFile := flag.Bool("append-to-devices", false, "Append to device file.")
	pathFile := flag.String("path-file", "devices.yaml", "Path to file to write to.")
	flag.Parse()

	if *address == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	s, err := shelly.New(*address)
	if err != nil {
		log.Fatal(err)
	}

	deviceInfoResponse, err := s.API().Shelly().GetDeviceInfo(shelly.ShellyGetDeviceInfoRequest{
		Ident: false,
	})
	if err != nil {
		log.Fatal(err)
	}

	d := &device{
		Name:    deviceInfoResponse.App,
		Methods: make(map[string][]string),
	}

	listMethods, err := s.API().Shelly().ListMethods(shelly.ShellyListMethodsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	for _, method := range listMethods.Methods {
		m := strings.Split(method, ".")
		d.Methods[m[0]] = append(d.Methods[m[0]], m[1])
	}

	out, err := yaml.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(out))

	if *appendToDeviceFile {
		devices := &DevicesFile{}

		f, err := os.OpenFile(*pathFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		buf, _ := ioutil.ReadAll(f)

		err = yaml.Unmarshal(buf, devices)
		if err != nil {
			log.Fatal(err)
		}

		if devices.Devices == nil {
			devices.Devices = make(map[string]device)
		}

		if _, ok := devices.Devices[d.Name]; ok {
			log.Fatalf("Device %s already exists in file %s", d.Name, *pathFile)
		}

		devices.Devices[d.Name] = device{
			Methods: d.Methods,
			Name:    d.Name,
		}

		out, err := yaml.Marshal(devices)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := f.Write(out); err != nil {
			log.Fatal(err)
		}

		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	} else {
		pretty.Print(d)
	}
}
