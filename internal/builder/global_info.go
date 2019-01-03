package builder

import (
	"html/template"

	"github.com/hellozimi/sidago/internal/builder/config"
)

// GlobalInfo struct for site rendering
type GlobalInfo struct {
	Title       string
	Description string
	Copyright   string
	Data        *map[string]interface{}
	baseURL     string
	sida        *Sida
}

// Config returns the shared config
func (g *GlobalInfo) Config() config.Config {
	return g.sida.config
}

// BaseURL returns the config base url
func (g *GlobalInfo) BaseURL() template.URL {
	return template.URL(g.baseURL)
}

// Posts proxy
func (g *GlobalInfo) Posts() []*Page {
	return g.sida.Posts()
}
