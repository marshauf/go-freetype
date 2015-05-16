package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

const (
	PixelModeNone  int = C.FT_PIXEL_MODE_NONE
	PixelModeMono  int = C.FT_PIXEL_MODE_MONO
	PixelModeGray  int = C.FT_PIXEL_MODE_GRAY
	PixelModeGray2 int = C.FT_PIXEL_MODE_GRAY2
	PixelModeGray4 int = C.FT_PIXEL_MODE_GRAY4
	PixelModeLCD   int = C.FT_PIXEL_MODE_LCD
	PixelModeLCDV  int = C.FT_PIXEL_MODE_LCD_V
	PixelModeBGRA  int = C.FT_PIXEL_MODE_BGRA
	PixelModeMax   int = C.FT_PIXEL_MODE_MAX
)
