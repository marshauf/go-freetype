package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

// This structure models the metrics of a bitmap strike
// (i.e., a set of glyphs for a given point size and resolution) in a bitmap font.
// It is used for the ‘available_sizes’ field of Face.
type BitmapSize struct {
	handle C.FT_Bitmap_Size
}

// The vertical distance, in pixels, between two consecutive baselines. It is always positive.
func (size *BitmapSize) Height() int16 {
	return int16(size.handle.height)
}

// The average width, in pixels, of all glyphs in the strike.
func (size *BitmapSize) Width() int16 {
	return int16(size.handle.width)
}

// The nominal size of the strike in 26.6 fractional points. This field is not very useful.
func (size *BitmapSize) Size() int64 {
	return int64(size.handle.size)
}

// The horizontal ppem (nominal width) in 26.6 fractional pixels.
func (size *BitmapSize) XPpem() int64 {
	return int64(size.handle.x_ppem)
}

// The vertical ppem (nominal height) in 26.6 fractional pixels.
func (size *BitmapSize) YPpem() int64 {
	return int64(size.handle.y_ppem)
}
