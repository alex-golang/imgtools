// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/jteeuwen/imgtools/lib"
	hash "github.com/jteeuwen/imgtools/imghash/lib"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root, hf := parseArgs()

	filepath.Walk(root, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !lib.ValidFile(file) {
			return nil
		}

		hash, err := getHash(hf, file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", file, err)
			return nil
		}

		fmt.Fprintf(os.Stdout, "%d %s\n", hash, file)
		return nil
	})
}

// getHash creates a perceptual hash for the given file.
func getHash(hf hash.HashFunc, file string) (uint64, error) {
	fd, err := os.Open(file)
	if err != nil {
		return 0, err
	}

	defer fd.Close()

	img, _, err := lib.Decode(fd)
	if err != nil {
		return 0, err
	}

	return hf(img), nil
}

// parseArgs parses command line arguments.
func parseArgs() (string, hash.HashFunc) {
	hashname := flag.String("hash", "average", "")
	version := flag.Bool("version", false, "")

	flag.Usage = usage
	flag.Parse()

	if *version {
		fmt.Printf("%s\n", Version())
		os.Exit(0)
	}

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "Missing directory path.\n")
		flag.Usage()
		os.Exit(1)
	}

	dir := filepath.Clean(flag.Args()[0])

	stat, err := os.Lstat(dir)
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Path %q does not exist.\n", dir)
		os.Exit(1)
	}

	if !stat.IsDir() {
		fmt.Fprintf(os.Stderr, "Path %q is not a directory.\n", dir)
		os.Exit(1)
	}

	var hf hash.HashFunc
	switch strings.ToLower(*hashname) {
	case "average":
		hf = hash.Average
	default:
		fmt.Fprintf(os.Stderr, "Unknown hash function: %s\n", *hashname)
		flag.Usage()
		os.Exit(1)
	}

	return dir, hf
}

func usage() {
	fmt.Printf(`Usage: %s [options] <path>

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

`, AppName)
}
