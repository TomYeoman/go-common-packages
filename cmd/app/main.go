package main

import (
	num1 "github.com/TomYeoman/go-common-packages/pkg/formatting/numbers"
	num2 "github.com/TomYeoman/go-common-packages/pkg/formatting/numbers"
	"github.com/TomYeoman/go-common-packages/pkg/formatting/string"
	"github.com/TomYeoman/go-common-packages/pkg/logging"
)

func main() {
	logging.PrintString("yaya")

	// Note how a single folder can only have a single package import
	// of which we can call whatever we like
	num1.Add(10, 50)
	num1.Multiply(10, 60)
	num2.Add(10, 50)
	num2.Multiply(10, 60)

	string.Length("Hiii")

}
