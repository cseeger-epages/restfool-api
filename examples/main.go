package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	rest "../restfool"
)

func main() {
	confFile := flag.String("c", "conf/api.conf", "path to config ile")
	flag.Parse()

	// initialize rest api using conf file
	api, err := rest.New(*confFile)
	if err != nil {
		log.Fatal(err)
	}

	// add handler
	err = api.AddHandler("Index", "GET", "/", "default page", Index)
	if err != nil {
		log.Fatal(err)
	}

	// start
	err = api.Serve()
	if err != nil {
		log.Fatal(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	// dont need to cache ?
	w.Header().Set("Cache-Control", "no-store")

	qs := rest.ParseQueryStrings(r)
	message := fmt.Sprintf("Welcome to restfool take a look at https://%s/help", r.Host)
	msg := rest.Msg{Message: message}
	rest.EncodeAndSend(w, r, qs, msg)
}
