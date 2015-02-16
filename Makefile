
ASSET_DIRS=assets/js assets/js/external assets/js/vendor assets/css assets/images assets/templates assets/fonts assets/translations

ARCH=$(shell uname -m)
OS=$(shell uname -s)
VER=$(shell uname -r)
BINVER=$(shell cat gosnatch/config.go | grep "VERSION" | awk '{print $$3}')



GOFLAGS=

# external sqlite libs on i386 osx
ifeq ($(OS),Darwin)
	ifeq ($(ARCH),i386)
		GOFLAGS+=-ldflags -linkmode=external
	endif
endif

default: clean depends build

depends:
	go get $(GOFLAGS)
	go get github.com/jteeuwen/go-bindata && go install github.com/jteeuwen/go-bindata/go-bindata
	go get -u github.com/nicksnyder/go-i18n/goi18n

assets: $(ASSET_DIRS)
	go-bindata -pkg gosnatch -o gosnatch/bindata.go $(ASSET_DIRS)

assets-debug: $(ASSET_DIRS)
	go-bindata -debug -pkg gosnatch -o gosnatch/bindata.go $(ASSET_DIRS)

build: assets
	go build $(GOFLAGS) -o dist/gosnatch-$(OS)-$(ARCH) -v

dist: clean depends build
	cd dist && tar cvzf gosnatch-$(OS)-$(ARCH).tgz gosnatch-$(OS)-$(ARCH)

install: assets
	go install $(GOFLAGS)


update: clean depends
	go build $(GOFLAGS) -o dist/$(OS)-$(ARCH)
	go-selfupdate ./dist/ $(BINVER)

push-update:
	go-selfupdate ./dist/ $(BINVER)

cross:
	goxc -pv=$(BINVER) -pr="alpha" -bu=1 -tasks-="validate" -d=./dist -o="{{.Dest}}{{.PS}}{{.Version}}{{.PS}}{{.Os}}-{{.Arch}}" -build-ldflags="-linkmode=external" -os="linux" -arch="386"

run: assets-debug
	go run $(GOFLAGS) main.go

debug: assets-debug
	export SNT_DEVEL=true && export SNT_DEBUG=true && go run $(GOFLAGS) main.go

trans:
	cd ./assets/translations && goi18n *.all.json

trans-update:
	cd ./assets/translations && goi18n *.all.json *.untranslated.json

clean:
	rm -f dist/*