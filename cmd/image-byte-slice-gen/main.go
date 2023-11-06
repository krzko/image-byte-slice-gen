package main

import (
	"fmt"
	"os"

	"github.com/krzko/byte-slice-gen/pkg/cli"
)

func main() {

	app := cli.SetupApp()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}

}
