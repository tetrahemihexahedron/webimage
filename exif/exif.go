package exif

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
)

type ImageMetadata struct {
	FileName    string
	FileType    string
	Title       string
	Description string
	Width       uint
	Height      uint
}

func FetchMetadata(path string) ([]ImageMetadata, []error) {
	return []ImageMetadata{}, nil
}

type exiftoolOutput struct {
	FileName    string `json:"FileName"`
	FileType    string `json:"FileType"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Width       uint   `json:"ImageWidth"`
	Height      uint   `json:"ImageHeight"`
	Error       string `json:"Error"`
}

func fetchExiftoolOutput(path string) ([]exiftoolOutput, error) {
	cmd := exec.Command("exiftool", "-json", path)

	out, err := cmd.Output()

	if len(out) == 0 {
		if exitErr, ok := errors.AsType[*exec.ExitError](err); ok {
			return nil, fmt.Errorf("running exiftool on %q: %s", path, exitErr.Stderr)
		}
		return nil, fmt.Errorf("running exiftool on %q: %w", path, err)
	}

	var structuredOutput []exiftoolOutput
	if err := json.Unmarshal(out, &structuredOutput); err != nil {
		return structuredOutput, fmt.Errorf("unmarshalling exiftool output for %q: %w", path, err)
	}
	return structuredOutput, nil
}
