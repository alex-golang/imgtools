// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/jteeuwen/imgtools/lib"
	_ "github.com/jteeuwen/imgtools/lib/gif"
	_ "github.com/jteeuwen/imgtools/lib/jpeg"
	_ "github.com/jteeuwen/imgtools/lib/png"
	_ "github.com/jteeuwen/imgtools/lib/pnm"
	"image"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	file, pipe, format, options := parseArgs()
	in, out := getStreams(file, pipe)

	if !pipe {
		defer in.Close()
	}

	img, _, err := image.Decode(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid input file: %v\n", err)
		return
	}

	err = lib.Encode(out, format, img, options)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Encode failed: %v\n", err)
		return
	}
}

// getStreams opens ands returns the input and output streams.
func getStreams(file string, pipe bool) (in io.ReadCloser, out io.Writer) {
	if pipe {
		return os.Stdin, os.Stdout
	}

	fd, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid input file: %v\n", err)
		os.Exit(1)
	}

	return fd, os.Stdout
}

// parseArgs parses command line arguments.
func parseArgs() (string, bool, string, string) {
	target := flag.String("type", "",
		fmt.Sprintf("Name of target image type: %s", strings.Join(lib.Formats(), ", ")))
	optstr := flag.String("options", "", "List of encoder options.")
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

	if flag.NArg() == 0 {
		return "", true, *target, *optstr
	}

	return filepath.Clean(flag.Args()[0]), false, *target, *optstr
}
