package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/thebaer/cdr"
)

var printUsage = func() {
	fmt.Fprintf(os.Stderr, "usage: %s [optional flags] filename\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = printUsage
	flag.Parse()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filepath.Walk(wd, func(path string, i os.FileInfo, err error) error {
		if !i.IsDir() && !strings.HasPrefix(i.Name(), ".") {
			fName := i.Name()
			trackName := cdr.RenameTrack(fName)
			fmt.Println("Renaming", fName, "to", trackName)
			os.Rename(fName, trackName)
		}

		return nil
	})
}
