package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
#include FT_GLYPH_H
*/
import "C"

import (
	"unsafe"
)

// A handle to an object used to model an outline glyph image.
// This is a sub-class of Glyph, and a pointer to OutlineGlyphRec.
type OutlineGlyph struct {
	handle C.FT_OutlineGlyph
}

func CastGlyphToOutlineGlyph(glyph *Glyph) *OutlineGlyph {
	var og C.FT_OutlineGlyph = (C.FT_OutlineGlyph)(unsafe.Pointer(glyph.handle))
	return &OutlineGlyph{og}
}

func (g *OutlineGlyph) Root() *Glyph {
	var handle C.FT_Glyph = (C.FT_Glyph)(unsafe.Pointer(&g.handle.root))
	return &Glyph{handle}
}

func (g *OutlineGlyph) Outline() *Outline {
	return &Outline{g.handle.outline}
}
