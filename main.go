package main

import (
	// Ensure all dependencies are imported for go generate.
	_ "github.com/Masterminds/sprig/v3"
	_ "github.com/iancoleman/strcase"
	_ "golang.org/x/exp/slices"
	_ "gopkg.in/yaml.v3"
)

//go:generate go run template/template.go
