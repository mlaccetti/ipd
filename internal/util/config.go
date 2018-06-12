package util

import (
	"flag"

	flags "github.com/spf13/pflag"

	"github.com/spf13/viper"
)

type opts struct {
	Help          bool
	Verbose       bool
	CountryDBPath string
	CityDBPath    string
	Listen        string
	ListenTLS     string
	TLSKey        string
	TLSCert       string
	ReverseLookup bool
	PortLookup    bool
	Template      string
	IPHeader      string
}

var _ = &opts{
	Help: *flags.BoolP("help", "h", false, "This help text"),
	Verbose: *flags.BoolP("verbose", "v", false, "verbose output (default false"),
	CountryDBPath: *flags.StringP("country-db", "f", "", "Path to GeoIP country database"),
	CityDBPath: *flags.StringP("city-db", "c", "", "Path to GeoIP city database"),
	Listen: *flags.StringP("listen", "l", ":8080", "Listening address"),
	ListenTLS: *flags.StringP("listen-tls", "s", ":8443", "Listening address for TLS"),
	TLSKey: *flags.StringP("tls-key", "k", "", "Path to the TLS key to use (ignored if no TLS listen address is specified)"),
	TLSCert: *flags.StringP("tls-cert", "e", "", "Path to the TLS certificate to use (ignored if no TLS listen address is specified)"),
	ReverseLookup: *flags.BoolP("reverse-lookup", "r", true, "Perform reverse hostname lookups"),
	PortLookup: *flags.BoolP("port-lookup", "p", true, "Perform port lookups"),
	Template: *flags.StringP("template", "t", "index.html", "Path to template"),
	IPHeader: *flags.StringP("trusted-header", "H", "X-Forwarded-For", "Header with 'real' IP, if present"),
}

func Config() (*viper.Viper ) {
	v := viper.New()

	/*v.SetDefault("help", false)
	v.SetDefault("verbose", false)
	v.SetDefault("listen", ":8080")
	v.SetDefault("listen-tls", ":8443")
	v.SetDefault("reverse-lookup", true)
	v.SetDefault("port-lookup", true)
	v.SetDefault("template", "index.html")
	v.SetDefault("trusted-header", "X-Forwarded-For")*/

	flags.CommandLine.SortFlags = false

	flags.CommandLine.AddGoFlagSet(flag.CommandLine)
	flags.Parse()
	v.BindPFlags(flags.CommandLine)

	v.AutomaticEnv()

	return v
}

func PrintHelp() {
	flags.PrintDefaults()
}
