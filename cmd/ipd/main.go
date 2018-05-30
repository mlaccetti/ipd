package main

import (
	"log"
	"os"

	"github.com/mlaccetti/ipd2/http"
	"github.com/mlaccetti/ipd2/iputil"
	"github.com/mlaccetti/ipd2/iputil/database"
)

func main() {
	opts := config()
	if opts.Help {
		printHelp()
		os.Exit(0)
	}

	log.SetFlags(0)
	
	db, err := database.New(opts.CountryDBPath, opts.CityDBPath)
	if err != nil {
		log.Fatal(err)
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
	}
}
