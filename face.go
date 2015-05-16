package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

import (
	"reflect"
	"unsafe"
)

type Face struct {
	handle C.FT_Face
}

// NewFace opens a font by its filePathName.
func NewFace(library *Library, filePathName string, faceIndex int64) (*Face, error) {
	face := &Face{}
	cfilePathName := C.CString(filePathName)
	defer C.free(unsafe.Pointer(cfilePathName))
	errno := C.FT_New_Face(library.handle, cfilePathName, C.FT_Long(faceIndex), &face.handle)
	if errno != 0 {
		return nil, GetError(errno)
	}
	return face, nil
}

// TODO OpenFace

func NewMemoryFace(library *Library, data []byte, faceIndex int64) (*Face, error) {
	face := &Face{}
	buffer := (*C.FT_Byte)(unsafe.Pointer(&data[0]))
	errno := C.FT_New_Memory_Face(library.handle, buffer, C.FT_Long(len(data)), C.FT_Long(faceIndex), &face.handle)
	if errno != 0 {
		return nil, GetError(errno)
	}
	return face, nil
}

func (f *Face) Done() error {
	errno := C.FT_Done_Face(f.handle)
	return GetError(errno)
}

func (f *Face) NumFaces() int64 {
	return int64(f.handle.num_faces)
}

func (f *Face) FaceIndex() int64 {
	return int64(f.handle.face_index)
}

func (f *Face) FaceFlags() int64 {
	return int64(f.handle.face_flags)
}

func (f *Face) StyleFlags() int64 {
	return int64(f.handle.style_flags)
}

func (f *Face) FamilyName() string {
	return C.GoString((*C.char)(unsafe.Pointer(f.handle.family_name)))
}

func (f *Face) StyleName() string {
	return C.GoString((*C.char)(unsafe.Pointer(f.handle.style_name)))
}

func (f *Face) NumFixedSizes() int {
	return int(f.handle.num_fixed_sizes)
}

func (f *Face) AvailableSizes() []BitmapSize {
	l := f.NumFixedSizes()
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(f.handle.available_sizes)),
		Len:  l,
		Cap:  l,
	}
	sizes := make([]BitmapSize, l)
	goSlice := *(*[]C.FT_Bitmap_Size)(unsafe.Pointer(&hdr))
	for i := range goSlice {
		sizes[i] = BitmapSize{goSlice[i]}
	}
	return sizes
}

func (f *Face) NumCharmaps() int {
	return int(f.handle.num_charmaps)
}

func (f *Face) CharMaps() []CharMap {
	l := f.NumCharmaps()
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(f.handle.charmaps)),
		Len:  l,
		Cap:  l,
	}
	charmaps := make([]CharMap, l)
	goSlice := *(*[]C.FT_CharMap)(unsafe.Pointer(&hdr))
	for i := range goSlice {
		charmaps[i] = CharMap{goSlice[i]}
	}
	return charmaps
}

// The following member variables (down to `underline_thickness')
// are only relevant to scalable outlines; cf. @FT_Bitmap_Size
// for bitmap fonts.

func (f *Face) BBox() BBox {
	return BBox{f.handle.bbox}
}

func (f *Face) Ascender() int16 {
	return int16(f.handle.ascender)
}

func (f *Face) Descender() int16 {
	return int16(f.handle.descender)
}

func (f *Face) Height() int16 {
	return int16(f.handle.height)
}

func (f *Face) MaxAdvanceHeight() int16 {
	return int16(f.handle.max_advance_height)
}

func (f *Face) MaxAdvanceWidth() int16 {
	return int16(f.handle.max_advance_width)
}

func (f *Face) UnderlinePosition() int16 {
	return int16(f.handle.underline_position)
}

func (f *Face) UnderlineThickness() int16 {
	return int16(f.handle.underline_thickness)
}

func (f *Face) Glyph() *GlyphSlot {
	return &GlyphSlot{handle: f.handle.glyph}
}

func (f *Face) Size() Size {
	return Size{f.handle.size}
}

func (f *Face) CharMap() CharMap {
	return CharMap{f.handle.charmap}
}

// functions

// GetCharIndex returns the glyph index of a given character code.
// This function uses a charmap object to do the mapping.
func (f *Face) GetCharIndex(char uint64) uint {
	index := C.FT_Get_Char_Index(f.handle, C.FT_ULong(char))
	return uint(index)
}

func (f *Face) LoadChar(char uint64, flags int) error {
	errno := C.FT_Load_Char(f.handle, C.FT_ULong(char), C.FT_Int32(flags))
	if errno != 0 {
		return GetError(errno)
	}
	return nil
}

func (f *Face) LoadGlyph(index uint, flags int) error {
	errno := C.FT_Load_Glyph(f.handle, C.FT_UInt(index), C.FT_Int32(flags))
	if errno != 0 {
		return GetError(errno)
	}
	return nil
}

func (f *Face) SelectCharmap(encoding int) error {
	errno := C.FT_Select_Charmap(f.handle, C.FT_Encoding(encoding))
	if errno != 0 {
		return GetError(errno)
	}
	return nil
}

func (f *Face) SetCharSize(charWidth, charHeight int64, horzResolution, vertResolution uint) error {
	errno := C.FT_Set_Char_Size(f.handle, C.FT_F26Dot6(charWidth), C.FT_F26Dot6(charHeight), C.FT_UInt(horzResolution), C.FT_UInt(vertResolution))
	if errno != 0 {
		return GetError(errno)
	}
	return nil
}

func (f *Face) SetPixelSizes(width, height uint) error {
	errno := C.FT_Set_Pixel_Sizes(f.handle, C.FT_UInt(width), C.FT_UInt(height))
	if errno != 0 {
		return GetError(errno)
	}
	return nil
}

func (f *Face) SetTransform(matrix *Matrix, delta *Vector) {
	C.FT_Set_Transform(f.handle, &matrix.handle, &delta.handle)
}
