package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

type SizeMetrics struct {
	handle C.FT_Size_Metrics
}

// horizontal pixels per EM
func (size *SizeMetrics) XPpem() int64 {
	return int64(size.handle.x_ppem)
}

// vertical pixels per EM
func (size *SizeMetrics) YPpem() int64 {
	return int64(size.handle.y_ppem)
}

// scaling values used to convert font
func (size *SizeMetrics) XScale() int64 {
	return int64(size.handle.x_scale)
}

// units to 26.6 fractional pixels
func (size *SizeMetrics) YScale() int64 {
	return int64(size.handle.y_scale)
}

// ascender in 26.6 frac. pixels
func (size *SizeMetrics) Ascender() int64 {
	return int64(size.handle.ascender)
}

// descender in 26.6 frac. pixels
func (size *SizeMetrics) Descender() int64 {
	return int64(size.handle.descender)
}

// text height in 26.6 frac. pixels
func (size *SizeMetrics) Height() int64 {
	return int64(size.handle.height)
}

// max horizontal advance, in 26.6 pixels
func (size *SizeMetrics) MaxAdvance() int64 {
	return int64(size.handle.max_advance)
}
