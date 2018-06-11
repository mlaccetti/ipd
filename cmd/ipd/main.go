package main

import (
	"log"
	"os"
	"reflect"

	"github.com/mlaccetti/ipd2/http"
	"github.com/mlaccetti/ipd2/iputil"
	"github.com/mlaccetti/ipd2/iputil/database"
)

func main() {
	viper, opts, err := config()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if viper.GetBool(reflect.Indirect(reflect.ValueOf(opts.Help)).Type().Field(0).Name) {
		printHelp()
		os.Exit(0)
	}

	log.SetFlags(0)
	
	db, err := database.New(opts.CountryDBPath, opts.CityDBPath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	server := http.New(db)
	server.Template = opts.Template
	server.IPHeader = opts.IPHeader
	if opts.ReverseLookup {
		log.Println("Enabling reverse lookup")
		server.LookupAddr = iputil.LookupAddr
	}
	if opts.PortLookup {
		log.Println("Enabling port lookup")
		server.LookupPort = iputil.LookupPort
	}
	if opts.IPHeader != "" {
		log.Printf("Trusting header %s to contain correct remote IP", opts.IPHeader)
	}

	log.Printf("Listening on http://%s", opts.Listen)
	if err := server.ListenAndServe(opts.Listen); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
