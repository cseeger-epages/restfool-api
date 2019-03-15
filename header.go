package restfool

import (
	"fmt"
	"net/http"
	"strings"
)

func (a *RestAPI) addDefaultHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		if a.Conf.Cors.AllowCrossOrigin {
			w.Header().Set("Access-Control-Allow-Origin", a.Conf.Cors.AllowFrom)
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(a.Conf.Cors.CorsMethods, ","))
			/* tbd
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			*/
		}

		if a.Conf.TLS.Hsts {
			hsts := fmt.Sprintf("max-age=%d; includeSubDomains", a.Conf.TLS.HstsMaxAge)
			w.Header().Add("Strict-Transport-Security", hsts)
		}
		h.ServeHTTP(w, r)
	})
}
