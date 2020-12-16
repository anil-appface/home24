package main

import (
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Matcher should return true when a desired node is found.
type Matcher func(node *html.Node) bool

// getAllNodes returns all nodes which match the provided Matcher.
func getAllNodes(node *html.Node, matcher Matcher) []*html.Node {
	return findAllInternal(node, matcher, false)
}

// getHTMLNode returns the first node which matches the matcher using depth-first search.
func getHTMLNode(node *html.Node, matcher Matcher) (n *html.Node, ok bool) {
	if matcher(node) {
		return node, true
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		n, ok := getHTMLNode(c, matcher)
		if ok {
			return n, true
		}
	}
	return nil, false
}

// getText gets the text of html node.
func getText(node *html.Node) string {
	joiner := func(s []string) string {
		n := 0
		for i := range s {
			trimmed := strings.TrimSpace(s[i])
			if trimmed != "" {
				s[n] = trimmed
				n++
			}
		}
		return strings.Join(s[:n], " ")
	}
	return textJoin(node, joiner)
}

func textJoin(node *html.Node, join func([]string) string) string {
	nodes := getAllNodes(node, func(n *html.Node) bool { return n.Type == html.TextNode })
	parts := make([]string, len(nodes))
	for i, n := range nodes {
		parts[i] = n.Data
	}
	return join(parts)
}

// Attr returns the value of an HTML attribute.
func getAttr(node *html.Node, key string) string {
	for _, a := range node.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}

// getElementByTag returns a Matcher which matches all nodes of the provided tag type.
func getElementByTag(a atom.Atom) Matcher {
	return func(node *html.Node) bool { return node.DataAtom == a }
}

// findAllInternal encapsulates the node tree traversal
func findAllInternal(node *html.Node, matcher Matcher, searchNested bool) []*html.Node {
	matched := []*html.Node{}

	if matcher(node) {
		matched = append(matched, node)

		if !searchNested {
			return matched
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		found := findAllInternal(c, matcher, searchNested)
		if len(found) > 0 {
			matched = append(matched, found...)
		}
	}
	return matched
}

//getNextSibling Find returns the first node which matches the matcher using next sibling search.
func getNextSibling(node *html.Node, matcher Matcher) (n *html.Node, ok bool) {

	for s := node.NextSibling; s != nil; s = s.NextSibling {
		if matcher(s) {
			return s, true
		}
	}
	return nil, false
}

//getHTMLDoctype gives the doctype tag in HTML
func getHTMLDoctype(r io.Reader) string {

	tokenizer := html.NewTokenizer(r)
	for {
		//get the next token type
		tokenType := tokenizer.Next()

		//if it's an error token, we either reached
		//the end of the file, or the HTML was malformed
		if tokenType == html.DoctypeToken {
			doctype := tokenizer.Token().Data
			startIndexStr := "-//W3C//DTD "
			startIndex := strings.Index(doctype, startIndexStr)
			endIndex := strings.Index(doctype, "//EN")
			if strings.ToLower(doctype) == "html" {
				return "HTML 5.0"
			}
			htmlVersion := doctype[startIndex+len(startIndexStr) : endIndex]
			return htmlVersion
			//To check htmlversion, I have created a small playgroud@ https://play.golang.org/p/EdSvwXzQPfA
		}
	}
}
