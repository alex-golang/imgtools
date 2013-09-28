// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/jteeuwen/imgtools/lib"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	file, target, pipe := parseArgs()
	println(file, target, pipe)
}

// parseArgs parses command line arguments.
func parseArgs() (string, string, bool) {
	target := flag.String("type", "",
		fmt.Sprintf("Name of target image type: %s", strings.Join(lib.Formats(), ", ")))
	version := flag.Bool("version", false, "Display version information.")
	flag.Parse()

	if *version {
		fmt.Printf("%s\n", Version())
		os.Exit(0)
	}

	if len(*target) == 0 {
		fmt.Fprintf(os.Stderr, "Missing target image format.\n")
		os.Exit(1)
	}

	if !lib.Supported(*target) {
		fmt.Fprintf(os.Stderr, "Invalid target image format: %s\n", *target)
		os.Exit(1)
	}

	if flag.NArg() == 0 {
		return "", *target, true
	} else {
		return filepath.Clean(flag.Args()[0]), *target, false
	}
}
