package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

// A structure used to hold an outline's bounding box, i.e.,
// the coordinates of its extrema in the horizontal and vertical directions.
type BBox struct {
	handle C.FT_BBox
}

// The horizontal minimum (left-most).
func (b *BBox) MinX() int64 {
	return int64(b.handle.xMin)
}

// The vertical minimum (bottom-most).
func (b *BBox) MinY() int64 {
	return int64(b.handle.yMin)
}

// The horizontal maximum (right-most).
func (b *BBox) MaxX() int64 {
	return int64(b.handle.xMax)
}

// The vertical maximum (top-most).
func (b *BBox) MaxY() int64 {
	return int64(b.handle.yMax)
}
