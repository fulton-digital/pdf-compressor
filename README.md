# pdf-compressor

This is a cross platform frontend for ghostscript that will use common sense defaults to reduce the size of your PDF file for use in print, web, and email contexts.

## Example Usage

Output a large (but compressed) PDF.

`pdfcompress -i test.pdf -s large`

Use lossy (JPEG) compression instead of Flate.

`pdfcompress -i test.pdf -l lossy`

Make output slides 4:3 aspect ratio.

`pdfcompress -i test.pdf -arx 4 -ary 3`

## Usage Flags

```
-arx float
    X aspect ratio. Defaults to 16:9. (default 16)
-ary float
    Y aspect ratio. Defaults to 16:9. (default 9)
-cl string
    PDF Compatibility Level. Reduce if sharing broadly. 1.3 - 1.7 are supported. (default "1.7")
-i string
    Path to the input file. REQUIRED
-l string
    Use Lossy or Lossless image compression. Lossy images are much smaller but tend to have compression artifacts. (default "lossless")
-nr
    Do not resize output. For presentation decks this should always be omitted.
-s string
    Size of output file. Valid choices are small, medium, or large (default "medium")
```