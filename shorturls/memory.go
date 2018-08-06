package shorturls

import (
	"encoding/hex"
	"math/rand"
)

// MemoryStore is the implementation of shorturls.Store that stores
// the urls in memory.
type MemoryStore struct {
	URLs map[string]string
}

// Save is MemoryStore's implementation of Store.Save
func (m *MemoryStore) Save(target string) (string, error) {
	b := make([]byte, 4)
	rand.Read(b)
	id := hex.EncodeToString(b)
	m.URLs[id] = target
	return id, nil
}

// Get is MemoryStore's implementation of Store.Get
func (m *MemoryStore) Get(id string) (string, error) {
	return m.URLs[id], nil
}
