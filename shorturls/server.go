package shorturls

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
)

// IndexData is the data that will be passed to the home page template.
type IndexData struct {
	Title     string
	HasNewURL bool
	NewURL    string
}

// RootHandler is the handler that hooks up routing
func RootHandler(shortener *Shortener) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		if r.URL.Path == "/" {
			switch r.Method {
			case http.MethodGet:
				HomePageHandler(w, r)
			case http.MethodPost:
				ShortenURLHandler(shortener, w, r)
			default:
				w.WriteHeader(404)
				return
			}
		} else if lookupRouteMatch, _ := regexp.Match(
			"/[0-9a-zA-Z]+",
			[]byte(r.URL.Path),
		); lookupRouteMatch {
			LookupURLHandler(shortener, w, r)
		} else {
			w.WriteHeader(404)
			return
		}
	}
}

// HomePageHandler handles the request for the homepage
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: proper error handling here? What happens if I don't?
	w.Header().Add("content-type", "text/html")
	t, _ := template.ParseFiles("./tmpl/index.html")

	id := r.URL.Query().Get("id")

	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}

	data := IndexData{
		Title:     "ShortURLs",
		HasNewURL: id != "",
		NewURL:    scheme + r.Host + "/" + id,
	}

	t.Execute(w, data)
}

// ShortenURLHandler accepts a URL and returns the shortened version
func ShortenURLHandler(shortener *Shortener, w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(500)
		return
	}

	url := r.Form.Get("url")
	if url == "" {
		w.WriteHeader(400)
		return
	}

	id, err := shortener.Shorten(url)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Add("location", fmt.Sprintf("%s?id=%s", r.URL.String(), id))
	w.WriteHeader(302)
	return
}

// LookupURLHandler accepts a URL id and returns the URL value for it
func LookupURLHandler(shortener *Shortener, w http.ResponseWriter, r *http.Request) {
	url, err := shortener.Lookup(r.URL.Path[1:])

	// TODO: Proper handling of NotFound
	if err != nil || url == "" {
		w.WriteHeader(500)
		return
	}

	w.Header().Add("location", url)
	w.WriteHeader(302)
	return
}
