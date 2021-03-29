package main

import (
	"bytes"
	"fmt"
	"godogs/helpers"
	"os"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

func main() {

	buf := new(bytes.Buffer)

	// Selects folder (Scenario) to run (info comming from queue)
	opts := godog.Options{
		Output: colors.Colored(buf),
		Format: "pretty",
		Paths:  []string{"taxes"},
	}

	// Run tests
	status := helpers.RunTests(opts)

	// Save this output to database
	fmt.Println(buf.String())

	os.Exit(status)
}
