package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path/filepath"
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

	for _, metadata := range allMetadata {
		if metadata.FileType != "JPEG" {
			log.Printf("Skipping file %s: file type is %s, not JPEG", metadata.FileName, metadata.FileType)
			continue
		}

		inPath := filepath.Join(config.InDir, metadata.FileName)
		hash, err := hashFile(inPath)
		if err != nil {
			log.Fatal(err)
		}
		shortHash := hash[:10]
		log.Printf("short hash for %s is %s", inPath, shortHash)
	}
}

func hashFile(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
