package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const testHTML = `
<html>
<title>Something</title>
  <body>
	<div class="container-div">
		<h1>h1 tag</h1>
		<h2>h2 tag</h2>
    </div>
  </body>
</html>
`

func TestFindHTMLTitle(t *testing.T) {
	node, err := html.Parse(strings.NewReader(testHTML))
	assert.NoError(t, err, "Error occured while reading test html data")
	title, ok := getHTMLNode(node, getElementByTag(atom.Title))
	assert.True(t, ok, "error while fetching the title tag")
	assert.Equal(t, title, "Something", "expected title text doesnt match")
}

func TestFindAllHeadingTags(t *testing.T) {
	node, err := html.Parse(strings.NewReader(testHTML))
	assert.NoError(t, err, "Error occured while reading test html data")
	h1Tags := getAllNodes(node, getElementByTag(atom.H3))
	assert.True(t, len(h1Tags) == 1, "error while fetching the h1 tag")
}

func Expect(exp bool) {
	if !exp {
		panic(exp)
	}
}

func ExpectNoErr(err error) {
	if err != nil {
		panic(err)
	}
}
