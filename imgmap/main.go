// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	input, expr := parseArgs()
	in, out := getStreams(input)

	_, _, _ = in, out, expr
}

// getStreams opens ands returns the input and output streams.
func getStreams(input string) (in io.ReadCloser, out io.Writer) {
	if len(input) == 0 {
		return os.Stdin, os.Stdout
	}

	fd, err := os.Open(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid input file: %v\n", err)
		os.Exit(1)
	}

	return fd, os.Stdout
}

// parseArgs parses command line arguments.
func parseArgs() (string, []byte) {
	version := flag.Bool("version", false, "")
	mapfile := flag.String("map", "", "")
	expr := flag.String("expr", "", "")

	flag.Usage = usage
	flag.Parse()

	if *version {
		fmt.Printf("%s\n", Version())
		os.Exit(0)
	}

	if len(*mapfile) == 0 && len(*expr) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	data := []byte(*expr)

	if len(*mapfile) > 0 {
		fd, err := os.Open(filepath.Clean(*mapfile))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		defer fd.Close()

		var buf bytes.Buffer
		io.Copy(&buf, fd)
		data = buf.Bytes()
	}

	data = bytes.TrimSpace(data)
	sz := len(data)

	if sz == 0 {
		fmt.Fprintf(os.Stderr, "Invalid map expression specified; no data.\n")
		os.Exit(1)
	}

	if data[sz-1] != '\n' {
		data = append(data, '\n')
	}

	if flag.NArg() > 0 {
		return flag.Args()[0], data
	}

	return "", data
}

func usage() {
	fmt.Printf(`Usage: %s [options] <path>
   or: cat <path> | %s [options]

 -version
    Displays version information.

 -map <file>
    Path to a text file with color map expressions.

 -expr <expression>
    A mapping expression which should be executed as-is.
    This is intended for simple, one-off operations you
    do not want to create a separate mapping file for.

`, AppName, AppName)

}
