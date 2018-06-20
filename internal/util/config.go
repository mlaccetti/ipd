package util

import (
	"bytes"
	"flag"
	"os"
	"text/template"

	flags "github.com/spf13/pflag"

	"github.com/spf13/viper"
)

func init() {
	flags.BoolP("help", "h", false, "Show this help message")
	flags.BoolP("verbose", "v", false, "Verbose output (default false)")
	flags.StringP("listen", "l", ":8080", "Listening address")
	flags.StringP("listen-tls", "s", ":8443", "Listening address for TLS")
	flags.StringP("tls-key", "k", "", "Path to the TLS key to use (ignored if no TLS listen address is specified)")
	flags.StringP("tls-cert", "e", "", "Path to the TLS certificate to use (ignored if no TLS listen address is specified)")
	flags.StringP("country-db", "f", "", "Path to GeoIP country database")
	flags.StringP("city-db", "c", "", "Path to GeoIP city database")
	flags.BoolP("port-lookup", "p", true, "Perform port lookups")
	flags.BoolP("reverse-lookup", "r", true, "Perform reverse hostname lookups")
	flags.StringP("template", "t", "index.html", "Path to template")
	flags.StringP("trusted-header", "H", "X-Forwarded-For", "Header with 'real' IP, if present")
}

type HelpTemplate struct {
	Flags string
}

const helpTemplate = `
Usage:
  ipd2 [OPTIONS]

Application Options:
{{.Flags}}

Help Options:
  -h, --help                    Show this help message
`

var buf = new(bytes.Buffer)

func Config() (*viper.Viper) {
	v := viper.New()

	flags.CommandLine.SortFlags = false
	flags.CommandLine.MarkHidden("help")
	flags.CommandLine.SetOutput(buf)

	flags.CommandLine.AddGoFlagSet(flag.CommandLine)
	flags.Parse()
	v.BindPFlags(flags.CommandLine)

	v.AutomaticEnv()

	return v
}

func PrintHelp() {
	flags.CommandLine.PrintDefaults()
	data := HelpTemplate{buf.String()}
	buf = new(bytes.Buffer)

	t := template.Must(template.New("help").Parse(helpTemplate))
	t.Execute(os.Stdout, data)
}
