package httputil

import (
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"go.dpb.io/sitemap"
	"go.dpb.io/sitemap/data"
)

var DefaultFetcher = Fetcher{
	Client:         http.DefaultClient,
	Decoder:        sitemap.Decode,
	RecurseSitemap: true,
	// IgnoreClientResponseErrors: true,
	// MaxRecursion: 2,
}

type Fetcher struct {
	Client         *http.Client
	Decoder        func(io.Reader, data.EntryCallback) error
	RecurseSitemap bool
	// IgnoreClientResponseErrors bool
	// MaxRecursion int
}

func (f Fetcher) Fetch(sitemapURL string, callback data.EntryCallback) error {
	return f.fetch(sitemapURL, f.wrapCallback(callback))
}

func (f Fetcher) fetch(sitemapURL string, callback data.EntryCallback) error {
	// TODO add visited arg to avoid accidental duplicates

	res, err := http.Get(sitemapURL)
	if err != nil {
		return errors.Wrap(err, "getting url")
	} else if _e, _a := http.StatusOK, res.StatusCode; _e != _a {
		return fmt.Errorf("expected status code `%v` but got: `%v`", _e, _a)
	}

	defer res.Body.Close()

	err = f.Decoder(res.Body, callback)
	if err != nil {
		return errors.Wrap(err, "decoding sitemap")
	}

	return nil
}

func (f Fetcher) wrapCallback(callback data.EntryCallback) data.EntryCallback {
	var builtCallback data.EntryCallback

	if f.RecurseSitemap {
		builtCallback = data.EntryCallbackFunc(func(entry data.Entry) error {
			err := callback.WithEntry(entry)
			if err != nil {
				return err
			}

			if entrySitemap, ok := entry.(*data.Sitemap); ok {
				return f.fetch(entrySitemap.Location, builtCallback)
			}

			return nil
		})
	}

	return builtCallback
}
