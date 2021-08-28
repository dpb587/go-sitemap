// This example shows how EntryAggregate can be used to read all entries into
// memory so they can be iterated on after (vs the traditional manner of
// handling each entry individually as they are parsed from a stream).
//
// Provide a local file to parse or an example sitemap will be used.
package main

import (
	_ "embed"
	"fmt"
	"os"

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

	aggregate := data.NewEntryAggregate()

	err = sitemap.Decode(fh, &aggregate)
	if err != nil {
		panic(err)
	}

	fmt.Printf("# found %d entries\n", len(aggregate))

	for _, entry := range aggregate {
		if entryURL, ok := entry.(*data.URL); ok {
			fmt.Printf("%s \n", entryURL.Location)
		}
	}
}
