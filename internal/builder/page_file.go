package builder

import "path/filepath"

const (
	buildDirectory = "build"
	postsDirectory = "posts"
	pagesDirectory = "pages"
)

func (p *Page) OutputPath() string {
	return filepath.Join(
		p.sida.basePath,
		buildDirectory,
		p.directoryForKind(),
		p.PageMeta.Slug+".html")
}

func (p *Page) directoryForKind() string {
	switch p.kind {
	case KindBlog:
		return postsDirectory
	case KindSingle:
		return pagesDirectory
	default:
		panic("Unknown post kind")
	}
}
