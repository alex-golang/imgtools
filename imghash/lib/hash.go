// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import "image"

// A HashFunc computes a Perceptual Hash for a given image.
type HashFunc func(image.Image) uint64
