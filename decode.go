//Package cifar is a small package that decodes cifar bin files
package cifar

import (
	"image"
	"image/color"
	"io"
)

//Cifar100Image is used for the Cifar 100 image dataset
type Cifar100Image struct {
	cifarimage
	course, fine byte
}

//Cifar10Image is a cifar 10 image
type Cifar10Image struct {
	cifarimage
	num byte
}

//CifarImage is an image that was taken from cifar.bin file
type cifarimage struct {
	data []byte
}

//ColorModel satisfies the image.Image interface
func (cif *cifarimage) ColorModel() color.Model {
	return color.RGBAModel
}

//Bounds satisfies the image.Image interface
func (cif *cifarimage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 32, 32)
}

//At satisfies the image.Image interface
func (cif *cifarimage) At(x, y int) color.Color {

	return color.RGBA{
		R: cif.data[y*32+x],
		G: cif.data[1024+y*32+x],
		B: cif.data[2048+y*32+x],
		A: 255,
	}
}

//RawData passes a slice of if the raw image data. Changes to data will change the underlying structure to cif.
//
//Label byte is excluded from data.
func (cif *cifarimage) RawData() (data []byte) {
	return cif.data
}

//Label returns the byte label
func (cif *Cifar10Image) Label() byte {
	return cif.num
}

//Label returns the course and fine bytes.
func (cif *Cifar100Image) Label() (course, fine byte) {
	return cif.course, cif.fine
}

//Decode10 decoes a cifar bin file
func Decode10(r io.Reader) ([]*Cifar10Image, error) {
	buffer := make([]byte, 3073)
	x := make([]*Cifar10Image, 0)
	for {
		_, err := r.Read(buffer)
		if err == io.EOF {
			return x, nil
		} else if err != nil {
			return x, err
		}

		cim := new(Cifar10Image)
		cim.num = buffer[0]
		cim.data = make([]byte, 3072)
		copy(cim.data, buffer[1:])
		x = append(x, cim)
	}
}

//Decode100 decoes the CIFAR 100
func Decode100(r io.Reader) ([]*Cifar100Image, error) {
	buffer := make([]byte, 3074)
	x := make([]*Cifar100Image, 0)
	for {
		_, err := r.Read(buffer)
		if err == io.EOF {
			return x, nil
		} else if err != nil {
			return x, err
		}

		cim := new(Cifar100Image)
		cim.course = buffer[0]
		cim.fine = buffer[1]
		cim.data = make([]byte, 3072)
		copy(cim.data, buffer[2:])
		x = append(x, cim)
	}

}
