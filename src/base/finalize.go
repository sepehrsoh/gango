package base

import (
	"gango/utils"
	"os/exec"
	"strings"
)

type FinalizeProject struct {
}

func (b *FinalizeProject) Run(name string) {
	downloadPackages(name)
	createDockerfile(name)
}

func downloadPackages(name string) {
	cmd := exec.Command("make", "dependencies")
	cmd.Dir = name
	err := cmd.Run()
	if err != nil {
		logger.Panic(err)
	}
}

func createDockerfile(name string) {
	err := utils.WriteFile(name, "Dockerfile", strings.ReplaceAll(dockerFile, "gango", name))
	if err != nil {
		logger.Panic(err)
	}
}

var dockerFile = `
FROM golang:1.20-alpine as base

WORKDIR /build
COPY . /build/.

CMD ["make","all"]

FROM base 

WORKDIR /app
COPY gango gango

CMD ["./gango","serve"]
`
