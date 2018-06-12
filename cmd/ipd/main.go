package main

import (
	"fmt"
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
	log.Println("ipd2 initializing...")

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

	countryDbPath := viper.GetString("CountryDBPath")
	cityDbPath := viper.GetString("CityDBPath")
	db, err := database.New(countryDbPath, cityDbPath)
	if err != nil {
		log.Fatal(err)
		return 1
	}

	server := http.New(db)
	server.Template = viper.GetString("Template")
	ipHeader := viper.GetString("IPHeader")
	server.IPHeader = ipHeader

	if viper.GetBool("ReverseLookup") {
		log.Println("Enabling reverse lookup")
		server.LookupAddr = iputil.LookupAddr
	}

	if viper.GetBool("PortLookup") {
		log.Println("Enabling port lookup")
		server.LookupPort = iputil.LookupPort
	}

	if ipHeader != "" {
		log.Println(fmt.Printf("Trusting header %s to contain correct remote IP\n", ipHeader))
	}

	listen := viper.GetString("listen")
	tlsConfig := viper.GetStringMapString("tls")
	isTlsSet := viper.IsSet("listen-tls")
	log.Printf("TLS is set: %v", isTlsSet)

	listenTls := ""
	if isTlsSet {
		listenTls = viper.GetString("listen-tls")
	}

	log.Println(fmt.Printf("Listening on http://%s, https://%s, using TLS config %s\n", listen, listenTls, tlsConfig))

	if !testMode {
		errs := server.ListenAndServe(listen, listenTls, tlsConfig)

		select {
		case err := <-errs:
			log.Fatalf("Could not start serving service due to (error: %s)", err)
			return 1
		}
	}

	return 0
}
