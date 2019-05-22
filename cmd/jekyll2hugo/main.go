package main

import (
	"flag"

	"github.com/miguelmota/go-jekyll-to-hugo/migration"
)

func main() {
	var src, dst string
	var dry bool
	flag.StringVar(&src, "src", "", "source")
	flag.StringVar(&dst, "dst", "", "destination")
	flag.BoolVar(&dry, "dry", false, "dry run")
	flag.Parse()

	migration.Migrate(&migration.Options{
		Source:      src,
		Destination: dst,
		Dry:         dry,
	})
}
