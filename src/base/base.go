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
	createReadMeFile(name)
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

func createReadMeFile(name string) {
	err := utils.WriteFile(name, "README.md", strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(readmeFile, "~", "`"),
			"gango", name),
		"\"", "```"))

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

.PHONY: build-docker
build-docker:
	make linux
	docker build . -t gango:latest
`

var readmeFile = `
# gango

Welcome to gango! This repository is generated using Gango, a Go-based backend service framework. This README will help you get started with your new project.

## Getting Started

To create this project on your local machine, use the following command:

" bash
make all
"

This command will set up your project environment and generate a new service unix file with the name gango.

You can run your service using the following command:

"bash
./gango serve
"

For example, we've included a sample "Hello, World!" API on the route ~/v1/hello-world~ to get you started.

## Building for Linux

To build your project for Linux, use the following command:

"bash
make linux
"

## Updating Dependencies

To update project dependencies, run:

"bash
make dependencies
"

## Dockerize Your Project

Dockerizing your project is made easy with Gango. You can create a Docker image with the following command:

"bash
make build-docker
docker run gango:latest
"

Once the image is built, you can run it using Docker.

## Project Structure

Inside the ~src/service~ directory, you'll find three important subdirectories:

1. **configs**: To add or modify configurations for your service, update the files in this directory.

2. **providers**: This directory is where you can add more service providers. Examples like Redis and Gin Gonic, which serve as the web server, can be found here.

3. **wiring**: In the ~wire.go~ file, you'll find a struct using the Singleton pattern. To add services to this wiring, you can modify the ~service.go~ file. Internal methods, such as metric host and port, controller registry, and other critical components, are located in ~internal.go~.

## Contributions

We welcome contributions from the community to enhance and improve gango. If you're interested in contributing, please follow our [Contribution Guidelines](CONTRIBUTING.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

Thank you for choosing Gango for your backend development needs. We hope this project simplifies the process of building robust and scalable backend services.

`
