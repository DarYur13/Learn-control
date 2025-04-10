package docsgenerator

type docxGenerator struct {
	templateDir string
}

func New(templateDir string) DocsGenerator {
	return &docxGenerator{
		templateDir: templateDir,
	}
}
