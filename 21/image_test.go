package day21

import "testing"

func TestImage_Enhance(t *testing.T) {
	input := []string{
		"../.# => ##./#../...",
		".#./..#/### => #..#/..../..../#..#",
	}
	enh := ParseEnhancements(input)
	img := DefaultImage()
	img = EnhanceImage(img, enh)
	img = EnhanceImage(img, enh)
	if on := img.Count(); on != 12 {
		t.Error("Expected 12 pixels turned on, got", on)
	}
}
