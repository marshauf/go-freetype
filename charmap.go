package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

// A handle to a given character map.
// A charmap is used to translate character codes in a given encoding into glyph indexes for its parent's face.
// Some font formats may provide several charmaps per font.
type CharMap struct {
	handle C.FT_CharMap
}

// A handle to the parent face object.
func (c *CharMap) Face() Face {
	return Face{handle: c.handle.face}
}

// An Encoding tag identifying the charmap.
// Use this with Face.SelectCharmap.
func (c *CharMap) Encoding() uint64 {
	return uint64(c.handle.encoding)
}

// An ID number describing the platform for the following encoding ID.
// This comes directly from the TrueType specification and should be emulated for other formats.
func (c *CharMap) PlatformID() uint16 {
	return uint16(c.handle.platform_id)
}

// A platform specific encoding number.
// This also comes from the TrueType specification and should be emulated similarly.
func (c *CharMap) EncodingID() uint16 {
	return uint16(c.handle.encoding_id)
}
