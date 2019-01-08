package restfool

import (
	"net/http"
)

/*
Default Handler Template
func Handler(w http.ResponseWriter, r*http.Request) {
	// caching stuff is handler specific
	w.Header().Set("Cache-Control", "no-store")

	qs := ParseQueryStrings(r)
	msg := HelpMsg{Message: "im a default Handler"}
	EncodeAndSend(w, r, qs, msg)
}
*/

// help reference for all routes
func (a *RestAPI) help(w http.ResponseWriter, r *http.Request) {
	// never cache help commands
	w.Header().Set("Cache-Control", "no-store")

	qs := ParseQueryStrings(r)

	var msg []pathList

	for _, m := range routes {
		msg = append(msg, pathList{m.Method, m.Pattern, m.Description})
	}

	EncodeAndSend(w, r, qs, msg)

}

func (a *RestAPI) corsHandler(w http.ResponseWriter, r *http.Request) {
	// caching stuff is handler specific
	w.Header().Set("Cache-Control", "no-store")

	// allow authorization to be send for CORS requests
	if a.Conf.Cors.AllowCrossOrigin {
		w.Header().Set("Access-Control-Allow-Headers", "authorization")
	}

	qs := ParseQueryStrings(r)
	msg := ""
	EncodeAndSend(w, r, qs, msg)
}
