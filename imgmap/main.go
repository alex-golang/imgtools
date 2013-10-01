// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/jteeuwen/imgtools/lib"
	"image"
	"image/draw"
	"io"
	"os"
	"path/filepath"
)

var (
	semicolon = []byte{';'}
	tab       = []byte{'\t'}
	space     = []byte{' '}
)

func main() {
	var linecount int
	var line []byte
	var err error

	file, expr := parseArgs()
	src, dst := getImages(file)

	for err != io.EOF {
		linecount++
		line, err = expr.ReadBytes('\n')

		if err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "Read expression: %v\n", err)
			os.Exit(1)
		}

		if parseExpression(linecount, line, src, dst) {
			src, dst = dst, src
		}
	}

	err = lib.Encode(os.Stdout, "png", src, "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Write output image: %v\n", err)
		os.Exit(1)
	}
}

// getStreams opens ands returns the input and output streams.
func getImages(input string) (draw.Image, draw.Image) {
	var fd io.ReadCloser
	var err error

	if len(input) == 0 {
		fd = os.Stdin
	} else {
		fd, err = os.Open(input)
		defer fd.Close()
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Open input file: %v\n", err)
		os.Exit(1)
	}

	img, _, err := lib.Decode(fd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Decode image: %v\n", err)
		os.Exit(1)
	}

	dimg, ok := img.(draw.Image)
	if !ok {
		fmt.Fprintf(os.Stderr, "Decode image: %T is not draw.Image; missing Set method.\n", img)
		os.Exit(1)
	}

	return dimg, image.NewRGBA(img.Bounds())
}

// parseArgs parses command line arguments.
func parseArgs() (string, *bufio.Reader) {
	var err error
	var data io.Reader

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

	data = bytes.NewBufferString(*expr)

	if len(*mapfile) > 0 {
		data, err = os.Open(filepath.Clean(*mapfile))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}

	if flag.NArg() > 0 {
		return flag.Args()[0], bufio.NewReader(data)
	}

	return "", bufio.NewReader(data)
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
