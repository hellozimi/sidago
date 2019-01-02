package builder

import "path/filepath"

const (
	buildDirectory = "build"
	indexDiretory  = ""
	postsDirectory = "posts"
	pagesDirectory = "pages"
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
	case KindBlog:
		return postsDirectory
	case KindPage:
		return pagesDirectory
	case KindIndex:
		return indexDiretory
	}
	return ""
}
