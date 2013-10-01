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
	"github.com/jteeuwen/resize"
	"image"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, strwidth, strheight, filter := parseArgs()

	src := load(file)

	width := realSize(src.Bounds().Dx(), strwidth)
	height := realSize(src.Bounds().Dy(), strheight)
	dst := resize.Resize(width, height, src, filter)

	save(dst)
}

// realSize returns the given size string as an integer.
// If it carries a percentage sign, this will return
// the appropriate size, relative to the given input image.
func realSize(imgsize int, str string) uint {
	if str == "" || str == "0" || str == "0%" {
		return 0
	}

	var percent bool
	if strings.HasSuffix(str, "%") {
		str = str[:len(str)-1]
		percent = true
	}

	n, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if percent {
		v := float64(imgsize) * 0.01 * float64(n)
		return uint(math.Ceil(v))
	}

	return uint(n)
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

// save encodes the given image as PNG and saves it to stdout.
func save(img image.Image) {
	err := lib.Encode(os.Stdout, "png", img, "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Encode image: %v\n", err)
		os.Exit(1)
	}
}

// parseArgs parses command line arguments.
func parseArgs() (string, string, string, resize.InterpolationFunction) {
	width := flag.String("width", "0", "")
	height := flag.String("height", "0", "")
	filter := flag.String("filter", "", "")
	version := flag.Bool("version", false, "")

	flag.Usage = usage
	flag.Parse()

	if *version {
		fmt.Printf("%s\n", Version())
		os.Exit(0)
	}

	if len(*filter) == 0 {
		fmt.Fprintf(os.Stderr, "Missing interpolation algorithm.\n")
		flag.Usage()
		os.Exit(1)
	}

	var interp resize.InterpolationFunction
	switch strings.ToLower(*filter) {
	case "nearestneighbor":
		interp = resize.NearestNeighbor
	case "bilinear":
		interp = resize.Bilinear
	case "bicubic":
		interp = resize.Bicubic
	case "mitchellnetravali":
		interp = resize.MitchellNetravali
	case "lanczos2lut":
		interp = resize.Lanczos2Lut
	case "lanczos2":
		interp = resize.Lanczos2
	case "lanczos3lut":
		interp = resize.Lanczos3Lut
	case "lanczos3":
		interp = resize.Lanczos3

	default:
		fmt.Fprintf(os.Stderr, "Unknown interpolation algorithm: %s\n", *filter)
		os.Exit(1)
	}

	if flag.NArg() == 0 {
		return "", *width, *height, interp
	}

	return flag.Args()[0], *width, *height, interp
}

func usage() {
	fmt.Printf(`Usage: %s [options] <path>
   or: cat <path> | %s [options]

 -version
    Displays version information.

 -width <N>
    Width of the target image, in pixels or percentage.
    
    When resizing an image, the original aspect ratio can be
    preserved  by leaving either width or height as 0.

 -height <N>
    Height of the target image, in pixels or percentage.
    
    When resizing an image, the original aspect ratio can be
    preserved  by leaving either width or height as 0.

 -filter <name>
    Name of the interpolation algorithm to use.
    Available algorithms, in order of fastest to slowest, are:

    * NearestNeighbor
    * Bilinear
    * Bicubic
    * MitchellNetravali
    * Lanczos2Lut
    * Lanczos2
    * Lanczos3Lut
    * Lanczos3

    Which of these gives the best results, depends on the input
    image and your use case.

`, AppName, AppName)

}
