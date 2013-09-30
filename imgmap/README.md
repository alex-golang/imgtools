## imgmap

imgmap remaps certain colours to new values. The color mapping is
supplied as a command line parameter or a builtin mapping is supplied.
Its output is always a PNG image.


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


### Named references

Named references are predefined to a limited set of values.
These are **case sensitive**.

* **#r**: Current red channel value.
* **#g**: Current green channel value.
* **#b**: Current blue channel value.
* **#a**: Current alpha channel value.
* **#A**: The average of the RGB channels: `(r + g + b) / 3`
* **#l**: The RGB lightness: Averages the most prominent amd least prominent
  channel values: `(max(r, g, b) + min(r + g + b)) / 2`
* **#L**: The RGB luminosity is a more sophisticated version of
  the average method. It also averages the values, but it forms a
  weighted average to account for human perception. We're more
  sensitive to green than other colors, so green is weighted
  most heavily. This uses the BT709 grayscale weights:
  `(R*0.2125 + G*0.7154 + B*0.0721) / 3`.

Swapping color channels can be done by referencing a channel by
its named placeholder in the destination color. For instance,
to swap all red and blue channels and leave the rest as-is, we can use:

    ? ? ? ?    #b ? #r ?

To turn all pixels into their grayscale equivalents, while setting
alpha to max, we can use:

	? ? ? ?    #m #m #m 255


### Numbers

So far, the examples show the use of base-10 numbers as color channel values.
The parser also supports hexadecimal values: `0xff`. For example, the following
two lines are functionally equivalent:

	255 153 0 255
	0xff 0x99 0x00 0xff


### Comments

A mapping file can contain comments. These are prefixed with `;`
and span the remainder of the line.


### Builtin mappings

The following set of builtin mappings can be used instead
of a custom mapping:

* **blackwhite**: Converts the input image to a black & white bitmap.
* **grayscale**: Converts the image to grayscale.
* **monochrome**: Same as grayscale, but allows specification of a custom colorspace.
* **saturation**: Adjusts the image color saturation to the given levels.
* **brightness**: Adjusts the image brightness to the given levels.


