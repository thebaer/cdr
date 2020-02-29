package cdr

type (
	Mixtape struct {
		Tracks []Track
	}

	Track struct {
		Num    int
		Title  string
		Artist string

		Filename string
	}
)
