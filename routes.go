package restfool

import (
	"fmt"
	"net/http"
)

// Route contains all information needed for path routing and help generation
type Route struct {
	Name        string
	Method      string
	Pattern     string
	Description interface{}
	HandlerFunc http.HandlerFunc
}

// AddHandler adds a new handler to the routing list
func (a *RestAPI) AddHandler(name string, method string, path string, description interface{}, callback http.HandlerFunc) error {
	if name == "" || method == "" || path == "" || callback == nil {
		return fmt.Errorf("name, method, path or callback function not set")
	}

	a.Routes = append(a.Routes, Route{
		name,
		method,
		path,
		description,
		callback,
	})
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
