package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pedrfelip/go-gateway/gateway/config"
)

func main() {
	configPath := os.Args[1]

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.LoadConfig(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
