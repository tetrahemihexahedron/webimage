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
	output, fetchErr := fetchExiftoolOutput(path)

	metadata, errs := processOutput(output)
	if fetchErr != nil {
		errs = append(errs, fetchErr)
	}

	return metadata, errs
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

	rawOutput, commandErr := cmd.Output()

	if len(rawOutput) == 0 {
		if commandErr == nil {
			return nil, fmt.Errorf("running exiftool on %q returned no output and no error", path)
		}
		if exitErr, ok := errors.AsType[*exec.ExitError](commandErr); ok {
			return nil, fmt.Errorf("running exiftool on %q: %s", path, exitErr.Stderr)
		}
		return nil, fmt.Errorf("running exiftool on %q: %w", path, commandErr)
	}

	// if there was output to stdout, then commandErr is probably not interesting
	// because exiftool reports the reasons for errors in stdout
	var output []exiftoolOutput
	if err := json.Unmarshal(rawOutput, &output); err != nil {
		return output, fmt.Errorf("unmarshalling exiftool output for %q: %w", path, err)
	}
	return output, nil
}

func processOutput(output []exiftoolOutput) ([]ImageMetadata, []error) {
	metadata := make([]ImageMetadata, 0, len(output))
	var errs []error

	for _, out := range output {
		if out.Error != "" {
			errs = append(errs, fmt.Errorf(
				"reported by exiftool for %q: %s",
				out.FileName,
				out.Error,
			))
			continue
		}
		if out.FileName == "" ||
			out.FileType == "" ||
			out.Width == 0 ||
			out.Height == 0 {
			errs = append(errs, fmt.Errorf(
				"missing required metadata in %+v: FileName, FileType, Width, and Height are required",
				out,
			))
			continue
		}

		metadata = append(metadata, ImageMetadata{
			FileName:    out.FileName,
			FileType:    out.FileType,
			Title:       out.Title,
			Description: out.Description,
			Width:       out.Width,
			Height:      out.Height,
		})
	}

	return metadata, errs
}
