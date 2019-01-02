package builder

import (
	"html/template"

	"github.com/hellozimi/sidago/internal/builder/config"
)

type GlobalInfo struct {
	Title     string
	baseURL   string
	Copyright string
	Data      *map[string]interface{}
	sida      *Sida
}

func (g *GlobalInfo) Config() config.Config {
	return g.sida.config
}

func (g *GlobalInfo) BaseURL() template.URL {
	return template.URL(g.baseURL)
}

func (g *GlobalInfo) Posts() []*Page {
	return g.sida.Posts()
}
