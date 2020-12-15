package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

//PageDetails Gives the page details.
type PageDetails struct {
	ErrorString       string
	HTMLVersion       string
	PageTitle         string
	InputURL          string
	InternalLinks     []string
	ExternalLinks     []string
	InAccessibleLinks []string
	HTags             map[string]int
	HasLoginForm      bool
}

type indexHandler struct {
	url string
}

func newIndexHandler() *indexHandler {
	return &indexHandler{}
}

func (me *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	me.handler(w, r)
}

func (me *indexHandler) handler(w http.ResponseWriter, r *http.Request) {

	data := PageDetails{}
	r.ParseForm()
	url := r.Form["urldetails"]

	if len(url) == 0 {
		data.ErrorString = "Please enter the URL to get the details"
		tpl.Execute(w, data)
		return
	}

	urlValue := url[0]

	data.InputURL = urlValue
	resp, err := http.Get(urlValue)

	if err != nil { // handle error
		data.ErrorString = "Enter valid url"
		fmt.Println("ERROR: Failed to get data:", url)
		tpl.Execute(w, data)
		return
	}

	defer resp.Body.Close() // close Body when the function completes

	//Read all the response
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil { // handle error
		fmt.Println("ERROR: Failed to read response:", url)
		return
	}

	root, err := html.Parse(bytes.NewReader(respBytes))
	if err != nil { // handle error
		fmt.Println("ERROR: Failed to get data:", url)
		return
	}

	//Get title
	if titleHTMLNode, ok := getHTMLNode(root, getElementByTag(atom.Title)); ok {
		data.PageTitle = getText(titleHTMLNode)
	}

	//Get HTags
	data.HTags = me.getAllHeadTagsCount(root)

	//Get internal & external links
	data.InternalLinks, data.ExternalLinks = me.getAllInternalAndExternalLinks(root)

	//Get HTML version
	data.HTMLVersion = getHTMLDoctype(bytes.NewReader(respBytes))

	//Find Login Form
	//Finding this could be dependant on the website & individial developer.
	//Here we assume that most logins are developed with form controls and it have one password field in it.
	data.HasLoginForm = me.hasLoginForm(root)

	//Get All Inaccessible links
	data.InAccessibleLinks = me.getAllInaccessibleLinks(data.ExternalLinks)

	err = tpl.Execute(w, data)
	if err != nil {
		//handle error
	}
}

//getAllHeadTagsCount returns the map of the all headings tags along with its count.
func (me *indexHandler) getAllHeadTagsCount(root *html.Node) map[string]int {

	hTags := make(map[string]int, 0)
	headingNode := getAllNodes(root, getElementByTag(atom.H1))
	hTags["h1"] = len(headingNode)

	headingNode = getAllNodes(root, getElementByTag(atom.H2))
	hTags["h2"] = len(headingNode)

	headingNode = getAllNodes(root, getElementByTag(atom.H3))
	hTags["h3"] = len(headingNode)

	headingNode = getAllNodes(root, getElementByTag(atom.H4))
	hTags["h4"] = len(headingNode)

	headingNode = getAllNodes(root, getElementByTag(atom.H5))
	hTags["h5"] = len(headingNode)

	headingNode = getAllNodes(root, getElementByTag(atom.H6))
	hTags["h6"] = len(headingNode)

	return hTags
}

//getAllInternalAndExternalLinks gets all the internal & external links based on dns look up of a href attribute.
func (me *indexHandler) getAllInternalAndExternalLinks(root *html.Node) ([]string, []string) {

	externalLinks := make([]string, 0)
	internalLinks := make([]string, 0)
	links := getAllNodes(root, getElementByTag(atom.A))
	var wg sync.WaitGroup
	wg.Add(len(links))

	for _, link := range links {
		go func(link *html.Node) {
			defer wg.Done()
			linkText := getAttr(link, "href")
			if isURL(linkText) { //External links
				externalLinks = append(externalLinks, linkText)
			} else { //Internal links
				internalLinks = append(internalLinks, linkText)
			}
		}(link)
	}
	wg.Wait()
	return internalLinks, externalLinks
}

func isURL(str string) bool {

	//Parse the string to url
	url, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}

	//dns look up
	address := net.ParseIP(url.Host)

	if address == nil {
		return strings.Contains(url.Host, ".")
	}
	//if address is not nil, then consider it as external link.
	return true
}

//hasLoginForm the validation for login form is, the webpage should have FORM tag, and one password inside it.
func (me *indexHandler) hasLoginForm(root *html.Node) bool {

	allformTags := getAllNodes(root, getElementByTag(atom.Form))

	for _, formTag := range allformTags {

		allFormInputs := getAllNodes(formTag, getElementByTag(atom.Input))
		passwordFieldCount := 0

		for _, input := range allFormInputs {
			if getAttr(input, "type") == "password" {
				passwordFieldCount++
			}
		}

		if passwordFieldCount == 1 {
			return true
		}
	}

	return false
}

// getAllInaccessibleLinks  checks the status of links to determine whether links is accessible or not.
func (me *indexHandler) getAllInaccessibleLinks(allLinks []string) []string {

	inaccessibleLinks := make([]string, 0)
	var wg sync.WaitGroup
	wg.Add(len(allLinks))

	for _, link := range allLinks {
		go func(link string) {
			defer wg.Done()

			resp, err := http.Head(link)
			if err != nil {
				fmt.Println(err)
				return
			}

			if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
				fmt.Println("HTTP Status is in the 2xx range")
			} else {
				inaccessibleLinks = append(inaccessibleLinks, link)
				fmt.Printf("Broken link: %q", link)
			}
		}(link)
	}

	wg.Wait()

	return inaccessibleLinks
}
