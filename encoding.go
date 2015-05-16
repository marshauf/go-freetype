package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

const (
	EncodingNone     int = C.FT_ENCODING_NONE
	EncodingMSSymbol int = C.FT_ENCODING_MS_SYMBOL
	EncodingUnicode  int = C.FT_ENCODING_UNICODE
	EncodingSjis     int = C.FT_ENCODING_SJIS
	EncodingGb2312   int = C.FT_ENCODING_GB2312
	EncodingBig5     int = C.FT_ENCODING_BIG5
	EncodingWansung  int = C.FT_ENCODING_WANSUNG
	EncodingJohab    int = C.FT_ENCODING_JOHAB

	/* for backwards compatibility */
	EncodingMSSjis    = EncodingSjis
	EncodingMSGb2312  = EncodingGb2312
	EncodingMSBig5    = EncodingBig5
	EncodingMSWansung = EncodingWansung
	EncodingMSJohab   = EncodingJohab

	EncodingAdobeStandard int = C.FT_ENCODING_ADOBE_STANDARD
	EncodingAdobeExpert   int = C.FT_ENCODING_ADOBE_EXPERT
	EncodingAdobeCustom   int = C.FT_ENCODING_ADOBE_CUSTOM
	EncodingAdobeLatin1   int = C.FT_ENCODING_ADOBE_LATIN_1
	EncodingOldLatin2     int = C.FT_ENCODING_OLD_LATIN_2
	EncodingAppleRoman    int = C.FT_ENCODING_APPLE_ROMAN
)
