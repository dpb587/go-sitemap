// This example uses the built-in fetcher which downloads a URL and, by default,
// automatically recurses any child sitemaps it finds. The local callback then
// only needs to care about URLs which are found in any sitemap.
//
// Provide a remote sitemap URL to parse or an example sitemap will be used.
package main

import (
	"fmt"
	"os"

	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/examples/examplesutil"
	"go.dpb.io/sitemap/httputil"
)

func main() {
	sitemapURL := examplesutil.URLFromArgOrDefault(os.Args, "https://wordpress.com/sitemap.xml")

	err := httputil.DefaultFetcher.Fetch(
		sitemapURL,
		data.EntryCallbackFunc(func(entry data.Entry) error {
			switch eT := entry.(type) {
			case *data.Sitemap:
				fmt.Printf("# %s\n", eT.Location)
			case *data.URL:
				fmt.Printf("> %s\n", eT.Location)
			}

			return nil
		}),
	)
	if err != nil {
		panic(err)
	}
}
