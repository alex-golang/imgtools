// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package pnm

import (
	"fmt"
	"github.com/jteeuwen/imgtools/lib"
	"github.com/jteeuwen/pnm"
	"image"
	"io"
	"strings"
)

func init() {
	lib.RegisterEncoder("pnm", encode, "format")
}

func encode(w io.Writer, m image.Image, options lib.OptionSet) error {
	var ptype pnm.PNMType
	value := options.String("format", "")

	if len(value) == 0 {
		return fmt.Errorf("Missing option 'format'; expected p1, p2, p3, p4, p5, p6")
	}

	switch strings.ToLower(value) {
	case "p1":
		ptype = pnm.BitmapAscii
	case "p2":
		ptype = pnm.GraymapAscii
	case "p3":
		ptype = pnm.PixmapAscii
	case "p4":
		ptype = pnm.BitmapBinary
	case "p5":
		ptype = pnm.GraymapBinary
	case "p6":
		ptype = pnm.PixmapBinary
	default:
		return fmt.Errorf("Invalid option 'format:%q'; expected p1, p2, p3, p4, p5, p6", value)
	}

	return pnm.Encode(w, m, ptype)
}
