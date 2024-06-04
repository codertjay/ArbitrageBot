package main

import (
	newConfiguration "ArbitrageBot/ArbitrageBot/newConfigurations"
	"log"
)

func main() {
	cfg := &newConfiguration.Config{}
	_, err := cfg.Setup()
	if err != nil {
		log.Fatalf("Failed to set up config: %v", err)
	}
}
