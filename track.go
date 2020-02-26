package cdr

type (
	Mixtape struct {
		Tracks []Track
	}

	Track struct {
		Title  string
		Artist string

		Filename string
	}
)
