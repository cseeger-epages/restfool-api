/*
Package restfool is a stupidly "foolish" and simple approach of implementing a JSON Restful API.

Initialize the api

	confFile := flag.String("c", "conf/api.conf", "path to config ile")
	flag.Parse()

	api, err := rest.New(*confFile)
	if err != nil {
		log.Fatal(err)
	}

Add your own handler and go

	err = api.AddHandler("Index", "GET", "/", "default page", Index)
	if err != nil {
		log.Fatal(err)
	}

	err = api.Serve()
	if err != nil {
		log.Fatal(err)
	}

Handler definition

	func Index(w http.ResponseWriter, r *http.Request) {
		qs := rest.ParseQueryStrings(r)
		message := fmt.Sprintf("Welcome to restfool take a look at https://%s/help", r.Host)
		msg := rest.Msg{Message: message}
		rest.EncodeAndSend(w, r, qs, msg)
	}

Configuration example

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

Additional filters

As every useful restful JSON API `?prettify` can be used for pretty print.

Default handlers

By default the throttled.HTTPRateLimiter is used, Basic Auth is supported and CORS Headers are set by the default handlers.

Default routes/path

By default the /help path is added and provides the description interface{} of your added handlers.
A nice looking description could look like this:

	description := map[string]interface{}{
		"Message": "description message",
		"Post-parameter": map[string]string{
			"parameter": "type - description",
		},
	}
*/
package restfool
