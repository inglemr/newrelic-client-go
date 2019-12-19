// +build tools

package main

import (
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/stretchr/testify/assert"
	_ "golang.org/x/tools/cmd/godoc"
)
