package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
void ftRasterParamsGraySpansCB(FT_Raster_Params* params);
*/
import "C"

import (
	"reflect"
	"unsafe"
)

const (
	RasterFlagDefault int = C.FT_RASTER_FLAG_DEFAULT
	RasterFlagAA      int = C.FT_RASTER_FLAG_AA
	RasterFlagDirect  int = C.FT_RASTER_FLAG_DIRECT
	RasterFlagClip    int = C.FT_RASTER_FLAG_CLIP
)

type RasterParams struct {
	handle       C.FT_Raster_Params
	graySpanFunc SpanFunc
}

func NewRasterParams() *RasterParams {
	return &RasterParams{handle: C.FT_Raster_Params{}}
}

// The target bitmap.
func (params *RasterParams) Target() *Bitmap {
	return &Bitmap{*params.handle.target}
}

func (params *RasterParams) SetTarget(target *Bitmap) {
	params.handle.target = &target.handle
}

func (params *RasterParams) Source() unsafe.Pointer {
	return params.handle.source
}

func (params *RasterParams) SetSource(source *interface{}) {
	params.handle.source = unsafe.Pointer(source)
}

// The rendering flags.
func (params *RasterParams) Flags() int {
	return int(params.handle.flags)
}

func (params *RasterParams) SetFlags(flags int) {
	params.handle.flags = C.int(flags)
}

func (params *RasterParams) GraySpans() SpanFunc {
	return params.graySpanFunc
}

func (params *RasterParams) SetGraySpans(f SpanFunc) {
	params.handle.user = unsafe.Pointer(&params.handle)
	rasterParams[&params.handle] = params
	params.graySpanFunc = f
	C.ftRasterParamsGraySpansCB(&params.handle)
}

var rasterParams map[*C.FT_Raster_Params]*RasterParams = make(map[*C.FT_Raster_Params]*RasterParams)

//export goRasterParamsGraySpans
func goRasterParamsGraySpans(y, count C.FT_Int, spans *C.FT_Span, user unsafe.Pointer) {
	c := int(count)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(spans)),
		Len:  c,
		Cap:  c,
	}
	goSlice := *(*[]C.FT_Span)(unsafe.Pointer(&hdr))
	s := make([]Span, c)
	for i := range goSlice {
		s[i] = Span{goSlice[i]}
	}
	params := rasterParams[(*C.FT_Raster_Params)(user)]
	params.graySpanFunc(int(y), c, s)
}

// An optional clipping box. It is only used in direct rendering mode.
// Note that coordinates here should be expressed in integer pixels
// (and not in 26.6 fixed-point units)
func (params *RasterParams) ClipBox() *BBox {
	return &BBox{params.handle.clip_box}
}

func (params *RasterParams) SetClipBox(bbox *BBox) {
	params.handle.clip_box = bbox.handle
}
