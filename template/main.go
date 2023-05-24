//go:build ignore

// go generate
package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/FrangipaneTeam/go-shelly-sdk/internal/tools"
	"github.com/Masterminds/sprig/v3"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

type devicesFile struct {
	Devices map[string]device `yaml:"devices"`
}

type device struct {
	Name    string              `yaml:"name"`
	Methods map[string][]string `yaml:"methods"`
}

type configFile struct {
	Commands []*templateCommandInfos `yaml:"commands"`
}

type templateCommandInfos struct {
	Category    string                                       `yaml:"category"`
	Name        string                                       `yaml:"name"`
	Description string                                       `yaml:"description"`
	Request     map[string]*templateCommandInfosArgsRequest  `yaml:"request"`
	Response    map[string]*templateCommandInfosArgsResponse `yaml:"response"`

	ExtraStructsRequest  map[string]*templateCommandInfosArgsRequest
	ExtraStructsResponse map[string]*templateCommandInfosArgsResponse
}

type templateCommandInfosArgsRequest struct {
	StructName  string
	CamelName   string
	LowerName   string
	Type        string `yaml:"type"`
	OmitEmpty   bool   `yaml:"omitempty"`
	Description string `yaml:"description"`
	Items       map[string]*templateCommandInfosArgsRequest
}

type templateCommandInfosArgsResponse struct {
	StructName  string
	CamelName   string
	LowerName   string
	Type        string `yaml:"type"`
	OmitEmpty   bool   `yaml:"omitempty"`
	Description string `yaml:"description"`
	Items       map[string]*templateCommandInfosArgsResponse
}

//go:embed cmd.go.tmpl
var templateCommand string

//go:embed clients.go.tmpl
var templateClient string

//go:embed devices.go.tmpl
var templateDevices string

//go:embed cmd.yaml
var commandsYaml string

//go:embed devices.yaml
var devicesYaml string

func main() {
	fmt.Println("generating commands files...")

	var (
		c        = &configFile{}
		d        = &devicesFile{}
		commands = make(map[string][]*templateCommandInfos)
		clients  = []string{}
		err      error
	)

	// Parse yaml commands
	if err = yaml.Unmarshal([]byte(commandsYaml), c); err != nil {
		log.Fatalf("error: %v", err)
	}

	// Parse yaml devices
	if err = yaml.Unmarshal([]byte(devicesYaml), d); err != nil {
		log.Fatalf("error: %v", err)
	}

	// Custom functions for template
	funcMap := template.FuncMap{
		"ToUpper":      strings.ToUpper,
		"ToLower":      strings.ToLower,
		"ToCamel":      strcase.ToCamel,
		"ToLowerCamel": strcase.ToLowerCamel,
	}

	// For each command construct a map with the category as key and the command as value
	for _, v := range c.Commands {

		if !slices.Contains(clients, v.Category) {
			clients = append(clients, v.Category)
		}

		v.ExtraStructsRequest = make(map[string]*templateCommandInfosArgsRequest)
		v.ExtraStructsResponse = make(map[string]*templateCommandInfosArgsResponse)

		// Request
		for kRequest, vRequest := range v.Request {
			vRequest.StructName = strcase.ToCamel(fmt.Sprintf("%s_%s_request_%s", v.Category, v.Name, kRequest))
			vRequest.CamelName = strcase.ToCamel(kRequest)
			vRequest.LowerName = strcase.ToLowerCamel(kRequest)
			if strings.HasSuffix(vRequest.Type, "object") {
				v.ExtraStructsRequest[vRequest.StructName] = vRequest
				x := recursiveRequestItems(vRequest.StructName, vRequest.Items)
				for _, value := range x {
					v.ExtraStructsRequest[value.StructName] = value
				}
			}
			v.Request[kRequest] = vRequest
		}

		// Response
		for kResponse, vResponse := range v.Response {
			vResponse.StructName = strcase.ToCamel(fmt.Sprintf("%s_%s_response_%s", v.Category, v.Name, kResponse))
			vResponse.CamelName = strcase.ToCamel(kResponse)
			vResponse.LowerName = strcase.ToLowerCamel(kResponse)
			if strings.HasSuffix(vResponse.Type, "object") {
				v.ExtraStructsResponse[vResponse.StructName] = vResponse
				x := recursiveResponseItems(vResponse.StructName, vResponse.Items)
				for _, value := range x {
					v.ExtraStructsResponse[value.StructName] = value
				}
			}
			v.Response[kResponse] = vResponse
		}

		commands[v.Category] = append(commands[v.Category], v)
	}

	// Create template for clients
	tmpl, err := template.New("template").Funcs(sprig.FuncMap()).Funcs(funcMap).Parse(templateClient)
	if err != nil {
		fmt.Printf("error from template parse : %v\n", err)
		os.Exit(1)
	}

	var tpl bytes.Buffer

	errExec := tmpl.Execute(&tpl, clients)

	if errExec != nil {
		fmt.Printf("error from template execute : %v\n", errExec)
		os.Exit(1)
	}

	// format the code
	formattedContent, errFormat := format.Source(tpl.Bytes())
	if errFormat != nil {
		fmt.Printf("error from format : %v\n", errFormat)
		os.Exit(1)
	}

	errWrite := os.WriteFile("shelly/generated_clients.go", formattedContent, 0o644)
	if errWrite != nil {
		fmt.Printf("write to file error : %v\n", errWrite)
	}

	// End template for clients

	// Create template for commands
	for k, v := range commands {

		tmpl, err := template.New("template").Funcs(funcMap).Parse(templateCommand)
		if err != nil {
			fmt.Printf("error from template parse : %v\n", err)
			os.Exit(1)
		}

		var tpl bytes.Buffer

		if err := tmpl.Execute(&tpl, v); err != nil {
			fmt.Printf("error from template execute : %v\n", errExec)
			os.Exit(1)
		}

		// format the code
		formattedContent, errFormat := format.Source(tpl.Bytes())
		if errFormat != nil {
			fmt.Printf("error from format for component %s : %v\n", k, errFormat)
			os.Exit(1)
		}

		errWrite := os.WriteFile("shelly/generated_cmd_"+strings.ToLower(k)+".go", formattedContent, 0o644)
		if errWrite != nil {
			fmt.Printf("write to file error : %v\n", errWrite)
		}
	}
	// End template for commands

	// Create template for devices
	for k, v := range d.Devices {

		for i, cmds := range v.Methods {
			xx := []string{}
			if !tools.MapItemExists(commands, i) {
				// Remove method if not in commands
				delete(v.Methods, i)
				continue
			}
			for _, c := range commands[i] {
				if slices.Contains(cmds, c.Name) {
					xx = append(xx, c.Name)
				}
			}

			v.Methods[i] = xx
		}

		tmpl, err = template.New("template").Funcs(funcMap).Parse(templateDevices)
		if err != nil {
			fmt.Printf("error from template devices parse : %v\n", err)
			os.Exit(1)
		}

		var tpl bytes.Buffer

		if err := tmpl.Execute(&tpl, v); err != nil {
			fmt.Printf("error from template devices execute : %v\n", errExec)
			os.Exit(1)
		}

		// format the code
		formattedContent, errFormat := format.Source(tpl.Bytes())
		if errFormat != nil {
			fmt.Printf("error from format for component %s : %v\n", k, errFormat)
			os.Exit(1)
		}

		if err := os.WriteFile("shelly/generated_device_"+strings.ToLower(k)+".go", formattedContent, 0o644); err != nil {
			fmt.Printf("write to file error : %v\n", err)
		}
	}
	// End template for devices
}

func recursiveRequestItems(base string, items map[string]*templateCommandInfosArgsRequest) map[string]*templateCommandInfosArgsRequest {
	x := make(map[string]*templateCommandInfosArgsRequest)

	for k, v := range items {
		v.CamelName = strcase.ToCamel(k)
		v.LowerName = strcase.ToLowerCamel(k)
		v.StructName = strcase.ToCamel(fmt.Sprintf("%s_%s", base, k))

		if strings.HasSuffix(v.Type, "object") {
			x[k] = v
			x := recursiveRequestItems(base, v.Items)
			for _, v2 := range x {
				if strings.HasSuffix(v2.Type, "object") {
					items[v2.StructName] = v2
				}
			}
		}
	}
	return x
}

func recursiveResponseItems(base string, items map[string]*templateCommandInfosArgsResponse) map[string]*templateCommandInfosArgsResponse {
	x := make(map[string]*templateCommandInfosArgsResponse)

	for k, v := range items {
		v.CamelName = strcase.ToCamel(k)
		v.LowerName = strcase.ToLowerCamel(k)
		v.StructName = strcase.ToCamel(fmt.Sprintf("%s_%s", base, k))

		if strings.HasSuffix(v.Type, "object") {
			x[k] = v
			x := recursiveResponseItems(base, v.Items)
			for _, v2 := range x {
				if strings.HasSuffix(v2.Type, "object") {
					items[v2.StructName] = v2
				}
			}
		}
	}
	return x
}
