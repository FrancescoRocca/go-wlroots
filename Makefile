WAYLAND_PROTOCOLS=/usr/share/wlr-protocols

all: tinywl

tinywl: prep layer-shell-protocol
	go build -o build/bin/tinywl github.com/swaywm/go-wlroots/cmd/tinywl

layer-shell-protocol:
	wayland-scanner client-header \
		$(WAYLAND_PROTOCOLS)/unstable/wlr-layer-shell-unstable-v1.xml \
		wlroots/wlr-layer-shell-unstable-v1-protocol.h
	wayland-scanner private-code \
		$(WAYLAND_PROTOCOLS)/unstable/wlr-layer-shell-unstable-v1.xml \
		wlroots/wlr-layer-shell-unstable-v1-protocol.c

prep:
	mkdir -p build/bin

clean:
	rm -rf build wlroots/xdg-shell-protocol.c wlroots/xdg-shell-protocol.h