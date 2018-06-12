OS := $(shell uname)
TARGET := ipd2

ifeq ($(OS),Linux)
	TAR_OPTS := --wildcards
endif

all: deps test vet install build

fmt:
	@echo "Formatting all the things..."
	go fmt ./...
	@echo ""

test: geoip-download certs
	@echo "Running tests"
	go test ./...
	@echo ""

vet:
	@echo "Vetting stuff"
	go vet ./...
	@echo ""

deps:
	@echo "Ensuring dependencies are in place"
	dep ensure
	@echo ""

install:
	@echo "Installing the stuffs"
	go install ./...
	@echo ""

build: build_darwin_amd64 \
	build_linux_amd64 \
	build_windows_amd64

build_darwin_%: GOOS := darwin
build_linux_%: GOOS := linux
build_windows_%: GOOS := windows
build_windows_%: EXT := .exe

build_%_amd64: GOARCH := amd64

build_%:
	env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o build/$(TARGET)-${TRAVIS_TAG}-$(GOOS)_$(GOARCH)$(EXT) ./cmd/ipd/main.go

databases := GeoLite2-City GeoLite2-Country

$(databases):
	@echo "Downloading GeoIP databases"
	mkdir -p data
	curl -fsSL -m 30 http://geolite.maxmind.com/download/geoip/database/$@.tar.gz | tar $(TAR_OPTS) --strip-components=1 -C $(PWD)/data -xzf - '*.mmdb'
	test ! -f data/GeoLite2-City.mmdb || mv data/GeoLite2-City.mmdb data/city.mmdb
	test ! -f data/GeoLite2-Country.mmdb || mv data/GeoLite2-Country.mmdb data/country.mmdb
	@echo ""

geoip-download: $(databases)

certs:
	@echo "Generating test certificates"
	mkdir -p certs
	openssl req \
		-subj "/C=CA/ST=Ontario/L=Toronto/O=Magic Test Ltd./CN=localhost" \
		-newkey rsa:4096 -nodes -keyout $(PWD)/certs/test-localhost.key \
		-x509 -days 365 -out $(PWD)/certs/test-localhost.crt
	@echo ""