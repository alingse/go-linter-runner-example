// Package revive a demo code copy from https://github.com/mgechev/revive/issues/1066
package revive

import "maps"

type Config struct {
	b bool         // required
	i int          // optional
	m map[int]bool // optional
}

func NewConfig(b bool) Config {
	return Config{b: b, i: 42}
}

func (c Config) WithI(i int) Config {
	c.i = i
	return c
}

func (c Config) WithM(m map[int]bool) Config {
	c.m = maps.Clone(m)
	return c
}

type Thing struct {
	cfg Config
}

func NewThing(cfg Config) *Thing {
	return &Thing{cfg: cfg}
}
