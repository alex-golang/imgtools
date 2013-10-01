#!/usr/bin/env bash

IMG="original.png"

cat $IMG | imgmap -expr "? ? ? ?   ? 0 0 255"        > "map_red.png"
cat $IMG | imgmap -expr "? ? ? ?   0 ? 0 255"        > "map_green.png"
cat $IMG | imgmap -expr "? ? ? ?   0 0 ? 255"        > "map_blue.png"
cat $IMG | imgmap -expr "? ? ? ?   #l #l #l ?"       > "map_lightness.png"
cat $IMG | imgmap -expr "? ? ? ?   #A #A #A ?"       > "map_average.png"
cat $IMG | imgmap -expr "? ? ? ?   #L #L #L ?"       > "map_luminosity.png"
cat $IMG | imgmap -expr "? ? ? ?   +20 +20 +20 ?"    > "map_lightena.png"
cat $IMG | imgmap -expr "? ? ? ?   +20% +20% +20% ?" > "map_lightenb.png"
cat $IMG | imgmap -expr "? ? ? ?   -20 -20 -20 ?"    > "map_darkena.png"
cat $IMG | imgmap -expr "? ? ? ?   -20% -20% -20% ?" > "map_darkenb.png"
cat $IMG | imgmap -expr "<200 ? ? ?  #L #L #L ?"     > "map_yellow.png" 

