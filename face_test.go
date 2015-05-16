package freetype

import (
	"image"
	"io/ioutil"
	"testing"
	//"image/png"
	//"os"
)

func TestFaceNil(t *testing.T) {
	face, err := NewFace(lib, "", 0)
	if err != ErrCanNotOpenResource {
		t.Errorf("err should be ErrCanNotOpenResource: %s", err)
	}
	if face != nil {
		t.Error("face should be nill")
	}
}

func TestFace(t *testing.T) {
	face, err := NewFace(lib, fileName, 0)
	if err != nil {
		t.Fatal(err)
		return
	}
	if face == nil {
		t.Fatal("face should not be nill")
	}

	err = face.SetCharSize(0, 16*64, 300, 300)
	if err != nil {
		t.Error(err)
	}

	err = face.SetPixelSizes(0, 16)
	if err != nil {
		t.Error(err)
	}

	index := face.GetCharIndex('A')
	if index != uint(36) {
		t.Error("GetCharIndex('A') should return an uint 36")
	}

	err = face.LoadGlyph(index, LoadDefault)
	if err != nil {
		t.Error(err)
	}

	err = face.Glyph().Render(RenderModeNormal)
	if err != nil {
		t.Error(err)
	}

	// From memory
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
	memoryFace, err := NewMemoryFace(lib, data, 0)
	if err != nil {
		t.Fatal(err)
	}

	err = memoryFace.SetCharSize(0, 16*64, 300, 300)
	if err != nil {
		t.Error(err)
	}

	err = memoryFace.SetPixelSizes(0, 16)
	if err != nil {
		t.Error(err)
	}

	memoryIndex := memoryFace.GetCharIndex('A')
	if memoryIndex != index {
		t.Errorf("the char index of A should be the same nomatter the loading function. FromFile: %v FromMemory: %v", index, memoryIndex)
	}

	err = memoryFace.LoadGlyph(memoryIndex, LoadDefault)
	if err != nil {
		t.Error(err)
	}

	err = memoryFace.Glyph().Render(RenderModeNormal)
	if err != nil {
		t.Error(err)
	}

	// free memory
	err = face.Done()
	if err != nil {
		t.Error(err)
	}

	err = memoryFace.Done()
	if err != nil {
		t.Error(err)
	}
}

func TestTutorial1Refined(t *testing.T) {
	face, err := NewFace(lib, fileName, 0)
	if err != nil {
		t.Fatal(err)
	}
	if face == nil {
		t.Fatal("face should not be nil")
	}

	err = face.SetCharSize(0, 16*64, 72, 72)
	if err != nil {
		t.Error(err)
	}

	slot := face.Glyph()
	var (
		penX = 16
		penY = 16
		n    int
		text = "Hello, World" // "Hello, 世界" doesn't render correctly
		// because the last the chars are under East Asian Scripts, CJK Unified Ideographs
		// the charmap picked by default is the first one which is European Alphabets, Basic Latin

		img = image.NewGray(image.Rect(0, 0, 256, 256))
	)

	for n = 0; n < len(text); n++ {
		err = face.LoadChar(uint64(text[n]), LoadRender)
		if err != nil {
			t.Error(err)
		}

		drawBitmap(img, slot.Bitmap(), penX, penY, slot.BitmapLeft(), slot.BitmapTop())

		penX += int(slot.Advance().X()) >> 6
	}
	/*
		file, err := os.Create("test.png")
		c.Assert(err, IsNil)
		err = png.Encode(file, img)
		c.Assert(err, IsNil)
	*/

	err = face.Done()
	if err != nil {
		t.Error(err)
	}
}

func drawBitmap(img *image.Gray, bitmap *Bitmap, x, y, left, top int) error {
	b, err := bitmap.GrayImage()
	if err != nil {
		return err
	}

	rec := b.Bounds()
	for i := 0; i < rec.Dx(); i++ {
		for j := 0; j < rec.Dy(); j++ {
			img.Set(x+i, y+j-top, b.At(i, j))
		}
	}
	return nil
}
