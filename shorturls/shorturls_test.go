package shorturls

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShorten(t *testing.T) {
	store := &MemoryStore{make(map[string]string)}
	shortener := NewShortener(store)

	id, err := shortener.Shorten("https://google.com")

	if assert.NoError(t, err) {
		assert.NotEmpty(t, id)
	}
}

func TestLookup(t *testing.T) {
	store := &MemoryStore{make(map[string]string)}
	shortener := NewShortener(store)

	target := "https://google.com"
	id, err := shortener.Shorten(target)

	if assert.NoError(t, err) {
		got, err := shortener.Lookup(id)
		assert.NoError(t, err)
		assert.Equal(t, target, got)
	}
}
