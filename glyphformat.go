package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

const (
	GlyphFormatNone      int = C.FT_GLYPH_FORMAT_NONE
	GlyphFormatComposite int = C.FT_GLYPH_FORMAT_COMPOSITE
	GlyphFormatOutline   int = C.FT_GLYPH_FORMAT_OUTLINE
	GlyphFormatPlotter   int = C.FT_GLYPH_FORMAT_PLOTTER
)
