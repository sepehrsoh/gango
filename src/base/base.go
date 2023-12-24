package base

import (
	"gango/lib"
	"gango/utils"
	"os"
	"os/exec"
	"strings"
)

var logger = lib.GetLogger("base")

type BaseProject struct {
}

func (b *BaseProject) Run(name string) {
	projectName := utils.GetProjectDir(name)
	makeMainDirectory(projectName)
	createGoMod(name)
	createMakefile(projectName)
	createReadMeFile(projectName)
	createLinter(projectName)
}

func makeMainDirectory(name string) {
	err := os.Mkdir(name, 0750)
	if err != nil {
		logger.Panic(err)
	}
}

func createGoMod(name string) {
	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = utils.GetProjectDir(name)
	err := cmd.Run()
	if err != nil {
		logger.Panic(err)
	}
}

type MakefileWriter struct{}

func (m MakefileWriter) TemplateName() string {
	return "makefile.tmpl"
}

func (m MakefileWriter) TemplateData(name string) map[string]interface{} {
	splitName := strings.Split(name, "/")
	return map[string]interface{}{
		"ProjectName": splitName[len(splitName)-1],
	}
}

func (m MakefileWriter) FilePath() string {
	return ""
}

func (m MakefileWriter) FileName() string {
	return "Makefile"
}

func createMakefile(name string) {
	err := utils.EnrichTemplate(name, MakefileWriter{})
	if err != nil {
		logger.Panic(err)
	}
}

type ReadmeWriter struct {
}

func (r ReadmeWriter) TemplateName() string {
	return "readmeFile.tmpl"
}

func (r ReadmeWriter) TemplateData(name string) map[string]interface{} {
	return map[string]interface{}{
		"ProjectName": name,
	}
}

func (r ReadmeWriter) FilePath() string {
	return ""
}

func (r ReadmeWriter) FileName() string {
	return "README.md"
}

func createReadMeFile(name string) {
	err := utils.EnrichTemplate(name, ReadmeWriter{})

	if err != nil {
		logger.Panic(err)
	}
}

type LinterWriter struct{}

func (m LinterWriter) TemplateName() string {
	return "linter.tmpl"
}

func (m LinterWriter) TemplateData(name string) map[string]interface{} {
	splitName := strings.Split(name, "/")
	return map[string]interface{}{
		"ProjectName": splitName[len(splitName)-1],
	}
}

func (m LinterWriter) FilePath() string {
	return ""
}

func (m LinterWriter) FileName() string {
	return ".golangci.yaml"
}

func createLinter(name string) {
	err := utils.EnrichTemplate(name, LinterWriter{})
	if err != nil {
		logger.Panic(err)
	}
}
