package cmd

import (
	"gango/utils"
	"path/filepath"
	"strings"
)

type ProjectName struct {
}

func (m ProjectName) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(m.FilePath(), m.FileName()), strings.ReplaceAll(mainFile, "gango", dir))
}

func (p ProjectName) FilePath() string {
	return "/cmd"
}

func (p ProjectName) FileName() string {
	return "base.go"
}

var mainFile = `
package main

import "gango/cmd/base"

func main() {
	base.Execute()
}
`
