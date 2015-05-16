package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

type GlyphMetrics struct {
	handle C.FT_Glyph_Metrics
}

// The glyph's width.
func (m *GlyphMetrics) Width() int64 {
	return int64(m.handle.width)
}

// The glyph's height.
func (m *GlyphMetrics) Height() int64 {
	return int64(m.handle.height)
}

// Left side bearing for horizontal layout.
func (m *GlyphMetrics) HoriBearingX() int64 {
	return int64(m.handle.horiBearingX)
}

// Top side bearing for horizontal layout.
func (m *GlyphMetrics) HoriBearingY() int64 {
	return int64(m.handle.horiBearingY)
}

// Advance width for horizontal layout.
func (m *GlyphMetrics) HoriAdvance() int64 {
	return int64(m.handle.horiAdvance)
}

// Left side bearing for vertical layout.
func (m *GlyphMetrics) VertBearingX() int64 {
	return int64(m.handle.vertBearingX)
}

// Top side bearing for vertical layout.
// Larger positive values mean further below the vertical glyph origin.
func (m *GlyphMetrics) VertBearingY() int64 {
	return int64(m.handle.vertBearingY)
}

// Advance height for vertical layout.
// Positive values mean the glyph has a positive advance downward.
func (m *GlyphMetrics) VertAdvance() int64 {
	return int64(m.handle.vertAdvance)
}
