#!/usr/bin/env bash

# Test if two given images are to be considered equal.
# This works even if they have different sizes and aspect ratios.

A="original.png"
B="original_thumb.png"

HASHA=$(imghash -hash average $A)
HASHB=$(imghash -hash average $B)

DISTANCE=$(imghash -diff $HASHA $HASHB)
EPSILON=3

if [ $DISTANCE -le $EPSILON ]; then
	echo "$A and $B are equal"
else
	echo "$A and $B are not equal"
fi


