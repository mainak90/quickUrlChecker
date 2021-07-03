package main

import (
	"flag"
	checker2 "github.com/mainak90/quickUrlChecker/checker"
)

func main() {
	var (
		filepath = flag.String("filepath", "null", "The filepath of the file to parse")
		url = flag.String("url", "null", "The url of the website to check")
	)
	flag.Parse()
	checker2.CheckUrlConFile(*filepath)
	checker2.CheckUrlCmdLine(*url)
}
