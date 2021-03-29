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

	opts := godog.Options{
		Output: colors.Colored(buf),
		Format: "pretty",
		Paths:  []string{"taxes"},
	}

	status := helpers.RunTests(opts)

	fmt.Println(buf.String())

	os.Exit(status)
}
