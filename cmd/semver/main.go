package main

import (
	"flag"
	"fmt"
	"github.com/maykonlf/semver-cli/internal/entities"
)

func main() {
	versionStr := flag.String("v", "", "version which next version should be based on")
	flag.Parse()

	commands := flag.Args()
	fmt.Println(commands)

	version, err := entities.NewVersion(*versionStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(version.String())
}
