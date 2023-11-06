package imageprocessor

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"os"
)

func GenerateGoFileContents(byteSlice []byte, varName string) string {
	contents := "package icon\n\n"
	contents += fmt.Sprintf("var %s = []byte{", varName)
	for i, b := range byteSlice {
		if i%12 == 0 && i > 0 {
			contents += "\n\t"
		}
		contents += fmt.Sprintf("%#x, ", b)
	}
	contents += "\n}\n"
	return contents
}

func ImageToByteSlice(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
