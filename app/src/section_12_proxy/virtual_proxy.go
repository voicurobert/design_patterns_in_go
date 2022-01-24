package proxy

import "fmt"

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func (b *Bitmap) Draw() {
	fmt.Println("drawing image", b.filename)
}

func NewBitmap(filename string) *Bitmap {
	fmt.Println("loading image from", filename)
	return &Bitmap{filename: filename}
}

func DrawImage(image Image) {
	fmt.Println("About to draw image")
	image.Draw()
	fmt.Println("done drawing the image")
}

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}

func (l *LazyBitmap) Draw() {
	if l.bitmap == nil {
		l.bitmap = NewBitmap(l.filename)
	}
	l.bitmap.Draw()
}

func VirtualProxyExample() {
	//bmp := NewBitmap("demo.png")
	//DrawImage(bmp)

	bmp := NewLazyBitmap("demo.png")
	DrawImage(bmp)
}
