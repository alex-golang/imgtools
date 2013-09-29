// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	parseArgs()
}

// parseArgs parses command line arguments.
func parseArgs() {
	version := flag.Bool("version", false, "")

	flag.Usage = usage
	flag.Parse()

	if *version {
		fmt.Printf("%s\n", Version())
		os.Exit(0)
	}
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

 -saturation <number>
    Adjusts the image color saturation to the given levels.
    Its value must be in the range 0-100.

 -brightness <number>
    Adjusts the image brightness to the given levels.
    Its value must be in the range 0-100.


 Numbers can be specified in various formats and they can include
 a positive or negative sign, as well as a percentile indicator.

    Binary      : 2#11111111
    Octal       : 8#377
    Decimal     : 10#255 or 255
    Hexadecimal : 16#ff

An positive increment must be prefixed with '+' and a negative one
with '-'. A percentile value must have a '%%' suffix:

    Absolute value     :  50 
    Positive increment : +50
    Negative increment : -50
    Percentile         :  50%%


A color value comes as a list of comma-separated RGBA values,
or a single hexadecimal string:

    rgba : 255,153,0,255
    hex  : 16#ff9900ff

`, AppName, AppName)

}
