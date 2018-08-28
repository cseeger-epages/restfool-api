# Simple Restfool JSON API

[![GoDoc](https://img.shields.io/badge/godoc-reference-green.svg)](https://godoc.org/github.com/cseeger-epages/restfool-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/cseeger-epages/restfool-go)](https://goreportcard.com/report/github.com/cseeger-epages/restfool-go)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/cseeger-epages/restfool-go/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/cseeger-epages/restfool-go.svg?branch=master)](https://travis-ci.org/cseeger-epages/restfool-go)

is a stupidly "foolish" and simple approach of implementing a JSON Restful API library.

## Features
- path routing using gorilla mux
- versioning
- database wrapper
- TLS
- pretty print
- Etag / If-None-Match Clientside caching
- rate limiting and headers using trottled middleware
- basic auth
- config using TOML format
- error handler
- logging

### Ratelimit Headers
```
X-Ratelimit-Limit - The number of allowed requests in the current period
X-Ratelimit-Remaining - The number of remaining requests in the current period
X-Ratelimit-Reset - The number of seconds left in the current period
```

## tbd
- X-HTTP-Method-Override
- caching serverside (varnish support ?)
- Authentication - oauth(2)

## simple example

```
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	rest "github.com/cseeger-epages/restfool-go"
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
```

## config example 

```
# default API configuration file
# using the TOML File format

[general]
# listen supports also IPv6 addresses like :: or ::1
listen = "127.0.0.1"
port = "9443"
basicauth = false

[certs]
public = "certs/server.crt"
private = "certs/server.key"

[tls]
# supported minimal ssl/tls version
# minversion = ["ssl30", "tls10", "tls11", "tls12"]
minversion = "tls12"
# used eliptical curves
# curveprefs = ["p256","p384","p521","x25519"]
curveprefs = ["p256","p384","p521"]
# allowed ciphers
# ciphers = [        
#  "TLS_RSA_WITH_RC4_128_SHA",
#  "TLS_RSA_WITH_3DES_EDE_CBC_SHA",
#  "TLS_RSA_WITH_AES_128_CBC_SHA",
#  "TLS_RSA_WITH_AES_256_CBC_SHA",
#  "TLS_RSA_WITH_AES_128_CBC_SHA256",
#  "TLS_RSA_WITH_AES_128_GCM_SHA256",
#  "TLS_RSA_WITH_AES_256_GCM_SHA384",
#  "TLS_ECDHE_ECDSA_WITH_RC4_128_SHA",
#  "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA",
#  "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA",
#  "TLS_ECDHE_RSA_WITH_RC4_128_SHA",
#  "TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA",
#  "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
#  "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
#  "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256",
#  "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256",
#  "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
#  "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
#  "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
#  "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
#  "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305",
#  "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305",
#]
ciphers = [
    "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
    "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
    "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
    "TLS_RSA_WITH_AES_256_GCM_SHA384",
]
# if not set equal to false
preferserverciphers = true
# HTTP Strict Transport Security
hsts = true
hstsmaxage = 63072000

# cross origin policy
[cors]
allowcrossorigin = true
# corsmethods = ["POST", "GET", "OPTIONS", "PUT", "DELETE"]
corsmethods = ["POST", "GET"]
allowfrom = "https://localhost:8443"

[logging]
# type = ["text","json"]
type = "text"
# loglevel = ["info","error","debug"]
loglevel = "debug"
# output = ["stdout","logfile"]
output = "stdout"
# only if output = "logfile"
logfile = "mylog.log"

[ratelimit]
limit = 1500
burst = 300

[[user]]
username = "testuser"
password = "testpass"

[[user]]
username = "username"
password = "password"
```

Dont fool the reaper ?
