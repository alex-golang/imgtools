// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"flag"
	"fmt"
	hashlib "github.com/jteeuwen/imgtools/imghash/lib"
	"github.com/jteeuwen/imgtools/lib"
	"image"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	file, hf, a, b := parseArgs()

	if a > 0 {
		distance := hashlib.Distance(a, b)
		fmt.Fprintf(os.Stdout, "%d\n", distance)
	} else {
		src := load(file)
		fmt.Fprintf(os.Stdout, "%d\n", hf(src))
	}
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
func parseArgs() (string, hashlib.HashFunc, uint64, uint64) {
	diff := flag.Bool("diff", false, "")
	hash := flag.String("hash", "average", "")
	version := flag.Bool("version", false, "")

	flag.Usage = usage
	flag.Parse()

	if *version {
		fmt.Printf("%s\n", Version())
		os.Exit(0)
	}

	if *diff {
		if flag.NArg() < 2 {
			fmt.Fprintf(os.Stderr, "Missing hash values.\n")
			flag.Usage()
			os.Exit(1)
		}

		str := flag.Args()[0]
		a, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid hash value: %s\n", str)
			flag.Usage()
			os.Exit(1)
		}

		str = flag.Args()[1]
		b, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid hash value: %s\n", str)
			flag.Usage()
			os.Exit(1)
		}

		return "", nil, a, b
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

	return file, hf, 0, 0
}

func usage() {
	fmt.Printf(`Usage: %s [options] <path>
   or: cat <path> | %s [options]
   or: %s -diff <hash1> <hash2>

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

 -diff
    This option compares the two given hashes and returns their
    Hamming Distance.
    
    Equality of two images is defined as the Hamming
    Distance between two hashes. This distance is a value
    in the range 0-64. Where 0 means the images are
    indentical and 64 means they are completely different.
    To account for minor scaling or aspect ratio artefacts,
    it is generally better to compare this distance to a
    threshold value in order to determine of the images
    are equal or not. For instance, a thumbnail and its
    full-size image version may have a distance of 3 or less.
    The input images are then to be considered equal if
    distance <= 3.

`, AppName, AppName, AppName)
}
