package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"
import "unsafe"

func (page *Page) BeginText() error {
	C.HPDF_Page_BeginText(page.page)
	return page.pdf.GetLastError()
}

func (page *Page) EndText() error {
	C.HPDF_Page_EndText(page.page)
	return page.pdf.GetLastError()
}

func (page *Page) SetCharSpace(value float32) error {
	C.HPDF_Page_SetCharSpace(page.page, C.HPDF_REAL(value))
	return page.pdf.GetLastError()
}

func (page *Page) SetWordSpace(value float32) error {
	C.HPDF_Page_SetWordSpace(page.page, C.HPDF_REAL(value))
	return page.pdf.GetLastError()
}

func (page *Page) ShowText(text string) error {
	cText := C.CString(text)

	C.HPDF_Page_ShowText(page.page, cText)

	C.free(unsafe.Pointer(cText))

	return page.pdf.GetLastError()
}

func (page *Page) TextOut(x float32, y float32, text string) error {
	cText := C.CString(text)

	C.HPDF_Page_TextOut(page.page, C.HPDF_REAL(x), C.HPDF_REAL(y), cText)

	C.free(unsafe.Pointer(cText))
	return page.pdf.GetLastError()
}

func (page *Page) MoveTextPos(x float32, y float32) error {
	C.HPDF_Page_MoveTextPos(page.page, C.HPDF_REAL(x), C.HPDF_REAL(y))
	return page.pdf.GetLastError()
}
