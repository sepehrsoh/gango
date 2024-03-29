package base

import (
	"os/exec"

	"gango/utils"
)

type FinalizeProject struct {
}

func (b *FinalizeProject) Run(name string, conf utils.Config) {
	projectName := utils.GetProjectDir(name)
	downloadPackages(projectName)
	createDockerfile(projectName)
	createEnvFile(projectName, conf)
}

func downloadPackages(name string) {
	cmd := exec.Command("make", "dependencies")
	cmd.Dir = name
	err := cmd.Run()
	if err != nil {
		logger.Errorf("download packages error: %v", err)
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
		logger.Errorf("create docker file error: %v", err)
	}
}

type EnvFileWriter struct {
	options utils.ServiceOptions
}

func (e EnvFileWriter) FilePath() string {
	return ""
}

func (e EnvFileWriter) FileName() string {
	return ".env"
}

func (e EnvFileWriter) TemplateName() string {
	return "envFile.tmpl"
}

func (e EnvFileWriter) TemplateData(name string) map[string]interface{} {
	tmpl := utils.GetDefaultTemplateValues(name)
	if e.options.WithRedis {
		tmpl["withRedis"] = true
	}
	if e.options.WithElastic {
		tmpl["withElastic"] = true
	}
	if e.options.WithPostgres {
		tmpl["withPostgres"] = true
	}
	return tmpl
}

func NewEnvFileWriter(configs ...utils.Options) EnvFileWriter {
	opts := utils.DefaultOptions
	for _, config := range configs {
		config.Apply(&opts)
	}
	return EnvFileWriter{options: opts}
}

func createEnvFile(name string, cnf utils.Config) {
	err := utils.EnrichTemplate(name, NewEnvFileWriter(utils.GetWireConfigs(cnf)...))
	if err != nil {
		logger.Errorf("create env file error: %v", err)
	}
}
