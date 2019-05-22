package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	var src, dst string
	var dry bool
	flag.StringVar(&src, "src", "", "source")
	flag.StringVar(&dst, "dst", "", "destination")
	flag.BoolVar(&dry, "dry", false, "dry run")
	flag.Parse()

	if src == "" {
		src = "."
	}
	if dst == "" {
		dst = "."
	}
	if !strings.HasSuffix(src, "/") {
		src = src + "/"
	}
	if !strings.HasSuffix(dst, "/") {
		src = dst + "/"
	}

	if _, err := os.Stat(src); os.IsNotExist(err) {
		log.Fatal("source directory does not exist")
	}

	if _, err := os.Stat(dst); os.IsNotExist(err) {
		if err := os.Mkdir(dst, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	files, err := ioutil.ReadDir(src)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filename := file.Name()
		in := fmt.Sprintf("%s%s", src, filename)
		data, err := ioutil.ReadFile(in)
		if err != nil {
			log.Fatal(err)
		}

		re := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2}).*`)
		date := re.ReplaceAllString(filename, "$1-$2-$3")
		date = fmt.Sprintf("%sT00:00:00-00:00", date)

		re = regexp.MustCompile(`(\d{4}-\d{2}-\d{2})-(.*)`)
		newFilename := re.ReplaceAllString(filename, "$2")

		re = regexp.MustCompile(`---\n([\s\S]*)---\n([\s\S]*)`)
		matches := re.FindSubmatch(data)
		header := string(matches[1])
		body := string(matches[2])
		dateline := fmt.Sprintf("date: %s\n", date)
		draftline := "draft: false\n"
		header = fmt.Sprintf("%s%s%s", header, dateline, draftline)

		content := fmt.Sprintf("---\n%s---\n%s", header, body)

		out := dst + newFilename
		fmt.Printf("%s -> %s", in, out)

		if !dry {
			if err := ioutil.WriteFile(out, []byte(content), 0644); err != nil {
				log.Fatal(err)
			}
		}
	}
}
