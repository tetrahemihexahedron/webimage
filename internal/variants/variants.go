package variants

import (
	"errors"
	"fmt"
)

type Spec struct {
	OutPath string
	Width   uint
}

func Generate(source string, specs []Spec) error {
	if source == "" {
		return errors.New("source file path cannot be empty")
	}

	var errs []error
	for _, spec := range specs {
		if err := generateVariant(source, spec); err != nil {
			errs = append(errs, fmt.Errorf(
				"generating %q with width %d: %w",
				spec.OutPath,
				spec.Width,
				err,
			))
		}
	}
	return errors.Join(errs...)
}

func generateVariant(source string, spec Spec) error {
	return nil
}
