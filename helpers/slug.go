package helpers

// IsSlug checks if input is a slug
func IsSlug(in string) bool {
	return Slugify(in) == in
}
