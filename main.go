package main

import (
	"fmt"
	"infra-lab-cli/cmd"
	"infra-lab-cli/config"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}

	cmd.Execute()
}
