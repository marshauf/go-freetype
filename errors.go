package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

import (
	"errors"
)

func GetError(errno C.FT_Error) error {
	switch errno {
	case 0x00:
		return nil
	case 0x01:
		return ErrCanNotOpenResource
	case 0x02:
		return ErrUnknownFileFormat
	case 0x03:
		return ErrInvalidFileFormat
	case 0x04:
		return ErrInvalidVersion
	case 0x05:
		return ErrLowerModuleVersion
	case 0x06:
		return ErrInvalidArgument
	case 0x07:
		return ErrUnimplementedFeature
	case 0x08:
		return ErrInvalidTable
	case 0x09:
		return ErrInvalidOffset
	case 0x0A:
		return ErrArrayTooLarge
	case 0x0B:
		return ErrMissingModule
	case 0x0C:
		return ErrMissingProperty
	case 0x10:
		return ErrInvalidGlyphIndex
	case 0x11:
		return ErrInvalidCharacterCode
	case 0x12:
		return ErrInvalidGlyphFormat
	case 0x13:
		return ErrCannotRenderGlyph
	case 0x14:
		return ErrInvalidOutline
	case 0x15:
		return ErrInvalidComposite
	case 0x16:
		return ErrTooManyHints
	case 0x17:
		return ErrInvalidPixelSize
	case 0x20:
		return ErrInvalidHandle
	case 0x21:
		return ErrInvalidLibraryHandle
	case 0x22:
		return ErrInvalidDriverHandle
	case 0x23:
		return ErrInvalidFaceHandle
	case 0x24:
		return ErrInvalidSizeHandle
	case 0x25:
		return ErrInvalidSlotHandle
	case 0x26:
		return ErrInvalidCharMapHandle
	case 0x27:
		return ErrCacheHandle
	case 0x28:
		return ErrInvalidStreamHandle
	}
	return errors.New("freetype: Unknown error")
}

var (
	// wrapper errors
	ErrUnsupportedPixelMode = errors.New("freetype: unsupported pixel mode")

	// glyph/character errors
	ErrCanNotOpenResource   = errors.New("freetype: cannot open resource")
	ErrUnknownFileFormat    = errors.New("freetype: unknown file format")
	ErrInvalidFileFormat    = errors.New("freetype: broken file")
	ErrInvalidVersion       = errors.New("freetype: invalid FreeType version")
	ErrLowerModuleVersion   = errors.New("freetype: module version is too low")
	ErrInvalidArgument      = errors.New("freetype: invalid argument")
	ErrUnimplementedFeature = errors.New("freetype: unimplemented feature")
	ErrInvalidTable         = errors.New("freetype: broken table")
	ErrInvalidOffset        = errors.New("freetype: broken offset within table")
	ErrArrayTooLarge        = errors.New("freetype: array allocation size too large")
	ErrMissingModule        = errors.New("freetype: missing module")
	ErrMissingProperty      = errors.New("freetype: missing property")

	// glyph/character errors
	ErrInvalidGlyphIndex    = errors.New("freetype: invalid glyph index")
	ErrInvalidCharacterCode = errors.New("freetype: invalid character code")
	ErrInvalidGlyphFormat   = errors.New("freetype: unsupported glyph image format")
	ErrCannotRenderGlyph    = errors.New("freetype: cannot render this glyph format")
	ErrInvalidOutline       = errors.New("freetype: invalid outline")
	ErrInvalidComposite     = errors.New("freetype: invalid composite glyph")
	ErrTooManyHints         = errors.New("freetype: too many hints")
	ErrInvalidPixelSize     = errors.New("freetype: invalid pixel size")

	// handle errors
	ErrInvalidHandle        = errors.New("freetype: invalid object handle")
	ErrInvalidLibraryHandle = errors.New("freetype: invalid library handle")
	ErrInvalidDriverHandle  = errors.New("freetype: invalid module handle")
	ErrInvalidFaceHandle    = errors.New("freetype: invalid face handle")
	ErrInvalidSizeHandle    = errors.New("freetype: invalid size handle")
	ErrInvalidSlotHandle    = errors.New("freetype: invalid glyph slot handle")
	ErrInvalidCharMapHandle = errors.New("freetype: invalid charmap handle")
	ErrCacheHandle          = errors.New("freetype: invalid cache manager handle")
	ErrInvalidStreamHandle  = errors.New("freetype: invalid stream handle")
)
