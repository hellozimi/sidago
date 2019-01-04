package content

import (
	"fmt"
	"time"
)

var (
	errNoContent = fmt.Errorf("no content for content type")
)

type boilerplate struct {
	title string
	date  time.Time
	t     contentType
}

func (b *boilerplate) content() (string, error) {
	switch b.t {
	case typePage:
		return pageContent(b.title), nil
	case typePost:
		return postContent(b.title, b.date), nil
	default:
		return "", errNoContent
	}
}

func newBoilerplate(title string, date time.Time, t contentType) boilerplate {
	return boilerplate{title, date, t}
}

func postContent(title string, date time.Time) string {
	return fmt.Sprintf(`---
title = "%s"
date = "%s"
---
# %s`, title, date.Format(time.RFC3339), title)
}

func pageContent(title string) string {
	return fmt.Sprintf("# %s", title)
}
