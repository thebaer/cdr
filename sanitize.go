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

// RenameTrack takes a filename, opens it, reads the metadata, and returns both
// the old and new filename.
func RenameTrack(file string) string {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("error loading file: %v", err)
		return ""
	}
	defer f.Close()

	fName := f.Name()
	fMatch := trackNameReg.FindStringSubmatch(fName)
	if len(fMatch) < 2 {
		log.Fatal("Unexpect filename format")
	}
	trackNum := fMatch[1]
	ext := fName[strings.LastIndex(fName, "."):]

	m, err := tag.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s-%s-%s%s", trackNum, Sanitize(m.Artist()), Sanitize(m.Title()), ext)
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
