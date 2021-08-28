package examplesutil

import (
	"embed"
	"io/fs"
	"os"
)

//go:embed file.fallback.xml
var embedFS embed.FS

func FileFromArgOrFallback(args []string) (fs.File, error) {
	var (
		fh  fs.File
		err error
	)

	if len(os.Args) > 1 {
		fh, err = os.Open(os.Args[1])
	} else {
		fh, err = embedFS.Open("file.fallback.xml")
	}
	if err != nil {
		return nil, err
	}

	return fh, nil
}
