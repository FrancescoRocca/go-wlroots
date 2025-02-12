package wlroots

/*
#cgo pkg-config: wlroots wayland-server
#cgo CFLAGS: -DWLR_USE_UNSTABLE
#include <wlr/types/wlr_layer_shell_v1.h>
*/
import "C"

type LayerShellV1 struct {
	p *C.struct_wlr_layer_shell_v1
}

type LayerSurfaceV1 struct {
	p *C.struct_wlr_layer_surface_v1
}

// wlr_layer_shell_v1_create
func LayerShellV1Create(display *C.struct_wl_display, version uint32) *LayerShellV1 {
	return &LayerShellV1{
		p: C.wlr_layer_shell_v1_create(display, C.uint32_t(version)),
	}
}

// wlr_layer_surface_v1_configure
func (surface *LayerSurfaceV1) Configure(width, height uint32) uint32 {
	return uint32(C.wlr_layer_surface_v1_configure(surface.p, C.uint32_t(width), C.uint32_t(height)))
}

// wlr_layer_surface_v1_destroy
func (surface *LayerSurfaceV1) Destroy() {
	C.wlr_layer_surface_v1_destroy(surface.p)
}

// wlr_layer_surface_v1_try_from_wlr_surface
func LayerSurfaceV1TryFromWlrSurface(surface *C.struct_wlr_surface) *LayerSurfaceV1 {
	p := C.wlr_layer_surface_v1_try_from_wlr_surface(surface)
	if p == nil {
		return nil
	}
	return &LayerSurfaceV1{p: p}
}

// wlr_layer_surface_v1_from_resource
func LayerSurfaceV1FromResource(resource *C.struct_wl_resource) *LayerSurfaceV1 {
	return &LayerSurfaceV1{
		p: C.wlr_layer_surface_v1_from_resource(resource),
	}
}

// wlr_layer_surface_v1_get_exclusive_edge
/*func (surface *LayerSurfaceV1) GetExclusiveEdge() Edges {
	return Edges(C.wlr_layer_surface_v1_get_exclusive_edge(surface.p))
}*/
