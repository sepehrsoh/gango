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
	makeMainDirectory(name)
	createGoMod(name)
	createMakefile(name)
}

func makeMainDirectory(name string) {
	err := os.Mkdir(name, 0755)
	if err != nil {
		logger.Panic(err)
	}
}

func createGoMod(name string) {
	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = name
	err := cmd.Run()
	if err != nil {
		logger.Panic(err)
	}
}

func createMakefile(name string) {
	splitName := strings.Split(name, "/")
	makeF := strings.ReplaceAll(makefile, "gango", splitName[len(splitName)-1])
	err := utils.WriteFile(name, "Makefile", makeF)
	if err != nil {
		logger.Panic(err)
	}

}

var makefile = `
BINARY_NAME=gango
BINARY_UNIX=$(BINARY_NAME)
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOLIST=$(GOCMD) list
GOTEST=$(GOCMD) test
GOMOD=GO111MODULE=on $(GOCMD) mod
GOGET=GO111MODULE=on $(GOCMD) get

.PHONY: all
all: clean dependencies build

.PHONY: linux
linux: clean dependencies build-linux

.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@rm -f $(BINARY_UNIX)
	$(GOMOD) tidy
	$(GOMOD) vendor
	$(GOCLEAN) -i
	@echo  "Done cleaning."

.PHONY: dependencies
dependencies:
	$(GOMOD) tidy
	$(GOMOD) download
	$(GOMOD) vendor

.PHONY: build
build:
	$(GOBUILD) ./...
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/*.go

.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v  cmd/*.go

.PHONY: build-release
build-release:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags "-s -w" -o $(BINARY_UNIX) -v cmd/*.go

.PHONY: go-lint
go-lint:
	@echo  "running go-lint ..."
	golangci-lint run -v

.PHONY: build-docker
build-docker:
	make linux
	docker build . -t gango:latest
`
