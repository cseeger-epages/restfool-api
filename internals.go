package restfool

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ParseQueryStrings parses filters
func ParseQueryStrings(r *http.Request) QueryStrings {
	vals := r.URL.Query()

	// set defaults
	qs := QueryStrings{false}

	// Parse
	_, ok := vals["prettify"]
	if ok {
		qs.prettify = true
	}

	return qs
}

// EncodeAndSend handles some filters, json encodes and outputs msg
func EncodeAndSend(w http.ResponseWriter, r *http.Request, qs QueryStrings, msg interface{}) {

	var err error
	// i need to encode the data twice for checking etag
	// and for sending with/without prettyfy maybe there
	// is a better way
	etagdata, err := json.Marshal(msg)
	Error("json marshal error etag", err)
	etagsha := sha256.Sum256([]byte(etagdata))
	etag := fmt.Sprintf("%x", etagsha)
	w.Header().Set("ETag", etag)

	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, etag) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	w.WriteHeader(http.StatusOK)

	if qs.prettify {
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", " ")
		err = encoder.Encode(msg)
	} else {
		err = json.NewEncoder(w).Encode(msg)
	}
	Error("json parse error", err)
}

// Serve creates and starts the restfull server and listener
func (a *RestAPI) Serve() error {
	err := a.initRoutes()
	if err != nil {
		return err
	}

	router := a.NewRouter()

	s, l, err := a.createServerAndListener(router, a.Conf.General.Listen, a.Conf.General.Port)
	if err != nil {
		return err
	}

	Info("starting server", map[string]interface{}{"ip": a.Conf.General.Listen, "port": a.Conf.General.Port})
	if a.Conf.TLS.Encryption {
		Info("SSL", map[string]interface{}{"status": "enabled", "cert": a.Conf.Certs.Public, "key": a.Conf.Certs.Private})
		err = s.ServeTLS(l, a.Conf.Certs.Public, a.Conf.Certs.Private)
	} else {
		Info("SSL", map[string]interface{}{"status": "disabled"})
		err = s.Serve(l)
	}

	if err != nil {
		return err
	}
	return nil
}

// New is the restfool constructor
func New(confFile string) (RestAPI, error) {
	var conf config
	err := parseConfig(confFile, &conf)
	if err != nil {
		return RestAPI{}, err
	}

	api := RestAPI{conf, []Route{}}
	api.initLogger()
	if err != nil {
		return RestAPI{}, err
	}
	Info("Basic Authentication", map[string]interface{}{"enabled": conf.General.BasicAuth})
	if api.Conf.TLS.Encryption {
		Info("HTTP Strict Transport Security", map[string]interface{}{"enabled": conf.TLS.Hsts})
	}
	Info("Cross Origin Policy", map[string]interface{}{"enabled": conf.Cors.AllowCrossOrigin})
	return api, nil
}
