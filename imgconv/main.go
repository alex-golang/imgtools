// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/jteeuwen/imgtools/lib"
	"image"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	file, format, options := parseArgs()
	in, out := getStreams(file)

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
func getStreams(file string) (in io.Reader, out io.Writer) {
	if len(file) == 0 {
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
func parseArgs() (string, string, string) {
	target := flag.String("type", "", "")
	optstr := flag.String("options", "", "")
	version := flag.Bool("version", false, "")

	flag.Usage = usage
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
		return "", *target, *optstr
	}

	return filepath.Clean(flag.Args()[0]), *target, *optstr
}

func usage() {
	fmt.Printf(`Usage: %s [options] <path>
   or: cat <path> | %s [options]

 -version
    Displays version information.

 -type <name>
    Name of target image format: %s

 -options <string>
    A semi-colon-separated list of key/value pairs.
    These pairs specify encoder properties specific to the
    output image type. For example:
    
        -options "key1: value1; key2: value2; ...; keyN: valueN"
    
    Known option key names for a given encoder:

`, AppName, AppName, strings.Join(lib.Formats(), ", "))

	for _, enc := range lib.Encoders {
		keys := enc.Keys()

		if len(keys) == 0 {
			continue
		}

		fmt.Printf("        %s: %s\n", enc.Name, strings.Join(keys, ", "))
	}

	fmt.Printf(`
    For example:
      
        %s -type jpeg -options "quality:75" file.png
        %s -type pnm -options "format:P6" file.png

`, AppName, AppName)
}
