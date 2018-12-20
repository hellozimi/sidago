package main

// Error proxy
type Error string

func (e Error) Error() string {
	return string(e)
}
