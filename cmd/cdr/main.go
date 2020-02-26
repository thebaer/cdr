package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/thebaer/cdr"
)

var printUsage = func() {
	fmt.Fprintf(os.Stderr, "usage: %s [optional flags] filename\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = printUsage
	flag.Parse()
	if flag.NArg() != 1 {
		printUsage()
		return
	}

	file := flag.Arg(0)
	oldFilename, trackName := cdr.RenameTrack(file)
	fmt.Println("Renaming", oldFilename, "to", trackName)
	os.Rename(oldFilename, trackName)
}
