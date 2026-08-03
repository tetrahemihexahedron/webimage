package image

import "strings"

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
