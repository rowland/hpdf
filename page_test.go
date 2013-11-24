package hpdf

import (
	. "testing"
)

func TestAddPage(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if page == nil || err != nil {
		t.Fatal(err)
	}
}

func TestInsertPage(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	insertedPage, err := pdf.InsertPage(page)

	if insertedPage == nil || err != nil {
		t.Fatal(err)
	}
}

func TestPageSetWidth(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	width := float32(200)

	err = page.SetWidth(width)

	if err != nil || page.GetWidth() != width {
		t.Fatal(err)
	}
}

func TestPageSetHeight(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	height := float32(200)

	err = page.SetHeight(height)

	if err != nil || page.GetHeight() != height {
		t.Fatal(err)
	}
}

func TestPageSetSize(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	err = page.SetSize(PAGE_SIZE_COMM10, PAGE_LANDSCAPE)

	if err != nil {
		t.Fatal(err)
	}
}

func TestPageSetRotate(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	err = page.SetRotate(90)

	if err != nil {
		t.Fatal(err)
	}
}
