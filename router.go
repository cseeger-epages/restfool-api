package restfool

import (
	"net/http"

	"github.com/didip/tollbooth"
	"goji.io"
	"goji.io/pat"
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

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(
			tollbooth.LimitHandler(
				tollbooth.NewLimiter(
					float64(a.Conf.RateLimit.Limit),
					nil,
				),
				handler,
			),
			route.Name,
		)

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
			use(handler, a.addDefaultHeader, a.basicAuthHandler),
		)
	}
}

// Middleware chainer
func use(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}
