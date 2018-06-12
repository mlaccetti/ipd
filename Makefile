OS := $(shell uname)
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

build:
	@echo "Building the ipd2 app"
	go build -o build/ipd2 ./cmd/ipd/main.go
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