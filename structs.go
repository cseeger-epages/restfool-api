package restfool

import (
	"net/http"
)

// RestAPI contains api data
type RestAPI struct {
	Conf   config
	Routes []route
}

// Msg is the standard message type
type Msg struct {
	Message string `json:"message"`
}

// ErrMsg is the standard error message type
type ErrMsg struct {
	Error string `json:"error"`
}

// QueryStrings contains all possible query options
type QueryStrings struct {
	prettify bool
}

type pathList struct {
	Method      string
	Pattern     string
	Description interface{}
}

type pathPrefix struct {
	Prefix  string
	Handler http.Handler
}
