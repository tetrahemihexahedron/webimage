package image

import "strings"

type Metadata struct {
	FileName    string
	FileType    Format
	Title       string
	Description string
	Width       int
	Height      int
}

type Processed struct {
	Hash     string
	ImageDir string
	Metadata Metadata
}

type Format string

const (
	FormatJPEG  Format = "JPEG"
	FormatAVIF  Format = "AVIF"
	FormatOther Format = "OTHER"
)

func ParseFormat(s string) Format {
	switch strings.ToLower(s) {
	case "jpeg", "jpg":
		return FormatJPEG
	case "avif":
		return FormatAVIF
	default:
		return FormatOther
	}
}
