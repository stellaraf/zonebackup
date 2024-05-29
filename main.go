package main

import (
	"fmt"
	"os"

	"github.com/stellaraf/zonebackup/internal/cmd"
)

func main() {
	if err := cmd.CLI.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
