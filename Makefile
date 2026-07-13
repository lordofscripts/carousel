# - Host OS Platform
ifeq ($(OS),Windows_NT) 
    detected_OS := Windows
else
    detected_OS := $(sh -c 'uname 2>/dev/null || echo Unknown')
endif

# - Go Build Environment
GO=go
GO_TAGS=-tags logx
ifeq ($(detected_OS), Windows)
	GOFLAGS = -v -buildmode=exe -gcflags all=-N 
	EXE_EXT=.exe
else
	GOFLAGS = -v -buildmode=pie
	EXE_EXT=
endif

# - Source Project Environment
# get the Makefile's directory (GNU Make >= v3.81)
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(dir $(mkfile_path))
# set the GO Project's BIN directory
GO_PROJ_BIN=${mkfile_dir}bin

# - Application Stanza
BIN_OUT=$(GOBIN)/goCarousel$(EXE_EXT)
BIN_OUT_UTIL=$(GOBIN)/goUnixStyle$(EXE_EXT)
MAIN=cmd/*.go

# - Packagers only
PKG_NAME=goCarousel
PKG_VERSION=1.1.0
PKG_REVISION=1
PKG_ARCH=amd64
PKG_SEMANTIC_VERSION=$(shell sed -n '/>>>BEGIN/,/>>>END/p' version.go > /tmp/mainversion_gc.go && go run /tmp/mainversion_gc.go short)
PKG_FULL_VERSION:= $(patsubst v%,%,$(PKG_SEMANTIC_VERSION))
PKG_PUBLIC_NAME=$(shell grep -m 1 'module' go.mod | sed -E 's/^module\s+//p')
PKG_FULLNAME=${PKG_NAME}_${PKG_VERSION}-${PKG_REVISION}_${PKG_ARCH}
PKG_BUILD_DIR=${HOME}/Develop/Distrib/Build/${PKG_NAME}
PKG_PPA_DIR=${HOME}/Develop/Distrib/PPA

# - Application Targets
.PHONY: clean build

build:
	$(GO) build $(GO_TAGS) $(GOFLAGS) -o ${BIN_OUT} ${MAIN}

buildwin:
	$(GO) build $(GO_TAGS) $(GOFLAGS) -o ${BIN_OUT}.exe ${MAIN}

release:
	$(GO) build $(GOFLAGS) -o ${BIN_OUT} ${MAIN}

version:
	@echo $(PKG_FULL_VERSION)

clean:
	go clean

util:
	$(GO) build $(GOFLAGS) -o ${BIN_OUT_UTIL} cmd/util/*go

run:
	go run -race  $MAIN

proxy:
	GOPROXY=proxy.golang.org go list -m $(PKG_PUBLIC_NAME)@v$(PKG_FULL_VERSION)

lint: 
	@gofmt -l . | grep ".*\.go"

test:
	go test tests/*test.go	

debian:
	rm -fR ${PKG_BUILD_DIR}
	mkdir -p ${PKG_BUILD_DIR}/DEBIAN
	#ln -s ${PKG_BUILD_DIR}/DEBIAN ${PKG_BUILD_DIR}/debian
	cp -R distrib/DEBIAN/* ${PKG_BUILD_DIR}/DEBIAN
	mkdir -p ${PKG_BUILD_DIR}/usr/bin
	mkdir -p ${PKG_BUILD_DIR}/usr/share/doc/${PKG_NAME}/assets
	mkdir -p ${PKG_BUILD_DIR}/usr/share/man/man1
	gzip -n -9 -c distrib/manpages/man1/goCarousel.1 > ${PKG_BUILD_DIR}/usr/share/man/man1/goCarousel.1.gz
	gzip -n -9 -c distrib/manpages/man1/goUnixStyle.1 > ${PKG_BUILD_DIR}/usr/share/man/man1/goUnixStyle.1.gz
	mkdir -p ${PKG_BUILD_DIR}/usr/share/man/man5
	gzip -n -9 -c distrib/manpages/man5/goCarousel.5 > ${PKG_BUILD_DIR}/usr/share/man/man1/goCarousel.5.gz
	strip --strip-unneeded ${BIN_OUT}
	cp ${BIN_OUT} ${PKG_BUILD_DIR}/usr/bin
	strip --strip-unneeded ${BIN_OUT_UTIL}
	cp ${BIN_OUT_UTIL} ${PKG_BUILD_DIR}/usr/bin
	cp distrib/DEBIAN/copyright ${PKG_BUILD_DIR}/usr/share/doc/${PKG_NAME}
	cp docs/README.md ${PKG_BUILD_DIR}/usr/share/doc/${PKG_NAME}
	cp docs/assets/* ${PKG_BUILD_DIR}/usr/share/doc/${PKG_NAME}/assets
	gzip -n -9 -c distrib/DEBIAN/changelog > ${PKG_BUILD_DIR}/usr/share/doc/${PKG_NAME}/changelog.gz
	(cd ${PKG_BUILD_DIR} && dpkg-deb --root-owner-group -b ./ ${PKG_FULLNAME}.deb)
	#(cd ${PKG_BUILD_DIR} && fakeroot /usr/bin/dpkg-buildpackage --build=binary -us -uc -b ./ ${PKG_FULLNAME})
	#@mv /tmp/${PKG_FULLNAME}.deb ${DEST_REPOSITORY}
