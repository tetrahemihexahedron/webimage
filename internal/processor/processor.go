package processor

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"tetrahemihexahedron/webimage/internal/config"
	"tetrahemihexahedron/webimage/internal/exif"
	"tetrahemihexahedron/webimage/internal/image"
	"time"
)

func ProcessDir(config config.Config) error {
	inDir := config.InDir
	allMetadata, errs := exif.FetchMetadata(inDir)

	log.Printf("Fetched metadata for %d files. %d error(s).", len(allMetadata), len(errs))
	for _, err := range errs {
		log.Printf("Metadata error: %v", err)
	}

	for _, metadata := range allMetadata {
		if metadata.FileType != image.FormatJPEG {
			log.Printf("Skipping file %s: file type is %s, not JPEG", metadata.FileName, metadata.FileType)
			continue
		}

		inPath := filepath.Join(inDir, metadata.FileName)
		hash, err := hashFile(inPath)
		if err != nil {
			log.Fatal(err)
		}

		image := image.Processed{
			Hash:     hash,
			ImageDir: imageDir(),
			Metadata: metadata,
		}

		log.Printf("image is %+v", image)
	}
	return nil
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

func imageDir() string {
	randId := ""
	b := make([]byte, 7)
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)

	for !re.MatchString(randId) {
		if _, err := rand.Read(b); err != nil {
			log.Fatal(err)
		}
		randId = base64.RawURLEncoding.EncodeToString(b)
	}

	year := time.Now().Year()
	month := time.Now().Month()

	return fmt.Sprintf("%d/%02d/%s/", year, month, randId)
}
