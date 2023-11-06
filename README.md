# image-byte-slice-gen

A Go-based command-line tool to convert images into [Go](https://go.dev/) byte slice files.

## Usage

Run the utility with the `-f` flag followed by the path to the image file you want to convert:

```sh
image-byte-slice-gen -f path/to/your/image.png
```

This will generate a Go source file with the same name as the image file, containing a byte slice with the image data.

```go
// image.go
package icon

var image = []byte{...}
```
