OS := $(shell uname)
TARGET := ipd2

ifeq ($(OS),Linux)
	TAR_OPTS := --wildcards
endif

guard-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Environment variable $* not set"; \
		exit 0; \
	fi

all: deps test vet build

fmt:
	@echo "Formatting all the things..."
	go fmt ./...
	@echo ""

vet:
	@echo "Vetting stuff"
	go vet ./...
	@echo ""

deps:
	@echo "Ensuring dependencies are in place"
	dep ensure
	@echo ""
	
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

test: geoip-download certs
	@echo "Running tests"
	go test ./...
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
	env GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -a -installsuffix cgo -o build/$(TARGET)-${TRAVIS_TAG}-$(GOOS)_$(GOARCH)$(EXT) ./cmd/ipd/main.go

docker-build:
	@echo "Building Docker image for compiling ipd2"
	docker build --build-arg TRAVIS_TAG=${TRAVIS_TAG} --target build --tag mlaccetti/ipd2:${TRAVIS_TAG}-build .

docker-release:	guard-TRAVIS_TAG
	@echo "Building Docker image for release"
	docker build --build-arg TRAVIS_TAG=${TRAVIS_TAG} --target runtime --tag mlaccetti/ipd2:${TRAVIS_TAG} .

release: docker-build
	set -e ;\
	CONTAINER_ID=$$(docker run -d mlaccetti/ipd2:$$TRAVIS_TAG-build /bin/false) ;\
	echo "Copying files from Docker container $$CONTAINER_ID for release" ;\
	rm -fr build ;\
	mkdir -p build ;\
	docker cp $$CONTAINER_ID:/go/src/github.com/mlaccetti/ipd2/build/ipd2-$$TRAVIS_TAG-darwin_amd64 build/ipd2-$$TRAVIS_TAG-darwin_amd64 ;\
	docker cp $$CONTAINER_ID:/go/src/github.com/mlaccetti/ipd2/build/ipd2-$$TRAVIS_TAG-linux_amd64 build/ipd2-$$TRAVIS_TAG-linux_amd64 ;\
	docker cp $$CONTAINER_ID:/go/src/github.com/mlaccetti/ipd2/build/ipd2-$$TRAVIS_TAG-windows_amd64.exe build/ipd2-$$TRAVIS_TAG-windows_amd64.exe
