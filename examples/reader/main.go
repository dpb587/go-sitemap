// This example shows the common approach of handling each entry as it is
// parsed and looking at individual fields for further use.
//
// Provide a local file to parse or an example sitemap will be used.
package main

import (
	"fmt"
	"os"
	"strings"

	"go.dpb.io/sitemap"
	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/examples/examplesutil"
)

func main() {
	fh, err := examplesutil.FileFromArgOrFallback(os.Args)
	if err != nil {
		panic(err)
	}

	defer fh.Close()

	err = sitemap.Decode(
		fh,
		data.EntryCallbackFunc(func(entry data.Entry) error {
			if entryURL, ok := entry.(*data.URL); ok {
				var extra []string
				if _v := entryURL.LastModified; _v != nil {
					extra = append(extra, fmt.Sprintf("lastmod=%s", *_v))
				}
				if _v := entryURL.ChangeFrequency; _v != nil {
					extra = append(extra, fmt.Sprintf("changefreq=%s", *_v))
				}
				if _v := entryURL.Priority; _v != nil {
					extra = append(extra, fmt.Sprintf("priority=%s", *_v))
				}

				fmt.Printf("%s [%s]\n", entryURL.Location, strings.Join(extra, "; "))
			}

			return nil
		}),
	)
	if err != nil {
		panic(err)
	}
}
