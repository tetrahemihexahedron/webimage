package variants

type Spec struct {
	FileName string
	Format   string
	Width    uint
	Height   uint
}

func Generate(source string, specs []Spec) error {
	return nil
}
