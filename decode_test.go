package cifar

import (
	"image/png"
	"os"
	"testing"
)

func TestDecode10(t *testing.T) {
	file, err := os.Open("10singleimage.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	cifarimages, err := Decode10(file)
	if err != nil {
		t.Error(err)
	}
	if len(cifarimages) != 1 {
		t.Error("Not One Image")
	}

	image, err := os.Create("10singleimage.png")
	if err != nil {
		t.Fatal(err)
	}
	defer image.Close()
	err = png.Encode(image, cifarimages[0])
	if err != nil {
		t.Error(err)
	}

}

func TestDecode100(t *testing.T) {
	file, err := os.Open("100singleimage.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	cifarimages, err := Decode100(file)
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Fatal(err)
	}
	if len(cifarimages) != 1 {
		t.Error("Not One Image")
	}
	image, err := os.Create("100singleimage.png")
	if err != nil {
		t.Fatal(err)
	}
	defer image.Close()
	err = png.Encode(image, cifarimages[0])
	if err != nil {
		t.Error(err)
	}

}
