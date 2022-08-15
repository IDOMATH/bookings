package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"generals", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"majors", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post-search-availability", "/search-availability", "POST", []postData{
		{key: "start", value: "2023-01-01"},
		{key: "end", value: "2023-01-02"},
	}, http.StatusOK},
	{"post-search-availability-json", "/search-availability", "POST", []postData{
		{key: "start", value: "2023-01-01"},
		{key: "end", value: "2023-01-02"},
	}, http.StatusOK},
	{"post-make-reservation", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Blane"},
		{key: "last_name", value: "Tschida"},
		{key: "email", value: "me@here.com"},
		{key: "phone", value: "555-555-5555"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, test := range theTests {
		if test.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + test.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s, expected %d, but got %d", test.name, test.expectedStatusCode, resp.StatusCode)
			}
		} else {
			// Code for POST tests
			values := url.Values{}
			for _, param := range test.params {
				values.Add(param.key, param.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+test.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s, expected %d, but got %d", test.name, test.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
