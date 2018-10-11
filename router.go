/*
   Restfool-go

   Copyright (C) 2018 Carsten Seeger

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.

   @author Carsten Seeger
   @copyright Copyright (C) 2018 Carsten Seeger
   @license http://www.gnu.org/licenses/gpl-3.0 GNU General Public License 3
   @link https://github.com/cseeger-epages/rest-api-go-skeleton
*/

package restfool

import (
	"net/http"

	"goji.io"
	"goji.io/pat"
	throttled "gopkg.in/throttled/throttled.v2"
	"gopkg.in/throttled/throttled.v2/store/memstore"
)

// NewRouter is the router constructor
func (a RestAPI) NewRouter() *goji.Mux {

	//router := mux.NewRouter().StrictSlash(true)
	router := goji.NewMux()

	// latest
	a.AddRoutes(router)

	return router
}

// AddRoutes add default handler, routing and ratelimit
func (a RestAPI) AddRoutes(router *goji.Mux) {
	store, err := memstore.New(65536)
	Error("ROUTES: could not create memstore", err)

	// rate limiter
	quota := throttled.RateQuota{
		MaxRate:  throttled.PerMin(a.Conf.RateLimit.Limit),
		MaxBurst: a.Conf.RateLimit.Burst,
	}
	rateLimiter, err := throttled.NewGCRARateLimiter(store, quota)
	Error("ROUTES: error in ratelimiting", err)

	httpRateLimiter := throttled.HTTPRateLimiter{
		RateLimiter: rateLimiter,
		VaryBy:      &throttled.VaryBy{Path: true},
	}
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		var pattern *pat.Pattern

		switch route.Method {
		case "DELETE":
			pattern = pat.Delete(route.Pattern)
		case "GET":
			pattern = pat.Get(route.Pattern)
		case "HEAD":
			pattern = pat.Head(route.Pattern)
		case "NEW":
			pattern = pat.New(route.Pattern)
		case "OPTIONS":
			pattern = pat.Options(route.Pattern)
		case "PATCH":
			pattern = pat.Patch(route.Pattern)
		case "POST":
			pattern = pat.Post(route.Pattern)
		case "PUT":
			pattern = pat.Put(route.Pattern)
		default:
			pattern = pat.Get(route.Pattern)
		}

		router.Handle(
			pattern,
			use(handler, a.addDefaultHeader, a.basicAuthHandler, httpRateLimiter.RateLimit),
		)

		/*
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(use(handler, a.addDefaultHeader, a.basicAuthHandler, httpRateLimiter.RateLimit))
		*/
	}
}

// Middleware chainer
func use(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}
