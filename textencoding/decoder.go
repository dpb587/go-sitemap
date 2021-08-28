package textencoding

import (
	"bufio"
	"io"
	"strings"

	"go.dpb.io/sitemap/data"
)

var entryTypeURL = data.EntryType{
	Space: "text",
	Local: "URL",
}

type Decoder struct {
	reader io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		reader: r,
	}
}

func (d *Decoder) Decode(callback data.EntryCallback) error {
	scanner := bufio.NewScanner(d.reader)

	for scanner.Scan() {
		var offset int64 = 0 // TODO

		text := strings.TrimSpace(scanner.Text())
		if len(text) == 0 {
			continue
		}

		e := data.NewURL(entryTypeURL, offset, nil)
		e.Location = text

		err := callback.WithEntry(e)
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
