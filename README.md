## imgtools

**imgtools** holds a set of command line tools intended for
image manipulation. Each tool performs a specific action on
a given input image.

An image can be supplied as a file path in a command line argument,
or its contents can be piped in through `stdin`. The latter allows all
commands to be chained together using pipes.

**Note**: This is work in progres. Some commands are not implemented yet
and others may benefit from optimization and feature enhancements.


### TODO

* Implement GIF encoder.
* Implement `imgquantize` tool.


### Tools

* **imgscale**: Resizes the given image to a new size.
* **imgconv**: Saves the image as a different image type.
* **imgmap**: Remaps specified colors in the input image to a set of new colors.
* **imgquantize**: Provides a set of different image quantization algorithms.


### Image types

The tools support all image types supported by the Go standard library,
as well as some externally defined formats:

* PNG
* Jpeg
* Gif
* PNM


### Usage

    go get github.com/jteeuwen/imgtools/...


### Dependencies

	go get github.com/jteeuwen/pnm


### Documentation

Documentation can be found at [godoc.org](http://godoc.org/github.com/jteeuwen/imgtools).


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

