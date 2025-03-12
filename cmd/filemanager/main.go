package main

import (
	"fmt"
	"os"

	"filemanager/internal/config"
	"filemanager/internal/operations"
)

func main() {
	cfg, err := config.ParseFlags()
	if err != nil {
		fmt.Println(err)
		config.PrintUsage()
		os.Exit(1)
	}

	err = operations.Execute(cfg)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}