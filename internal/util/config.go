package util

import (
	"flag"
	"log"

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
	ReverseLookup bool
	PortLookup    bool
	Template      string
	IPHeader      string
}

var _ = &opts{
	Help: *flags.BoolP("help", "h", false, "This help text"),
	Verbose: *flags.BoolP("verbose", "v", false, "verbose output"),
	CountryDBPath: *flags.StringP("country-db", "f", "", "Path to GeoIP country database"),
	CityDBPath: *flags.StringP("city-db", "c", "", "Path to GeoIP city database"),
	Listen: *flags.StringP("listen", "l", ":8080", "Listening address"),
	ListenTLS: *flags.StringP("listen-tls", "s", ":8443", "Listening address for TLS"),
	ReverseLookup: *flags.BoolP("reverse-lookup", "r", true, "Perform reverse hostname lookups"),
	PortLookup: *flags.BoolP("port-lookup", "p", true, "Perform port lookups"),
	Template: *flags.StringP("template", "t", "index.html", "Path to template"),
	IPHeader: *flags.StringP("trusted-header", "H", "X-Real-IP", "Header with 'real' IP, if present (i.e. X-Real-IP)"),
}

func Config() (*viper.Viper, error ) {
	log.Println("Configuration initializing...")

	flags.CommandLine.SortFlags = false

	flags.CommandLine.AddGoFlagSet(flag.CommandLine)
	flags.Parse()
	viper.BindPFlags(flags.CommandLine)

	viper.AutomaticEnv()

	viper.SetConfigName("ipd2") // name of util file (without extension)
	viper.AddConfigPath("/etc/ipd2/")   // path to look for the util file in
	viper.AddConfigPath("$HOME/.ipd2")  // call multiple times to add many search paths
	viper.AddConfigPath("./configs")            // optionally look for util in the working directory

	err := viper.ReadInConfig() // Find and read the util file
	if err != nil { // Handle errors reading the util file
		return nil, err
	}

	log.Println("Configuration completed.")
	return viper.GetViper(), nil
}

func PrintHelp() {
	flags.PrintDefaults()
}
