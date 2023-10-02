package main

import (
	"fmt"
	"github.com/leeturner/wiremock-cli/cmd"
	"os"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
