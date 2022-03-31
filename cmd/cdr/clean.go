package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/thebaer/cdr"
	"github.com/urfave/cli/v2"
)

var cmdClean = cli.Command{
	Name:   "clean",
	Usage:  "clean and organize audio files in the current directory",
	Action: cleanAction,
}

func cleanAction(c *cli.Context) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	filepath.Walk(wd, func(path string, i os.FileInfo, err error) error {
		if !i.IsDir() && !strings.HasPrefix(i.Name(), ".") {
			fName := i.Name()
			trackName := cdr.RenameTrack(fName)
			if trackName == "" {
				return nil
			}
			fmt.Println("Renaming", fName, "to", trackName)
			os.Rename(fName, trackName)
		}

		return nil
	})
	return nil
}
