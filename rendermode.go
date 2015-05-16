package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

const (
	RenderModeNormal int = C.FT_RENDER_MODE_NORMAL
	RenderModeLight  int = C.FT_RENDER_MODE_LIGHT
	RenderModeMono   int = C.FT_RENDER_MODE_MONO
	RenderModeLCD    int = C.FT_RENDER_MODE_LCD
	RenderModeLCDV   int = C.FT_RENDER_MODE_LCD_V
)
