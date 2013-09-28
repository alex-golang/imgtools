## imgmap

imgmap remaps certain colours to new values. The color mapping is
supplied as a command line parameter or a builtin mapping is supplied.


### Mappings

A color map comes in the form of a list of whitespace-separated source- and
destination colors. Each color is represented as a whitespace-separated set
of RGBA values. For example, to swap red and blue pixels, the following map
is used:

	255 0 0 255   0 255 0 255
	0 255 0 255   255 0 0 255

A wildcard `?` in the source color can be used for any of the color channel
values, to specify that any value for that channel is to be processed.
For example, in order to replace all colors with a red channel of `255`
with black, we use the following:

	255 ? ? ?   0 0 0 255

We can further filter color channels by applying conditionals. For example,
to replace all colors with a blue channel `< 100` with white, we use:

	? <100 ? ?   255 255 255 255

Supported conditionals are:

* `<N`: Channel value is less than N.
* `<=N`: Channel value is less than or equal to N.
* `>N`: Channel value is greater than N.
* `>=N`: Channel value is greater than or equal to N.

In order to change only one color channel and leave the rest as-is,
we can use the wildcard `?` in the destination color. For example,
to clear all red channels and leave the rest intact, we use:

	? ? ? ?    0 ? ? ?


So far any color changes have been to absolute values. We can also
perform a relative mapping. For instance, to reduce all color
channel intensities by 10%, while setting alpha to max, we use:

	? ? ? ?   -10% -10% -10% 255

Similarly, to increase by 10%, we use:

	? ? ? ?   +10% +10% +10% 255

We don't have to use percentages to perform relative changes.
Absolute increases or reductions are accepted as well:

	? ? ? ?   +10 +20 +30 255
	? ? ? ?   -10 -20 -30 255

Note the use of the `+` operator. If we omit this, we are saying
"set this channel to value 10.". With the `+` operator, we are
saying "Add 10 the current channel value".


### Builtin mappings

The following set of builtin mappings can be used instead
of a custom mapping:

* **blackwhite**: Converts the input image to a black & white bitmap.
* **grayscale**: Converts the image to grayscale.
* **monochrome**: Same as grayscale, but allows specification of a custom colorspace.
* **saturation**: Adjusts the image color saturation to the given levels.
* **brightness**: Adjusts the image brightness to the given levels.


### Color values

So far, the examples show the use of base-10 numbers as color channel values.
The parser supports additional bases by prefixing the number with the base. 
With the exception of base-10 (decimal) numbers, the prefixes are mandatory.

Here is the number `255` in all the supported base notations:

* **Binary**: `2#11111111`
* **Octal**: `8#377`
* **Decimal**: `10#255` or `255`
* **Hexadecimal**: `16#ff`



