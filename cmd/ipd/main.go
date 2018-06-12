package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mlaccetti/ipd2/internal/http"
	"github.com/mlaccetti/ipd2/internal/iputil"
	"github.com/mlaccetti/ipd2/internal/iputil/database"
	"github.com/mlaccetti/ipd2/internal/util"
)

func main() {
	retVal := runServer()
	os.Exit(retVal)
}

func runServer() int {
	log.Println("Kicking off ipd2")

	viper, err := util.Config()
	if err != nil {
		log.Fatal(err)
		return 1
	}

	if viper.GetBool("Help") {
		log.Println("Help mode invoked")
		util.PrintHelp()
		return 0
	}

	log.SetFlags(0)

	db, err := database.New(viper.GetString("CountryDBPath"), viper.GetString("CityDBPath"))
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

	listen := viper.GetString("Listen")
	tlsConfig := viper.GetStringMapString("tls")
	isTlsSet := viper.IsSet("listen-tls")
	log.Printf("TLS is set: %v", isTlsSet)

	listenTls := ""
	if isTlsSet {
		listenTls = viper.GetString("listen-tls")
	}

	log.Println(fmt.Printf("Listening on http://%s, https://%s, using TLS config %s\n", listen, listenTls, tlsConfig))
	errs := server.ListenAndServe(listen, listenTls, tlsConfig)
	select {
	case err := <-errs:
		log.Fatalf("Could not start serving service due to (error: %s)", err)
		return 1
	}


	return 0
}
