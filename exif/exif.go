package exif

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
