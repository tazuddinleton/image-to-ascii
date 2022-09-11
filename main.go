package main

import (
	"fmt"
	"image"
	png "image/png"
	"os"

	draw "golang.org/x/image/draw"
)

var (
	chars  = "Ñ@#W$9876543210?!abc;:+=-,._ "
	runes  = []rune(chars)
	length = len(runes)
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Provide image path")

		os.Exit(0)
	}

	// Todo: Accept size as argument
	path := args[0]
	imageToASCII(path)
}

// Todo: Move to its own package, handle other image formats
func imageToASCII(path string) {
	f, _ := os.Open(path)
	g, _ := png.Decode(f)

	ascImage := []string{}
	for i := 0; i < g.Bounds().Dx(); i++ {
		for j := 0; j < g.Bounds().Dy(); j++ {
			clr := g.At(i, j)
			r, g, b, _ := clr.RGBA()
			t := avg(r, g, b)
			if t > 255 {
				t = 255
				fmt.Println(t)
			}

			charIndex := mapVal(t, 0, 255, 0, length-1)
			c := fmt.Sprintf("%c", runes[charIndex])
			ascImage = append(ascImage, c)
		}
		ascImage = append(ascImage, "\n")
	}

	drawASCII(ascImage)
}

func drawASCII(img []string) {
	for _, pix := range img {
		fmt.Print(" ", pix)
	}
	fmt.Println()
}

// Todo: Map mapVal to util packaage and rename
func mapVal(val, inStart, inEnd, outStart, outEnd int) int {
	slope := 1.0 * (float32(outEnd-outStart) / float32(inEnd-inStart))
	out := float32(outStart) + slope*(float32(val)-float32(inStart))
	return int(out)
}

// Todo: Move avg to util package and rename
func avg(r, g, b uint32) int {
	return int((r/256 + g/256 + b/256) / 3)
}

// Todo: Move resize to its own package
func resize(h, w int, path, name string) {
	file, _ := os.Open(path)
	defer file.Close()
	src, _ := png.Decode(file)

	output, _ := os.Create(fmt.Sprintf("./%s%dx%d.png", name, h, w))

	outImg := image.NewRGBA64(image.Rect(0, 0, w, h))

	draw.CatmullRom.Scale(outImg, outImg.Rect, src, src.Bounds(), draw.Over, nil)

	png.Encode(output, outImg)
}
