package main

import (
	"log"
	"tetrahemihexahedron/webimage/internal/config"
	"tetrahemihexahedron/webimage/internal/exif"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded config: %+v", config)

	allMetadata, errs := exif.FetchMetadata(config.InDir)

	log.Printf("Fetched metadata for %d files. %d error(s).", len(allMetadata), len(errs))
	for _, err = range errs {
		log.Printf("Metadata error: %v", err)
	}
}
