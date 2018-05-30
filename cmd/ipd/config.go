package main

import (
	flags "github.com/spf13/pflag"

	"github.com/spf13/viper"
)

type opts struct {
	Help          bool
	Verbose       bool
	CountryDBPath string
	CityDBPath    string
	Listen        string
	ReverseLookup bool
	PortLookup    bool
	Template      string
	IPHeader      string
}

var opt = &opts{
	Help: *flags.BoolP("help", "h", false, "This help text"),
	Verbose: *flags.BoolP("verbose", "v", false, "verbose output"),
	CountryDBPath: *flags.StringP("country-db", "f", "", "Path to GeoIP country database"),
	CityDBPath: *flags.StringP("city-db", "c", "", "Path to GeoIP city database"),
	Listen: *flags.StringP("listen", "l", ":8080", "Listening address"),
	ReverseLookup: *flags.BoolP("reverse-lookup", "r", true, "Perform reverse hostname lookups"),
	PortLookup: *flags.BoolP("port-lookup", "p", true, "Perform port lookups"),
	Template: *flags.StringP("template", "t", "index.html", "Path to template"),
	IPHeader: *flags.StringP("trusted-header", "H", "X-Real-IP", "Header with 'real' IP, if present (i.e. X-Real-IP)"),
}

func init() {
	flags.CommandLine.SortFlags = false

	flags.Parse()
	viper.BindPFlags(flags.CommandLine)

	viper.AutomaticEnv()
}

func config() *opts {
	return opt
}

func printHelp() {
	flags.PrintDefaults()
}
