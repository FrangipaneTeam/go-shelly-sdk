//go:build ignore

// go generate
package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

type configFile struct {
	Commands []templateCommandInfos `yaml:"commands"`
}

type templateCommandInfos struct {
	Category    string                                      `yaml:"category"`
	Name        string                                      `yaml:"name"`
	Description string                                      `yaml:"description"`
	Request     map[string]templateCommandInfosArgsRequest  `yaml:"request"`
	Response    map[string]templateCommandInfosArgsResponse `yaml:"response"`

	ExtraStructsRequest  map[string]templateCommandInfosArgsRequest
	ExtraStructsResponse map[string]templateCommandInfosArgsResponse
}

type templateCommandInfosArgsRequest struct {
	Type        string `yaml:"type"`
	OmitEmpty   bool   `yaml:"omitempty"`
	Description string `yaml:"description"`
	Items       map[string]templateCommandInfosArgsRequest
}

type templateCommandInfosArgsResponse struct {
	Type        string `yaml:"type"`
	OmitEmpty   bool   `yaml:"omitempty"`
	Description string `yaml:"description"`
	Items       map[string]templateCommandInfosArgsResponse
}

//go:embed cmd.go.tmpl
var templateCommand string

//go:embed clients.go.tmpl
var templateClient string

//go:embed cmd.yaml
var commandsYaml string

func main() {
	fmt.Println("generating commands files...")

	c := &configFile{}

	err := yaml.Unmarshal([]byte(commandsYaml), c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	funcMap := template.FuncMap{
		"ToUpper":      strings.ToUpper,
		"ToLower":      strings.ToLower,
		"ToCamel":      strcase.ToCamel,
		"ToLowerCamel": strcase.ToLowerCamel,
	}

	// fmt.Printf("%+v\n", c)

	commands := make(map[string][]templateCommandInfos)
	clients := []string{}

	for _, v := range c.Commands {
		extraStructsRequest := make(map[string]templateCommandInfosArgsRequest)
		extraStructsResponse := make(map[string]templateCommandInfosArgsResponse)

		if !slices.Contains(clients, v.Category) {
			clients = append(clients, v.Category)
		}

		for k, v := range v.Request {
			if v.Type == "object" {
				extraStructsRequest[k] = v
				for k2, v2 := range v.Items {
					if v2.Type == "object" {
						extraStructsRequest[strcase.ToCamel(fmt.Sprintf("%s_%s", k, k2))] = v2
					}
				}
			}
		}

		for k, v := range v.Response {
			if v.Type == "object" {
				extraStructsResponse[k] = v
				for k2, v2 := range v.Items {
					if v2.Type == "object" {
						extraStructsResponse[strcase.ToCamel(fmt.Sprintf("%s_%s", k, k2))] = v2
					}
				}
			}
		}

		v.ExtraStructsRequest = extraStructsRequest
		v.ExtraStructsResponse = extraStructsResponse

		commands[v.Category] = append(commands[v.Category], v)

	}

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

	for k, v := range commands {
		tmpl, err := template.New("template").Funcs(funcMap).Parse(templateCommand)
		if err != nil {
			fmt.Printf("error from template parse : %v\n", err)
			os.Exit(1)
		}

		var tpl bytes.Buffer

		errExec := tmpl.Execute(&tpl, v)

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

		errWrite := os.WriteFile("shelly/generated_cmd_"+strings.ToLower(k)+".go", formattedContent, 0o644)
		if errWrite != nil {
			fmt.Printf("write to file error : %v\n", errWrite)
		}
	}

	return

}
