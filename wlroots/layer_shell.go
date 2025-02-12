package wlroots

// #cgo pkg-config: wlroots wayland-server wayland-client
// #cgo CFLAGS: -DWLR_USE_UNSTABLE
// #include <wlr/types/wlr_layer_shell_v1.h>
// #include <wayland-server-core.h>
// #include <wayland-client-core.h>
// #include <wlr-layer-shell-unstable-v1-protocol.h>
import "C"
import "unsafe"

type Layer uint32

const (
	LayerBackground Layer = C.ZWLR_LAYER_SHELL_V1_LAYER_BACKGROUND
	LayerBottom     Layer = C.ZWLR_LAYER_SHELL_V1_LAYER_BOTTOM
	LayerTop        Layer = C.ZWLR_LAYER_SHELL_V1_LAYER_TOP
	LayerOverlay    Layer = C.ZWLR_LAYER_SHELL_V1_LAYER_OVERLAY
)

type Anchor uint32

const (
	AnchorTop    Anchor = C.ZWLR_LAYER_SURFACE_V1_ANCHOR_TOP
	AnchorBottom Anchor = C.ZWLR_LAYER_SURFACE_V1_ANCHOR_BOTTOM
	AnchorLeft   Anchor = C.ZWLR_LAYER_SURFACE_V1_ANCHOR_LEFT
	AnchorRight  Anchor = C.ZWLR_LAYER_SURFACE_V1_ANCHOR_RIGHT
)

type Resource struct {
	p *C.struct_wl_resource
}

type LayerShellV1 struct {
	p *C.struct_wlr_layer_shell_v1
}

type LayerSurfaceV1 struct {
	p *C.struct_wlr_layer_surface_v1
	z *C.struct_zwlr_layer_surface_v1
}

func NewLayerShellV1(display Display, version uint32) *LayerShellV1 {
	return &LayerShellV1{
		p: C.wlr_layer_shell_v1_create(display.p, C.uint32_t(version)),
	}
}

func (s *LayerSurfaceV1) Configure(width, height uint32) uint32 {
	return uint32(C.wlr_layer_surface_v1_configure(s.p, C.uint32_t(width), C.uint32_t(height)))
}

func (s *LayerSurfaceV1) Destroy() {
	if s.z != nil {
		C.zwlr_layer_surface_v1_destroy(s.z)
	}
	if s.p != nil {
		C.wlr_layer_surface_v1_destroy(s.p)
	}
}

func TryFromWlrSurface(surface *Surface) *LayerSurfaceV1 {
	p := C.wlr_layer_surface_v1_try_from_wlr_surface(surface.p)
	if p == nil {
		return nil
	}
	return &LayerSurfaceV1{
		p: p,
		z: (*C.struct_zwlr_layer_surface_v1)(unsafe.Pointer(p)),
	}
}

func FromResource(resource *Resource) *LayerSurfaceV1 {
	p := C.wlr_layer_surface_v1_from_resource(resource.p)
	return &LayerSurfaceV1{
		p: p,
		z: (*C.struct_zwlr_layer_surface_v1)(unsafe.Pointer(p)),
	}
}

func (s *LayerSurfaceV1) SetSize(width, height uint32) {
	C.zwlr_layer_surface_v1_set_size(s.z, C.uint32_t(width), C.uint32_t(height))
}

func (s *LayerSurfaceV1) SetAnchor(anchor Anchor) {
	C.zwlr_layer_surface_v1_set_anchor(s.z, C.uint32_t(anchor))
}

func (s *LayerSurfaceV1) SetExclusiveZone(zone int32) {
	C.zwlr_layer_surface_v1_set_exclusive_zone(s.z, C.int32_t(zone))
}

func (s *LayerSurfaceV1) SetMargin(top, right, bottom, left int32) {
	C.zwlr_layer_surface_v1_set_margin(s.z, C.int32_t(top), C.int32_t(right),
		C.int32_t(bottom), C.int32_t(left))
}

func (s *LayerSurfaceV1) SetKeyboardInteractivity(interactive bool) {
	var val C.uint32_t
	if interactive {
		val = 1
	}
	C.zwlr_layer_surface_v1_set_keyboard_interactivity(s.z, val)
}

func (s *LayerSurfaceV1) AckConfigure(serial uint32) {
	C.zwlr_layer_surface_v1_ack_configure(s.z, C.uint32_t(serial))
}

func (s *LayerSurfaceV1) SetLayer(layer Layer) {
	C.zwlr_layer_surface_v1_set_layer(s.z, C.uint32_t(layer))
}

func (s *LayerSurfaceV1) GetSurface() *Surface {
	return &Surface{p: s.p.surface}
}
