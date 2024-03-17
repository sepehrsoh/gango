package utils

func GetDefaultTemplateValues(name string) map[string]interface{} {
	return map[string]interface{}{
		"ProjectName": name,
	}
}
