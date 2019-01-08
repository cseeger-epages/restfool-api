package restfool

import (
	"fmt"
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	Description interface{}
	HandlerFunc http.HandlerFunc
}

var routes []route

// AddHandler adds a new handler to the routing list
func (a *RestAPI) AddHandler(name string, method string, path string, description interface{}, callback http.HandlerFunc) error {
	if name == "" || method == "" || path == "" || callback == nil {
		return fmt.Errorf("name, method, path or callback function not set")
	}

	routes = append(routes, route{
		name,
		method,
		path,
		description,
		callback,
	})
	a.Routes = routes
	return nil
}

func (a *RestAPI) initRoutes() error {
	err := a.AddHandler("Help", "GET", "/help", "help page", a.help)
	if err != nil {
		return err
	}
	err = a.AddHandler("Help", "POST", "/help", "help page", a.help)
	if err != nil {
		return err
	}
	err = a.AddHandler("Cors", "OPTIONS", "/{endpoint}", "Cross origin preflight", a.corsHandler)
	if err != nil {
		return err
	}
	err = a.AddHandler("Cors", "OPTIONS", "/{endpoint}/{id}", "Cross origin preflight", a.corsHandler)
	return err
}
