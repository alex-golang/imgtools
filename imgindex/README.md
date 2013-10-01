## imgindex

imgindex accepts the path to a given directory.
It constructs perceptual hashes for all supported images it
finds and saves them in a database. This database can be used to quickly
find similar images using the `imgfind` tool.

This command should be run at least once before using `imgfind`.
After that, one should run it whenever the images in the given path
change or new ones are added. This is kept separate from `imgfind`,
because it takes a while to hash large libraries. Specially with
large image files. The tool is smart enough not to re-calculate hashes
for files which have not changed. So repeatedly running the tool over
the same set of files, will only update those files which have been
altered since the last time it was run.

The generated database is a flat list of entries, written to stdout.


