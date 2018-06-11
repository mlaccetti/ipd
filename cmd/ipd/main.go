package main

import (
	"log"
	"os"

	"github.com/mlaccetti/ipd2/http"
	"github.com/mlaccetti/ipd2/iputil"
	"github.com/mlaccetti/ipd2/iputil/database"
)

func main() {
	retVal := runServer()
	os.Exit(retVal)
}

func runServer() int {
	log.Println("Kicking off ipd2")

	viper, err := config()
	if err != nil {
		log.Fatal(err)
		return 1
	}

	if viper.GetBool("Help") {
		log.Println("Help mode invoked")
		printHelp()
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
		log.Printf("Trusting header %s to contain correct remote IP", ipHeader)
	}

	listen := viper.GetString("Listen")
	log.Printf("Listening on http://%s", listen)
	if err := server.ListenAndServe(listen); err != nil {
		log.Fatal(err)
		return 1
	}

	return 0
}
