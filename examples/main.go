package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"net/http"
	"os"

	rest "github.com/cseeger-epages/restfool-go"
)

func main() {
	confFile := flag.String("c", "conf/api.conf", "path to config ile")
	flag.Parse()

	var conf rest.Config
	err := parseConfig(*confFile, &conf)
	if err != nil {
		log.Fatal(err)
	}

	// initialize rest api using conf file
	api, err := rest.New(conf)
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

// Index handler
func Index(w http.ResponseWriter, r *http.Request) {
	// dont need to cache ?
	w.Header().Set("Cache-Control", "no-store")

	qs := rest.ParseQueryStrings(r)
	message := fmt.Sprintf("Welcome to restfool take a look at https://%s/help", r.Host)
	msg := rest.Msg{Message: message}
	rest.EncodeAndSend(w, r, qs, msg)
}

func parseConfig(fileName string, conf interface{}) error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		rest.Error("config error", err)
		os.Exit(1)
	}
	_, err := toml.DecodeFile(fileName, conf)
	return err
}