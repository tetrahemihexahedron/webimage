package exif

import (
	"slices"
	"testing"
)

func TestFetchMetadata(t *testing.T) {
	var tests = map[string]struct {
		path            string
		desiredMetadata []ImageMetadata
	}{
		"IMG_3916.jpeg": {
			path: "../test_images/IMG_3916.jpeg",
			desiredMetadata: []ImageMetadata{
				{
					FileName:    "IMG_3916.jpeg",
					FileType:    "JPEG",
					Title:       "2024 August After kleenex destruction",
					Description: "A fluffy Rosie, looking innocent after having shredded a kleenex lying nearby",
					Width:       4032,
					Height:      3024,
				},
			},
		},
		"riveter_chew_it_1024x1535.jpeg": {
			path: "../test_images/riveter_chew_it_1024x1535.jpeg",
			desiredMetadata: []ImageMetadata{
				{
					FileName:    "riveter_chew_it_1024x1535.jpeg",
					FileType:    "JPEG",
					Title:       "AIgen We can chew it poster",
					Description: "A parody of the classic Rosie the Riveter poster, with a poodle in a red bandana saying \"We can chew it\"",
					Width:       1024,
					Height:      1535,
				},
			},
		},
		"squash.jpg": {
			path: "../test_images/squash.jpg",
			desiredMetadata: []ImageMetadata{
				{
					FileName:    "squash.jpg",
					FileType:    "JPEG",
					Title:       "",
					Description: "",
					Width:       1200,
					Height:      799,
				},
			},
		},
	}

	for testname, testdata := range tests {
		t.Run(testname, func(t *testing.T) {
			metadata, errs := FetchMetadata(testdata.path)
			if len(errs) != 0 {
				t.Errorf("Unexpected errors returned when fetching metadata for %q: %v", testdata.path, errs)
			}
			if !slices.Equal(metadata, testdata.desiredMetadata) {
				t.Errorf("Image metadata incorrect for source file %q.\n\n   Got: %+v\n\n   Wanted: %+v", testdata.path, metadata, testdata.desiredMetadata)
			}
		})
	}
}
