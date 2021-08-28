package textencoding

import (
	"io"

	"go.dpb.io/sitemap/data"
)

// Decode processes the input of a plain text files with one URL per line.
func Decode(r io.Reader, callback data.EntryCallback) error {
	return NewDecoder(r).Decode(callback)
}
