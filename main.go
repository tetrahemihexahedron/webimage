package main

import (
	"log"
	"tetrahemihexahedron/webimage/internal/config"
	"tetrahemihexahedron/webimage/internal/processor"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded config: %+v", config)

	processor.ProcessDir(config)
}
