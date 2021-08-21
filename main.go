package main

import (
	"bufio"
	"os"
	"fmt"
	"flag"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	tag_ptr := flag.String("tag", "default", "tag name")
	attr_ptr := flag.String("attr", "default", "attribute name")
	html_ptr := flag.String("html", "false", "output as html")
	flag.Parse()

	tag := *tag_ptr
	attr := *attr_ptr
	html := *html_ptr

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

	if attr == "default" {
		if html == "true" {
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
