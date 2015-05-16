#include "_cgo_export.h"

void ftRasterParamsGraySpans(int y, int count, const FT_Span* spans, void* user) {
	goRasterParamsGraySpans(y, count,(FT_Span*)spans, user);
}
void ftRasterParamsGraySpansCB(FT_Raster_Params* params) {
	params->gray_spans = ftRasterParamsGraySpans;
}
