package base

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gango/lib"
	"gango/utils"
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
		fmt.Printf("There exist folder with name %v\nAre your sure to rewrite this folder?\n[y/n]: ", name)
		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			logger.Panic(err)
		}
		if strings.ToLower(strings.TrimSpace(response)) == "y" {
			_ = os.RemoveAll(name)
			makeMainDirectory(name)
			return
		}
		logger.Panic("process stopped")
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
	return utils.GetDefaultTemplateValues(splitName[len(splitName)-1])
}

func (m LinterWriter) FilePath() string {
	return ""
}

func (m LinterWriter) FileName() string {
	return ".golang-ci.yaml"
}

func createLinter(name string) {
	err := utils.EnrichTemplate(name, LinterWriter{})
	if err != nil {
		logger.Panic(err)
	}
}
