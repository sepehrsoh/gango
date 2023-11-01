package utils

import (
	"gango/lib"
	"os"
	"path/filepath"
	"text/template"
)

func EnrichTemplate(dir string, file lib.IWriteTemplate) error {
	templateFile := file.TemplateName();
	templateContent, err := os.ReadFile(filepath.Join("templates", templateFile))
	if err != nil {
		return err
	}
	tmpl, err := template.New(templateFile).Parse(string(templateContent))
	if err != nil {
		return err
	}

	output, err := os.Create(filepath.Join(dir, file.FilePath(), file.FileName()))
	if err != nil {
		return err
	}
	defer output.Close()

	err = tmpl.Execute(output, file.TemplateData(dir))
	if err != nil {
		return err
	}
	return nil
}