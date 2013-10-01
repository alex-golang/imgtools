// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/jteeuwen/imgtools/lib"
	"os"
	"path/filepath"
)

func main() {
	root := parseArgs()

	filepath.Walk(root, walk)
}

func walk(file string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() || !lib.ValidFile(file) {
		return nil
	}

	println(file)

	return nil
}

// parseArgs parses command line arguments.
func parseArgs() string {
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

	return dir
}

func usage() {
	fmt.Printf(`Usage: %s [options] <path>

 -version
    Displays version information.

`, AppName)
}
