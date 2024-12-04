package main

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/webp"
)

func main() {
	// Directory containing .webp files
	dir := "../go2d/assets" // Change this to your directory path

	// Process all files in the directory
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check if the file is a .webp file
		if strings.ToLower(filepath.Ext(info.Name())) == ".webp" {
			// Convert the file
			err := convertWebPToPNG(path)
			if err != nil {
				fmt.Printf("Failed to convert %s: %v\n", path, err)
			} else {
				fmt.Printf("Successfully converted %s\n", path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
	}
}

func convertWebPToPNG(inputPath string) error {
	// Open the input WebP file
	webpFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer webpFile.Close()

	// Decode the WebP image
	img, err := webp.Decode(webpFile)
	if err != nil {
		return fmt.Errorf("failed to decode WebP image: %v", err)
	}

	// Create the output PNG file path
	outputPath := strings.TrimSuffix(inputPath, filepath.Ext(inputPath)) + ".png"

	// Create the output PNG file
	pngFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer pngFile.Close()

	// Encode the image as PNG
	err = png.Encode(pngFile, img)
	if err != nil {
		return fmt.Errorf("failed to encode PNG image: %v", err)
	}

	return nil
}
