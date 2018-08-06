// Package shorturls is the core of the url shortening service.
package shorturls

// Store is the interface that storage adapters must adhere to
type Store interface {
	Save(string) (string, error)
	Get(string) (string, error)
}

// Shortener defines a service that creates shortened URL resources.
type Shortener struct {
	store Store
}

// NewShortener returns a new Shortener.
func NewShortener(store Store) *Shortener {
	return &Shortener{store}
}

// Shorten takes a target URL and creates then returns the lookup id.
func (s *Shortener) Shorten(target string) (string, error) {
	id, err := s.store.Save(target)
	if err != nil {
		return "", err
	}
	return id, nil
}

// Lookup takes an id and returns the stored url for it, if available.
func (s *Shortener) Lookup(id string) (string, error) {
	target, err := s.store.Get(id)
	if err != nil {
		return "", err
	}
	return target, nil
}
