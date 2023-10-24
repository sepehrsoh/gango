package errs

import (
	"path/filepath"

	"gango/utils"
)

type Errors struct{}

func (s Errors) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(s.FilePath(), s.FileName()), signals)
}

func (s Errors) FileName() string {
	return "errs.go"
}

func (s Errors) FilePath() string {
	return "/src/lib/errs"
}

var signals = `// Code generate by Gogang
package errs

var (
	ErrInternalServer = "internal server error"
)
`
