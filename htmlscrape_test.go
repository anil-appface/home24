package main

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const testHTMLStr = `
<!Doctype html>
<html>
<title>Something</title>
  <body>
	<div class="container-div">
		<h1>h1 tag</h1>
		<h2>h2 tag</h2>
		<h3>h3 tag 1</h3>
		<h3>h3 tag 2</h3>
		<h4>h4 tag</h4>
		<h5>h5 tag</h5>
		<h6>h6 tag</h6>
		<a href="https://google.com">external link</a>
		<a href="#internalLink">internal link</a>
    </div>
  </body>
</html>
`

func TestFindHTMLTitle(t *testing.T) {
	node, err := html.Parse(strings.NewReader(testHTMLStr))
	ExpectNoErr(err, "Error occured while reading test html data")
	title, ok := getHTMLNode(node, getElementByTag(atom.Title))
	Expect(ok, "error while fetching the title tag")
	Expect(getText(title) == "Something", "expected title text doesnt match, expected:%s, found: %s", "Something", getText(title))
}

func TestFindAllHeadingTags(t *testing.T) {
	node, err := html.Parse(strings.NewReader(testHTMLStr))
	ExpectNoErr(err, "Error occured while reading test html data")
	h1Tags := getAllNodes(node, getElementByTag(atom.H1))
	Expect(len(h1Tags) == 1, "error while fetching the h1 tag")
	h2Tags := getAllNodes(node, getElementByTag(atom.H2))
	Expect(len(h2Tags) == 1, "error while fetching the h1 tag")
	h3Tags := getAllNodes(node, getElementByTag(atom.H3))
	Expect(len(h3Tags) == 2, "error while fetching the h1 tag")
	h4Tags := getAllNodes(node, getElementByTag(atom.H4))
	Expect(len(h4Tags) == 1, "error while fetching the h1 tag")
	h5Tags := getAllNodes(node, getElementByTag(atom.H5))
	Expect(len(h5Tags) == 1, "error while fetching the h1 tag")
	h6Tags := getAllNodes(node, getElementByTag(atom.H6))
	Expect(len(h6Tags) == 1, "error while fetching the h1 tag")
}

func TestFindHTMLVersion(t *testing.T) {
	version := getHTMLDoctype(strings.NewReader(testHTMLStr))
	Expect(version == "HTML 5.0", "expected version text doesnt match, expected:%s, found: %s", "HTML 5.0", version)
}

func TestLinks(t *testing.T) {
	node, err := html.Parse(strings.NewReader(testHTMLStr))
	ExpectNoErr(err, "Error occured while reading test html data")
	links := getAllNodes(node, getElementByTag(atom.A))
	Expect(len(links) == 2, "expected length of links doesnt match , expected: %d, found: %d", 2, len(links))
}

//utility methods to check the errors on test cases
func Expect(exp bool, errorMessage string, args ...interface{}) {
	if !exp {
		fmt.Printf(errorMessage, args...)
		panic(exp)
	}
}

func ExpectNoErr(err error, errorMessage string, args ...interface{}) {
	if err != nil {
		fmt.Printf(errorMessage, args...)
		panic(err)
	}
}
