package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Usage prints usage.
func Usage() {
	fmt.Println("usage: grepfiles <tag> <attr (optional)> [-html]")
}

func main() {
	htmlPtr := flag.Bool("html", false, "output as html")
	flag.Parse()
	cmdArgs := flag.Args()

	tag := cmdArgs[0]
	var attr string
	if len(cmdArgs) >= 2 {
		attr = cmdArgs[1]
	}

	if tag == "default" {
		fmt.Println("[!] tag name must be specified.")
		os.Exit(0)
	}

	scanner := bufio.NewScanner(os.Stdin)
	input := ""
	for scanner.Scan() {
		line := scanner.Text()
		input += line + "\n"
	}

	//fmt.Println(input)
	r := strings.NewReader(input)
	doc, _ := goquery.NewDocumentFromReader(r)

	if attr == "" {
		if *htmlPtr {
			doc.Find(tag).Each(func(i int, el *goquery.Selection) {
				fmt.Println(el.Html())
			})
		} else {
			doc.Find(tag).Each(func(i int, el *goquery.Selection) {
				output := el.Text()
				output = strings.Replace(output, "\n", "", -1)
				output = strings.TrimSpace(output)
				if output != "" {
					fmt.Println(output)
				}
			})
		}
	} else {
		doc.Find(tag).Each(func(i int, el *goquery.Selection) {
			attr_val, _ := el.Attr(attr)
			if attr_val != "" {
				fmt.Println(attr_val)
			}
		})
	}
}
