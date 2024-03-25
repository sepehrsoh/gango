package providers

import "gango/utils"

type Grpc struct {
}

func (g Grpc) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, g)
}

func (g Grpc) FilePath() string {
	return "src/service/providers"
}

func (g Grpc) FileName() string {
	return "grpc.go"
}

func (g Grpc) TemplateName() string {
	return "grpcFile.tmpl"
}

func (g Grpc) TemplateData(name string) map[string]interface{} {
	return map[string]interface{}{}
}
