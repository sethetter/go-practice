package shorturls

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomePageHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "https://example.com/", nil)
	w := httptest.NewRecorder()

	HomePageHandler(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Contains(t, string(body), "ShortURLs")
}

func TestHomePageShowsCreatedURLInfo(t *testing.T) {
	req := httptest.NewRequest("GET", "https://example.com/?id=abc123", nil)
	w := httptest.NewRecorder()

	HomePageHandler(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Contains(t, string(body), "abc123")
}

func TestShortenUrlHandler(t *testing.T) {
	body := &url.Values{}
	body.Set("url", "https://google.com")
	req := httptest.NewRequest("POST", "https://example.com/", strings.NewReader(body.Encode()))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()

	store := &MemoryStore{make(map[string]string)}
	shortener := NewShortener(store)

	ShortenURLHandler(shortener, w, req)
	resp := w.Result()

	assert.Equal(t, 302, resp.StatusCode)

	foundID := grabIDFromStore(store)

	locationHeader := resp.Header.Get("location")
	expectedLocation := fmt.Sprintf("https://example.com/?id=%s", foundID)
	assert.Contains(t, locationHeader, expectedLocation)
}

func TestShortenUrlHandlerEmptyURL(t *testing.T) {
	body := url.Values{"url": {""}}
	req := httptest.NewRequest("POST", "https://example.com/", strings.NewReader(body.Encode()))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()

	store := &MemoryStore{make(map[string]string)}
	shortener := NewShortener(store)

	ShortenURLHandler(shortener, w, req)
	resp := w.Result()

	assert.Equal(t, 400, resp.StatusCode)
}

func TestLookupUrlHandler(t *testing.T) {
	store := &MemoryStore{make(map[string]string)}
	shortener := NewShortener(store)

	url := "https://google.com"
	id, err := shortener.Shorten(url)
	assert.NoError(t, err)

	req := httptest.NewRequest("GET", fmt.Sprintf("https://example.com/%s", id), nil)
	w := httptest.NewRecorder()

	LookupURLHandler(shortener, w, req)
	resp := w.Result()

	assert.Equal(t, 302, resp.StatusCode)

	locationHeader := resp.Header.Get("location")
	assert.Equal(t, url, locationHeader)
}

func grabIDFromStore(store *MemoryStore) string {
	var foundID string
	for id := range store.URLs {
		foundID = id
	}
	return foundID
}
