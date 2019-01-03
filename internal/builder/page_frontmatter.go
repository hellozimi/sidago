package builder

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

func (p *Page) updateWithFrontmatter(f map[string]interface{}) {
	if x, ok := f["title"]; ok {
		if val, ok := x.(string); ok && len(val) > 0 {
			p.Title = val
		}
	}
	if x, ok := f["date"]; ok {
		if val, ok := x.(string); ok && len(val) > 0 {
			if !strings.HasSuffix(val, "Z") {
				val = val + "Z"
			}
			date, err := time.Parse(time.RFC3339, val)
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
