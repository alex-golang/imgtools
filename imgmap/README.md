## imgmap

imgmap remaps certain colours in an input image to new values.
Color mappings are defined as one or more expressions, either
in an external text file, supplied through the `-map` command line
argument, or as a single expression in the `-expr` command line argument.


### Mappings

A color map comes in the form of a list of whitespace-separated source- and
destination colors. Each color is represented as a whitespace-separated set
of RGBA values. For example, to turn red pixels into blue, the following map
is used:

	255 0 0 255   0 255 0 255

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
* **#l**: The RGB lightness: Averages the most prominent amd least
  prominent channel values: `(max(r, g, b) + min(r + g + b)) / 2`
* **#L**: The RGB luminosity is a more sophisticated version of
  the average method. It also averages the values, but it forms a
  weighted average to account for human perception. We're more
  sensitive to green than other colors, so green is weighted
  most heavily. The formula for this is as follows:

```go
	lum := gamma_sRGB(
		0.212655 * inverse_gamma_sRGB(R) +
		0.715158 * inverse_gamma_sRGB(G) +
		0.072187 * inverse_gamma_sRGB(B)
	)
```

Swapping color channels can be done by referencing a channel by
its named placeholder in the destination color. For instance,
to swap all red and blue channels and leave the rest as-is, we can use:

    ? ? ? ?    #b ? #r ?

To turn all pixels into their grayscale equivalents, while setting
alpha to max, we can use:

	? ? ? ?    #L #L #L 255


### Numbers

So far, the examples show the use of base-10 numbers as color channel values.
The parser also supports hexadecimal values: `0xff`. For example, the following
two lines are functionally equivalent:

	255 153 0 255
	0xff 0x99 0x00 0xff


### Comments

A mapping file can contain comments. These are prefixed with `;`
and span the remainder of the line.


### Examples

<table>
<tr><th>Original image</th></tr>
<tr><th><img src="http://jteeuwen.nl/img/imgtools/original.png" /></th></tr>
</table>

<table>
<tr><th colspan="3">Filter color channels</th></tr>
<tr>
<th><img src="http://jteeuwen.nl/img/imgtools/red.png" /><br>red</th>
<th><img src="http://jteeuwen.nl/img/imgtools/green.png" /><br>green</th>
<th><img src="http://jteeuwen.nl/img/imgtools/blue.png" /><br>blue</th>
</tr>
<tr><td colspan="3"><pre>
cat orig.png | imgmap -expr "? ? ? ?   ? 0 0 255" > red.png
cat orig.png | imgmap -expr "? ? ? ?   0 ? 0 255" > green.png
cat orig.png | imgmap -expr "? ? ? ?   0 0 ? 255" > blue.png
</pre></td></tr></table>

<table>
<tr><th colspan="3">Grayscale conversions</th></tr>
<tr>
<th><img src="http://jteeuwen.nl/img/imgtools/average.png" /><br>average</th>
<th><img src="http://jteeuwen.nl/img/imgtools/lightness.png" /><br>lightness</th>
<th><img src="http://jteeuwen.nl/img/imgtools/luminosity.png" /><br>luminosity</th>
</tr>
<tr><td colspan="3"><pre>
cat orig.png | imgmap -expr "? ? ? ?   #l #l #l ?" > lightness.png
cat orig.png | imgmap -expr "? ? ? ?   #A #A #A ?" > average.png
cat orig.png | imgmap -expr "? ? ? ?   #L #L #L ?" > luminosity.png
</pre></td></tr></table>

<table>
<tr><th colspan="4">Lighten/darken</th></tr>
<tr>
<th><img src="http://jteeuwen.nl/img/imgtools/lightena.png" /><br>increase by 20</th>
<th><img src="http://jteeuwen.nl/img/imgtools/lightenb.png" /><br>increase by 20%</th>
<th><img src="http://jteeuwen.nl/img/imgtools/darkena.png" /><br>decrease by 20</th>
<th><img src="http://jteeuwen.nl/img/imgtools/darkenb.png" /><br>decrease by 20%</th>
</tr>
<tr><td colspan="4"><pre>
cat orig.png | imgmap -expr "? ? ? ?   +20 +20 +20 ?"    > lightena.png
cat orig.png | imgmap -expr "? ? ? ?   +20% +20% +20% ?" > lightenb.png
cat orig.png | imgmap -expr "? ? ? ?   -20 -20 -20 ?"    > darkena.png
cat orig.png | imgmap -expr "? ? ? ?   -20% -20% -20% ?" > darkenb.png
</pre></td></tr></table>

<table>
<tr><th>Custom filter</th></tr>
<tr>
<th><img src="http://jteeuwen.nl/img/imgtools/custom.png" /><br>Bring out the flower</th>
</tr>
<tr><td colspan="4"><pre>
cat orig.png | imgmap -expr "&lt;200 ? ? ?  #L #L #L ?" > custom.png
</pre></td></tr></table>


