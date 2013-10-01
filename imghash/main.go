// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"flag"
	"fmt"
	hashlib "github.com/jteeuwen/imgtools/imghash/lib"
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
	file, hash := parseArgs()
	src := load(file)

	fmt.Fprintf(os.Stdout, "%d\n", hash(src))
}

// load loads the given image.
func load(input string) image.Image {
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

	return img
}

// parseArgs parses command line arguments.
func parseArgs() (string, hashlib.HashFunc) {
	hash := flag.String("hash", "average", "")
	version := flag.Bool("version", false, "")

	flag.Usage = usage
	flag.Parse()

	if *version {
		fmt.Printf("%s\n", Version())
		os.Exit(0)
	}

	var file string
	var hf hashlib.HashFunc

	if flag.NArg() > 0 {
		file = filepath.Clean(flag.Args()[0])
	}

	if len(*hash) == 0 {
		fmt.Fprintf(os.Stderr, "Missing hash function.\n")
		os.Exit(1)
	}

	switch strings.ToLower(*hash) {
	case "average":
		hf = hashlib.Average
	default:
		fmt.Fprintf(os.Stderr, "Unknown hash function: %s\n", *hash)
		flag.Usage()
		os.Exit(1)
	}

	return file, hf
}

func usage() {
	fmt.Printf(`Usage: %s [options] <path>
   or: cat <path> | %s [options]

 -version
    Displays version information.

 -hash <name>
    Name of the hash function to use.
    Known hash functions include:
    
    - Average: Average computes a Perceptual Hash using a naive,
      but very fast method. It holds up to minor colour changes,
      changing brightness and contrast and is indifferent to
      aspect ratio and image size differences.
  
      Average Hash is a great algorithm if you are looking for
      something specific. For example, if we have a small thumbnail
      of an image and we wish to know if the original exists
      somewhere in our collection. Average Hash will find  it very
      quickly. However, if there are modifications -- like text
      was added or a head was spliced into place, then Average
      Hash probably won't do the job.
  
      The Average Hash is quick and easy, but it can generate false
      misses if gamma correction or color histogram is applied to
      the image. This is because the colors move along a non-linear
      scale -- changing where the "average" is located and therefore
      changing which bits are above/below the average.

`, AppName, AppName)
}
