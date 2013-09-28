// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

/*
imgconv convcerts one image type to another.
All supported image formats can be the input and output.

Different output formats may require some encoder options.
These can be supplied as a semi-colon separated list of
key/value pairs with the `-options` command line parameter.

For example:

	cat img.png | imgconv -type jpeg -options "quality:75"
	cat img.png | imgconv -type pnm -options "Format:P6"
*/
package main
