package data

import (
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

func NewEntryAggregate() EntryAggregate {
	return EntryAggregate{}
}

// EntryAggregate is an EntryList which can be used as an EntryCallback with the behavior of appending entries to it.
type EntryAggregate EntryList

var _ EntryCallback = &EntryAggregate{}

func (va *EntryAggregate) WithEntry(entry Entry) error {
	*va = append(*va, entry)

	return nil
}

// RequireValidLocation drops entries which do not have a location that is within the given base URL. If there is an
// error parsing a URL, then it will be dropped. Entries other than Sitemap and URL remain unfiltered.
func RequireValidLocation(baseURL string, callback EntryCallback) EntryCallback {
	baseParsedURL, err := url.Parse(baseURL)
	if err != nil {
		return EntryCallbackFunc(func(entry Entry) error {
			return errors.Wrap(err, "parsing base url")
		})
	}

	return EntryCallbackFunc(func(entry Entry) error {
		var loc string

		switch entryT := entry.(type) {
		case *Sitemap:
			loc = entryT.Location
		case *URL:
			loc = entryT.Location
		default:
			return callback.WithEntry(entry)
		}

		locParsedURL, err := url.Parse(loc)
		if err != nil {
			return nil
		} else if baseParsedURL.Scheme != locParsedURL.Scheme {
			return nil
		} else if baseParsedURL.Host != locParsedURL.Host {
			return nil
		} else if !strings.HasPrefix(locParsedURL.Path, baseParsedURL.Path) {
			return nil
		}

		return callback.WithEntry(entry)
	})
}

// ResolveRelativeLocations will resolve entries which have a relative URL for their location (although, according to
// specifications these must already be absolute). If there is an error parsing a URL, then it will be dropped. Entries
// other than Sitemap and URL remain unfiltered.
func ResolveRelativeLocations(baseURL string, callback EntryCallback) EntryCallback {
	baseParsedURL, err := url.Parse(baseURL)
	if err != nil {
		return EntryCallbackFunc(func(entry Entry) error {
			return errors.Wrap(err, "parsing base url")
		})
	}

	return EntryCallbackFunc(func(entry Entry) error {
		switch entryT := entry.(type) {
		case *Sitemap:
			locURL, err := url.Parse(entryT.Location)
			if err != nil {
				return nil
			}

			entryT.Location = baseParsedURL.ResolveReference(locURL).String()
		case *URL:
			locURL, err := url.Parse(entryT.Location)
			if err != nil {
				return nil
			}

			entryT.Location = baseParsedURL.ResolveReference(locURL).String()
		}

		return callback.WithEntry(entry)
	})
}
