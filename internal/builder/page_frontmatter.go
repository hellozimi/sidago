package builder

import (
	"fmt"
	"path/filepath"

	"github.com/hellozimi/sidago/helpers"
)

func (p *Page) updateWithFrontmatter(f map[string]interface{}) {
	if x, ok := f["title"]; ok {
		if val, ok := x.(string); ok && len(val) > 0 {
			p.Title = val
		}
	}
	if x, ok := f["date"]; ok {
		if val, ok := x.(string); ok && len(val) > 0 {
			date, err := helpers.ParseDate(val)
			if err != nil {
				fname := filepath.Base(p.path)
				fmt.Printf("error parsing frontmatter date for %s\n\n", fname)
			} else {
				p.Date = date
			}
		}
	}

	if x, ok := f["draft"]; ok {
		if val, ok := x.(bool); ok {
			p.Draft = val
		}
	}
}
