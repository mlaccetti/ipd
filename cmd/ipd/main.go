package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/mlaccetti/ipd2/internal/http"
	"github.com/mlaccetti/ipd2/internal/iputil"
	"github.com/mlaccetti/ipd2/internal/iputil/database"
	"github.com/mlaccetti/ipd2/internal/util"
)

func main() {
	retVal := runServer(false)
	os.Exit(retVal)
}

func runServer(testMode bool) int {
	viper := util.Config()

	if viper.IsSet("help") && viper.GetBool("help") {
		util.PrintHelp()
		return 0
	}

	if !viper.IsSet("verbose") || !viper.GetBool("verbose") {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	} else {
		log.Println("Verbose mode: enabled")
	}

	countryDbPath := viper.GetString("country-db")
	if countryDbPath != "" {
		log.Printf("Loading country DB from %s", countryDbPath)
	}

	cityDbPath := viper.GetString("city-db")
	if cityDbPath != "" {
		log.Printf("Loading city DB from %s", cityDbPath)
	}

	db, err := database.New(countryDbPath, cityDbPath)
	if err != nil {
		log.Printf("Could not load GeoIP information: %s", err)
		return 1
	}

	if !viper.IsSet("template") || viper.GetString("template") == "" {
		log.Printf("No template set, cannot proceed!")
		return 1
	}

	if !viper.IsSet("listen") || viper.GetString("listen") == "" {
		log.Printf("No listen address/port set, cannot proceed!")
		return 1
	}

	server := http.New(db)
	server.Template = viper.GetString("template")

	if viper.IsSet("trusted-header") {
		ipHeader := viper.GetString("trusted-header")
		server.IPHeader = ipHeader
		log.Printf("Trusting header %s (if found) to contain correct remote IP\n", ipHeader)
	}

	if viper.IsSet("reverse-lookup") && viper.GetBool("reverse-lookup") {
		log.Println("Enabling reverse lookup")
		server.LookupAddr = iputil.LookupAddr
	}

	if viper.IsSet("port-lookup") && viper.GetBool("port-lookup") {
		log.Println("Enabling port lookup")
		server.LookupPort = iputil.LookupPort
	}

	listen := viper.GetString("listen")
	log.Printf("Listening on http://%s", listen)

	isTlsSet := viper.IsSet("listen-tls") && viper.IsSet("tls-key") && viper.IsSet("tls-cert") &&
		viper.GetString("listen-tls") != "" && viper.GetString("tls-key") != "" && viper.GetString("tls-cert") != ""

	listenTls := ""
	tlsKey := ""
	tlsCert := ""

	if isTlsSet {
		listenTls = viper.GetString("listen-tls")
		tlsKey = viper.GetString("tls-key")
		tlsCert = viper.GetString("tls-cert")

		log.Printf("TLS is enabled, listening on https://%s, using certificate %s and key %s", listenTls, tlsCert, tlsKey)
	} else {
		log.Printf("TLS is not enabled, either due to an empty listen address, or a missing certificate/key.")
	}

	if !testMode {
		errs := server.ListenAndServe(listen, listenTls, map[string]string{"cert": tlsCert, "key": tlsKey})

		select {
		case err := <-errs:
			log.Fatalf("Could not start serving service due to (error: %s)", err)
			return 1
		}
	}

	return 0
}
