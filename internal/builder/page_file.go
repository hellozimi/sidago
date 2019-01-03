package builder

import "path/filepath"

const (
	buildDirectory = "build"
	indexDiretory  = ""
	postsDirectory = "posts"
	pagesDirectory = ""
)

func (p *Page) OutputPath() string {
	return filepath.Join(
		p.sida.basePath,
		buildDirectory,
		p.RelOutputPath(),
	)
}

func (p *Page) RelOutputPath() string {
	return filepath.Join(
		p.directoryForKind(),
		p.PageMeta.Slug+".html",
	)
}

func (p *Page) directoryForKind() string {
	switch p.kind {
	case kindBlog:
		return postsDirectory
	case kindPage:
		return pagesDirectory
	case kindIndex:
		return indexDiretory
	}
	return ""
}
