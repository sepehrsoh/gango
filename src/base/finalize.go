package base

import (
	"gango/utils"
	"os/exec"
)

type FinalizeProject struct {
}

func (b *FinalizeProject) Run(name string) {
	projectName := utils.GetProjectDir(name)
	downloadPackages(projectName)
	createDockerfile(projectName)
}

func downloadPackages(name string) {
	cmd := exec.Command("make", "dependencies")
	cmd.Dir = name
	err := cmd.Run()
	if err != nil {
		logger.Panic(err)
	}
}

type DockerfileWriter struct {
}

func (d DockerfileWriter) TemplateName() string {
	return "dockerFile.tmpl"
}

func (d DockerfileWriter) TemplateData(name string) map[string]interface{} {
	return map[string]interface{}{
		"ProjectName": name,
	}
}

func (d DockerfileWriter) FilePath() string {
	return ""
}

func (d DockerfileWriter) FileName() string {
	return "Dockerfile"
}

func createDockerfile(name string) {
	err := utils.EnrichTemplate(name, DockerfileWriter{})
	if err != nil {
		logger.Panic(err)
	}
}
