package restfool

import (
	"net/http"
	"strings"
)

func (a *RestAPI) basicAuthHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			h.ServeHTTP(w, r)
			return
		}

		if !a.Conf.General.BasicAuth {
			h.ServeHTTP(w, r)
			return
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		username, password, authOK := r.BasicAuth()
		if authOK == false {
			http.Error(w, "Not authorized", 401)
			xff := r.Header.Get("X-FORWARDED-FOR")
			if xff == "" {
				xff = "not set"
			}
			Debug("Authorization error", map[string]interface{}{
				"RemoteAddr":     r.RemoteAddr,
				"X-FORWARDD-FOR": xff,
			})
			return
		}

		valid := false

		for _, v := range a.Conf.Users {
			if username == v.Username && strings.TrimSuffix(password, "\n") == v.Password {
				valid = true
			}
		}
		if !valid {
			http.Error(w, "Not authorized", 401)
			return
		}

		h.ServeHTTP(w, r)
	})
}
