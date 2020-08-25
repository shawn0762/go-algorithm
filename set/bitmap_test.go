package set

import "testing"

func TestBitMap32_PrintMap(t *testing.T) {
	bmp := NewBitMap(64)
	bmp.Add(5)
	bmp.Add(7)
	bmp.Add(15)
	bmp.PrintMap()
}

func TestBitMap32_Has(t *testing.T) {
	bmp := NewBitMap(64)
	bmp.Add(5)
	bmp.Add(7)
	bmp.Add(15)
	if bmp.Has(7) == false {
		t.Errorf("num %d exists", 7)
	}
	if bmp.Has(3) == true {
		t.Errorf("num %d not exists", 3)
	}
}
