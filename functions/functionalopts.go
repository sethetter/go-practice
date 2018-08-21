package main

import "fmt"

type Layout int

const (
	landscape Layout = 0
	portrait  Layout = 1
)

type Size struct {
	x, y int
}

type Config struct {
	layout Layout
	name   string
	size   Size
}

type Terrain struct {
	config Config
}

func (t *Terrain) Layout() string {
	switch t.config.layout {
	case landscape:
		return "landscape"
	case portrait:
		return "portrait"
	default:
		return "unknown"
	}
}

func NewTerrain(options ...func(*Config)) *Terrain {
	var t Terrain
	for _, opt := range options {
		opt(&t.config)
	}
	return &t
}

func withPortrait(c *Config) {
	c.layout = portrait
}

func main() {
	t := NewTerrain(withPortrait)
	fmt.Printf("Terrain layout: %s\n", t.Layout())
}
