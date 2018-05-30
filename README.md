# ipd

[![Build Status](https://travis-ci.org/mpolden/ipd.svg)](https://travis-ci.org/mpolden/ipd)

A simple service for looking up your IP address. This is the code that powers
https://ifconfig2.co.

## Usage

Just the business, please:

```
$ curl ifconfig2.co
127.0.0.1

$ http ifconfig2.co
127.0.0.1

$ wget -qO- ifconfig2.co
127.0.0.1

$ fetch -qo- https://ifconfig2.co
127.0.0.1

$ bat -print=b ifconfig2.co/ip
127.0.0.1
```

Country and city lookup:

```
$ curl ifconfig2.co/country
Elbonia

$ curl ifconfig2.co/country-iso
EB

$ curl ifconfig2.co/city
Bornyasherk
```

As JSON:

```
$ curl -H 'Accept: application/json' ifconfig2.co  # or curl ifconfig2.co/json
{
  "city": "Bornyasherk",
  "country": "Elbonia",
  "country_iso": "EB",
  "ip": "127.0.0.1",
  "ip_decimal": 2130706433
}
```

Port testing:

```
$ curl ifconfig2.co/port/80
{
  "ip": "127.0.0.1",
  "port": 80,
  "reachable": false
}
```

Pass the appropriate flag (usually `-4` and `-6`) to your client to switch
between IPv4 and IPv6 lookup.

The subdomains https://v4.ifconfig2.co and https://v6.ifconfig2.co can be used to
force IPv4 or IPv6 lookup.

## Features

* Easy to remember domain name
* Fast
* Supports IPv6
* Supports HTTPS
* Supports common command-line clients (e.g. `curl`, `httpie`, `wget` and `fetch`)
* JSON output
* Country and city lookup using the MaxMind GeoIP database
* Port testing
* Open source under the [BSD 3-Clause license](https://opensource.org/licenses/BSD-3-Clause)

## Why?

* To scratch an itch
* An excuse to use Go for something
* Faster than ifconfig.me and has IPv6 support

## Building

Compiling requires the [Golang compiler](https://golang.org/) to be installed.
This package can be installed with `go get`:

`go get github.com/mlaccetti/ipd2/...`

For more information on building a Go project, see the [official Go
documentation](https://golang.org/doc/code.html).

### Usage

```
$ ipd -h
Usage:
  ipd [OPTIONS]

Application Options:
  -f, --country-db=FILE        Path to GeoIP country database
  -c, --city-db=FILE           Path to GeoIP city database
  -l, --listen=ADDR            Listening address (default: :8080)
  -r, --reverse-lookup         Perform reverse hostname lookups
  -p, --port-lookup            Enable port lookup
  -t, --template=FILE          Path to template (default: index.html)
  -H, --trusted-header=NAME    Header to trust for remote IP, if present (e.g. X-Real-IP)

Help Options:
  -h, --help                   Show this help message
```
