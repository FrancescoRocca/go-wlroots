package wlroots

/*
#cgo pkg-config: wlroots wayland-server
#cgo CFLAGS: -DWLR_USE_UNSTABLE
#include <wlr/types/wlr_layer_shell_v1.h>
#include <wayland-server-core.h>
*/
import "C"

type Resource struct {
	p *C.struct_wl_resource
}

type LayerShellV1 struct {
	p *C.struct_wlr_layer_shell_v1
}

type LayerSurfaceV1 struct {
	p *C.struct_wlr_layer_surface_v1
}

func LayerShellV1Create(display Display, version uint32) *LayerShellV1 {
	return &LayerShellV1{
		p: C.wlr_layer_shell_v1_create(display.p, C.uint32_t(version)),
	}
}

func (surface *LayerSurfaceV1) Configure(width, height uint32) uint32 {
	return uint32(C.wlr_layer_surface_v1_configure(surface.p, C.uint32_t(width), C.uint32_t(height)))
}

func (surface *LayerSurfaceV1) Destroy() {
	C.wlr_layer_surface_v1_destroy(surface.p)
}

func LayerSurfaceV1TryFromWlrSurface(surface *Surface) *LayerSurfaceV1 {
	p := C.wlr_layer_surface_v1_try_from_wlr_surface(surface.p)
	if p == nil {
		return nil
	}
	return &LayerSurfaceV1{p: p}
}

func LayerSurfaceV1FromResource(resource *Resource) *LayerSurfaceV1 {
	return &LayerSurfaceV1{
		p: C.wlr_layer_surface_v1_from_resource(resource.p),
	}
}
