// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	cfg := parseArgs()
	in, out := getStreams(cfg)

	if !cfg.Pipe() {
		defer in.Close()
	}

	_, _ = in, out

	c := cfg.Monochrome
	fmt.Printf("%T %v\n", c.R, c.R)
	fmt.Printf("%T %v\n", c.G, c.G)
	fmt.Printf("%T %v\n", c.B, c.B)
	fmt.Printf("%T %v\n", c.A, c.A)
}

// getStreams opens ands returns the input and output streams.
func getStreams(c *Config) (in io.ReadCloser, out io.Writer) {
	if c.Pipe() {
		return os.Stdin, os.Stdout
	}

	fd, err := os.Open(c.Input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid input file: %v\n", err)
		os.Exit(1)
	}

	return fd, os.Stdout
}

// parseArgs parses command line arguments.
func parseArgs() *Config {
	var cfg Config
	var err error

	monochrome := flag.String("monochrome", "", "")
	saturate := flag.String("saturate", "", "")
	brightness := flag.String("brightness", "", "")
	version := flag.Bool("version", false, "")

	flag.StringVar(&cfg.MapFile, "map", "", "")
	flag.BoolVar(&cfg.Grayscale, "grayscale", false, "")
	flag.BoolVar(&cfg.BlackWhite, "bw", false, "")
	flag.Usage = usage
	flag.Parse()

	if *version {
		fmt.Printf("%s\n", Version())
		os.Exit(0)
	}

	if flag.NArg() > 0 {
		cfg.Input = flag.Args()[0]
	}

	if len(cfg.MapFile) == 0 && len(*monochrome) == 0 &&
		len(*saturate) == 0 && len(*brightness) == 0 &&
		!cfg.Grayscale && !cfg.BlackWhite {
		flag.Usage()
		os.Exit(1)
	}

	if len(*monochrome) > 0 {
		cfg.Monochrome, err = ParseColor(*monochrome)
	}

	if len(*saturate) > 0 {
		cfg.Saturate, err = parseChannel(*saturate)
	}

	if len(*brightness) > 0 {
		cfg.Brightness, err = parseChannel(*brightness)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	return &cfg
}

func usage() {
	fmt.Printf(`Usage: %s [options] <path>
   or: cat <path> | %s [options]

 -version
    Displays version information.

 -map <file>
    Textfile with custom color mappings.

 -bw
    Convert the input image to a black & white bitmap.

 -grayscale
    Converts the input image to grayscale.

 -monochrome <color>
    Same as grayscale, but allows specification of
    a custom colorspace.

 -saturate <number>
    Adjusts the image color saturation to the given level.

 -brightness <number>
    Adjusts the image brightness to the given level.


 Numbers can be specified in decimal and hexadecima; formats.
 They can carry a positive or negative sign, as well as a
 percentile indicator.

 A positive increment must be prefixed with '+' and a negative one
 with '-'. A percentile value must have a '%%' suffix:

    Absolute value     :  50 
    Positive increment : +50
    Negative increment : -50
    Percentile         :  50%%


 A color value comes as a list of comma-separated RGBA values,
 or a single hexadecimal string:

    rgba : 255,153,0,255
    hex  : 0xff9900ff

`, AppName, AppName)

}
