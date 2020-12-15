package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestWithMultipleUrls(t *testing.T) {

	urls := []string{
		"https://stackoverflow.com/users/login?ssrc=head&returnurl=https%3a%2f%2fstackoverflow.com%2fquestions%2f2818852%2fis-there-a-queue-implementation",
		"http://jkorpela.fi/HTML3.2/",
		"https://github.com",
		"https://id.heroku.com/login",
	}

	for _, urlString := range urls {
		err := testServeHTTPScripts(urlString)
		if err != nil {
			t.Error(err)
		}
	}

}

func testServeHTTPScripts(urlString string) error {
	handler := indexHandler{}
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		return err
	}
	req.Form = url.Values{}
	req.Form.Add("urldetails", urlString)

	rr := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(handler.ServeHTTP)
	handlerFunc.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		return err
	}

	return nil
}
