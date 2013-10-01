## imgindex

imgindex accepts the path to a given directory.

It constructs perceptual hashes for all supported images it finds.
For every processed image, it outputs the image path and its hash
to `stdout`.

Note that this may take quite a while to complete on large image
repositories. Specially when dealing with very large images.
The reason for this, is that the hashes are computed in a way
which requires the image to be resized. This resize operation is
very expensive.


