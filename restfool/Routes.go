/*
   GOLANG REST API Skeleton

   Copyright (C) 2017 Carsten Seeger

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
   @copyright Copyright (C) 2017 Carsten Seeger
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
	return a.AddHandler("Help", "GET", "/help", "help page", Help)
}

/*
func init() {
	routes = Routes{
		Route{
			"Index",
			"GET",
			"/",
			"default page",
			Index,
		},
		Route{
			"Help",
			"GET",
			"/help",
			"help page",
			Help,
		},
		Route{
			"projects",
			"GET",
			"/projects",
			"show all projects",
			ProjectsHandler,
		},
		Route{
			"project",
			"GET",
			"/project/{project}",
			"get specific project",
			ProjectHandler,
		},
		Route{
			"project",
			"POST",
			"/project/{project}",
			map[string]interface{}{
				"Message": "description message",
				"Post-parameter": map[string]string{
					"parameter": "type - description",
				},
			},
			ProjectHandler,
		},
		Route{
			"Cors",
			"OPTIONS",
			"/{endpoint}",
			"Cross origin preflight",
			CorsHandler,
		},
		Route{
			"Cors",
			"OPTIONS",
			"/{endpoint}/{id}",
			"Cross origin preflight",
			CorsHandler,
		},
	}
}
*/
