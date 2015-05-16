package main

import (
	"fmt"
	ft "github.com/Agon/go-freetype"
	"image"
	"image/draw"
	"image/png"
	"math"
	"os"
)

const (
	Width  = 640
	Height = 480
)

var (
	library *ft.Library
	face    *ft.Face

	slot   *ft.GlyphSlot
	matrix *ft.Matrix = ft.NewMatrix()
	pen    *ft.Vector = ft.NewVector()
	err    error

	font     string
	text     string
	fileName string

	angle        float64
	targetHeight int
	n, numChars  int

	img *image.Gray
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("usage: %s font sample-text outpu-file\n", os.Args[0])
		return
	}

	font = os.Args[1]
	text = os.Args[2]
	fileName = os.Args[3]
	numChars = len(text)
	angle = (25.0 / 360) * 3.14159 * 2
	targetHeight = Height

	library, err = ft.InitFreeType()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer library.Done()

	face, err = ft.NewFace(library, font, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer face.Done()

	// use 50pt at 100dpi
	err = face.SetCharSize(50*64, 0, 100, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	slot = face.Glyph()

	// set up matrix
	matrix.SetXX(int64(math.Cos(angle) * 0x10000))
	matrix.SetXY(int64(-math.Sin(angle) * 0x10000))
	matrix.SetYX(int64(math.Sin(angle) * 0x10000))
	matrix.SetYY(int64(math.Cos(angle) * 0x10000))

	pen.SetX(300 * 64)
	pen.SetY(int64((targetHeight - 200) * 64))

	img = image.NewGray(image.Rect(0, 0, Width, Height))

	var x, y int = 0, 50

	for n = 0; n < numChars; n++ {
		// set transformation
		face.SetTransform(matrix, pen)

		// load glyph image into slot (erase previous one)
		err = face.LoadChar(uint64(text[n]), ft.LoadRender)
		if err != nil {
			continue
		}

		bitmap := slot.Bitmap()
		drawBitmap(bitmap, x, Height-y)

		advance := slot.Advance()
		pen.SetX(pen.X() + advance.X())
		pen.SetY(pen.Y() + advance.Y())
		x += int(bitmap.Width())
		y += int(bitmap.Rows())
	}

	err = saveImage(img, fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func saveImage(img image.Image, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, img)
}

func drawBitmap(bitmap *ft.Bitmap, x, y int) error {
	var (
		xMax int = x + bitmap.Width()
		yMax int = y + bitmap.Rows()
		r        = image.Rect(x, y, xMax, yMax)
		sp       = image.Pt(0, 0)
	)
	src, err := bitmap.GrayImage()
	if err != nil {
		return err
	}
	fmt.Printf("src.Bounds: %v\n", src.Bounds())
	fmt.Printf("r: %v\n", r)
	draw.Draw(img, r, src, sp, draw.Src)
	return nil
}
