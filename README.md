# pdf-compressor

This is a cross platform frontend for ghostscript that will use common sense defaults to reduce the size of your PDF file for use in print, web, and email contexts.

## Usage Flags

```
-aspect-ratio-x float
    X aspect ratio. Defaults to 16:9. (default 16)
-aspect-ratio-y float
    Y aspect ratio. Defaults to 16:9. (default 9)
-compatibility-level string
    PDF Compatibility Level. Reduce if sharing broadly. 1.3 - 1.7 are supported. (default "1.7")
-i string
    Path to the input file
-lossiness string
    Use Lossy or Lossless image compression. Lossy images are much smaller but tend to have (default "lossless")
-resize-output
    Whether to resize output or not. For presentation decks this should always be true. (default true)
-size string
    Size of output file. Valid choices are small, medium, or large (default "medium")
```