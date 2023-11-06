package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/krzko/byte-slice-gen/pkg/imageprocessor"

	"github.com/urfave/cli/v2"
)

func convertToMixedCaps(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '-' || r == '_'
	})

	for i, word := range words {
		if i == 0 {
			words[i] = strings.ToLower(word)
		} else {
			words[i] = title(word)
		}
	}

	return strings.Join(words, "")
}

func ProcessImage(inputPath string) error {
	bytes, err := imageprocessor.ImageToByteSlice(inputPath)
	if err != nil {
		return err
	}

	varName := filepath.Base(inputPath)
	varName = strings.TrimSuffix(varName, filepath.Ext(varName))
	varName = convertToMixedCaps(varName)

	goContents := imageprocessor.GenerateGoFileContents(bytes, varName)

	outputPath := strings.TrimSuffix(inputPath, filepath.Ext(inputPath)) + ".go"
	err = os.WriteFile(outputPath, []byte(goContents), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Generated file: %s\n", outputPath)
	return nil
}

func SetupApp() *cli.App {
	return &cli.App{
		Name:  "image-byte-slice-gen",
		Usage: "Converts image files to Go byte slices",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "Input image file",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			inputPath := c.String("file")
			return ProcessImage(inputPath)
		},
	}
}

func title(word string) string {
	if word == "" {
		return ""
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}
