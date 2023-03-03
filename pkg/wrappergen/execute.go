package wrappergen

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"go.fd.io/govpp/binapigen"
)

type Template struct {
	templates map[string]*template.Template
	input     fs.FS
}

func ParseFS(input fs.FS, patterns ...string) (*Template, error) {
	rv := &Template{
		input:     input,
		templates: make(map[string]*template.Template),
	}
	var err error
	fs.WalkDir(rv.input, ".", rv.addAllToTemplateWalkFn)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (t *Template) addAllToTemplateWalkFn(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if d.IsDir() {
		return nil
	}
	if !strings.HasSuffix(d.Name(), ".tmpl") {
		return nil
	}
	tmpl, err := template.ParseFS(t.input, path)
	if err != nil {
		return err
	}
	t.templates[path] = tmpl
	return nil
}

func (t *Template) createExecuteWalkFn(outputDir string, data interface{}, gen *binapigen.Generator) func(path string, d fs.DirEntry, err error) error {
	return func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(d.Name(), ".tmpl") {
			return nil
		}

		outputBuffer := bytes.NewBuffer([]byte{})
		tmpl, ok := t.templates[path]
		if !ok {
			return nil
		}
		if err := tmpl.Execute(outputBuffer, data); err != nil {
			return err
		}

		if strings.TrimSpace(outputBuffer.String()) == "" {
			return nil
		}

		outputPath := filepath.Join(outputDir, path)

		if err := os.MkdirAll(filepath.Dir(outputPath), 0700); err != nil {
			return err
		}

		genFile := gen.NewGenFile(strings.TrimSuffix(outputPath, ".tmpl"), nil)
		_, err = outputBuffer.WriteTo(genFile)

		return err
	}
}

func (t *Template) ExecuteAll(outputDir string, data interface{}, gen *binapigen.Generator) error {
	return fs.WalkDir(t.input, ".", t.createExecuteWalkFn(outputDir, data, gen))
}
