package wlroots

// #cgo pkg-config: wlroots wayland-server
// #cgo CFLAGS: -DWLR_USE_UNSTABLE
// #include <wlr/render/wlr_texture.h>
import "C"
import (
	"image"
	"unsafe"

	"golang.org/x/image/draw"
)

type Texture struct {
	p *C.struct_wlr_texture
}

func (t Texture) Destroy() {
	C.wlr_texture_destroy(t.p)
}

func (t Texture) Nil() bool {
	return t.p == nil
}

func TextureFromPixels(renderer *Renderer, fmt uint32, stride uint32, width uint32, height uint32, data []byte) Texture {
	cFmt := C.uint32_t(fmt)
	cStride := C.uint32_t(stride)
	cWidth := C.uint32_t(width)
	cHeight := C.uint32_t(height)
	cData := unsafe.Pointer(&data[0])

	texture := C.wlr_texture_from_pixels(
		renderer.p,
		cFmt,
		cStride,
		cWidth,
		cHeight,
		cData,
	)

	if texture == nil {
		return Texture{p: nil}
	}

	return Texture{p: texture}
}

// wlr fork
func TextureFromImage(renderer *Renderer, img image.Image) Texture {
	nrgba, err := img.(*image.NRGBA)
	if err {
		nrgba = image.NewNRGBA(img.Bounds())
		draw.Copy(nrgba, image.ZP, img, nrgba.Bounds(), draw.Src, nil)
	}

	return TextureFromPixels(
		renderer,
		uint32('A'|('B'<<8)|('2'<<16)|('4'<<24)),
		uint32(nrgba.Stride),
		uint32(nrgba.Bounds().Dx()),
		uint32(nrgba.Bounds().Dy()),
		nrgba.Pix,
	)
}
