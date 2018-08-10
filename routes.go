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
func (a RestAPI) AddHandler(name string, method string, path string, description interface{}, callback http.HandlerFunc) error {
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

func (a RestAPI) initRoutes() error {
	err := a.AddHandler("Help", "GET", "/help", "help page", a.help)
	err = a.AddHandler("Help", "POST", "/help", "help page", a.help)
	err = a.AddHandler("Cors", "OPTIONS", "/{endpoint}", "Cross origin preflight", a.corsHandler)
	err = a.AddHandler("Cors", "OPTIONS", "/{endpoint}/{id}", "Cross origin preflight", a.corsHandler)
	return err
}
