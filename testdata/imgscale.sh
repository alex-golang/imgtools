#!/usr/bin/env bash

IMG="original.png"

cat $IMG | imgscale -width 200% -filter nearestneighbor   >  "scale_up_nearestneighbor.png"
cat $IMG | imgscale -width 200% -filter bilinear          >  "scale_up_bilinear.png"
cat $IMG | imgscale -width 200% -filter bicubic           >  "scale_up_bicubic.png"
cat $IMG | imgscale -width 200% -filter mitchellnetravali >  "scale_up_mitchellnetravali.png"
cat $IMG | imgscale -width 200% -filter lanczos2lut       >  "scale_up_lanczos2lut.png"
cat $IMG | imgscale -width 200% -filter lanczos2          >  "scale_up_lanczos2.png"
cat $IMG | imgscale -width 200% -filter lanczos3lut       >  "scale_up_lanczos3lut.png"
cat $IMG | imgscale -width 200% -filter lanczos3          >  "scale_up_lanczos3.png"

cat $IMG | imgscale -width 50% -filter NearestNeighbor   >  "scale_down_nearestneighbor.png"
cat $IMG | imgscale -width 50% -filter bilinear          >  "scale_down_bilinear.png"
cat $IMG | imgscale -width 50% -filter bicubic           >  "scale_down_bicubic.png"
cat $IMG | imgscale -width 50% -filter mitchellnetravali >  "scale_down_mitchellnetravali.png"
cat $IMG | imgscale -width 50% -filter lanczos2lut       >  "scale_down_lanczos2lut.png"
cat $IMG | imgscale -width 50% -filter lanczos2          >  "scale_down_lanczos2.png"
cat $IMG | imgscale -width 50% -filter lanczos3lut       >  "scale_down_lanczos3lut.png"
cat $IMG | imgscale -width 50% -filter lanczos3          >  "scale_down_lanczos3.png"

