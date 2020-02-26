package cdr

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/dhowden/tag"
)

var trackNameReg = regexp.MustCompile("^([0-9]{2}).+")

func NewTrack(file string) *Track {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("error loading file: %v", err)
		return nil
	}
	defer f.Close()

	m, err := tag.ReadFrom(f)
	if err != nil {
		return nil
	}

	return &Track{
		Title:    m.Title(),
		Artist:   m.Artist(),
		Filename: f.Name(),
	}
}

// RenameTrack takes a filename, opens it, reads the metadata, and returns both
// the old and new filename.
func RenameTrack(file string) string {
	t := NewTrack(file)

	// Extract playlist track number from filename
	fMatch := trackNameReg.FindStringSubmatch(t.Filename)
	if len(fMatch) < 2 {
		log.Fatal("Unexpect filename format")
	}
	trackNum := fMatch[1]

	ext := t.Filename[strings.LastIndex(t.Filename, "."):]

	return fmt.Sprintf("%s-%s-%s%s", trackNum, Sanitize(t.Artist), Sanitize(t.Title), ext)
}

// Sanitize takes a string and removes problematic characters from it.
func Sanitize(s string) string {
	s = strings.Map(func(r rune) rune {
		if r == '(' || r == ')' || r == '[' || r == ']' || r == '.' {
			return -1
		}
		if unicode.IsSpace(r) {
			return '_'
		}
		return r
	}, s)
	return s
}