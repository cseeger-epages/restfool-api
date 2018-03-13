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
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
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

// root handler giving basic API information
func Index(w http.ResponseWriter, r *http.Request) {
	// dont know what should happen here
	w.Header().Set("Cache-Control", "no-store")

	qs := ParseQueryStrings(r)
	message := fmt.Sprintf("Welcome to GOLANG REST API SKELETON please take a look at https://%s/help", r.Host)
	msg := HelpMsg{Message: message}
	EncodeAndSend(w, r, qs, msg)
}

// help reference for all routes
func Help(w http.ResponseWriter, r *http.Request) {
	// never cache help commands
	w.Header().Set("Cache-Control", "no-store")

	qs := ParseQueryStrings(r)

	var msg []PathList

	for _, m := range routes {
		msg = append(msg, PathList{m.Method, m.Pattern, m.Description})
	}

	EncodeAndSend(w, r, qs, msg)

}

// Handler returning a list of all Projects
func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	// 30 days cache since project did not change often
	w.Header().Set("Cache-Control", "max-age=2592000")

	qs := ParseQueryStrings(r)
	projects := Projects{[]Project{Project{1, "P1"}, Project{2, "P2"}, Project{3, "P3"}}}
	EncodeAndSend(w, r, qs, projects)
}

// Project specific information
func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	// 1 day cache since data only change once a day
	w.Header().Set("Cache-Control", "max-age=86400")

	qs := ParseQueryStrings(r)
	vars := mux.Vars(r)
	project := vars["project"]

	var msg interface{}
	var projects Projects
	var perr error = nil

	projectlist := []Project{Project{1, "P1"}, Project{2, "P2"}, Project{3, "P3"}}

	if pid, err := strconv.Atoi(project); err == nil {
		projects, perr = Projects{}, errors.New("project does not exists")
		for _, v := range projectlist {
			if v.Id == pid {
				projects, perr = Projects{[]Project{v}}, nil
			}
		}
	} else {
		project = strings.ToLower(project)
		projects, perr = Projects{}, errors.New("project does not exists")
		for _, v := range projectlist {
			if strings.ToLower(v.Name) == project {
				projects, perr = Projects{[]Project{v}}, nil
			}
		}
	}
	if perr != nil {
		msg = ErrorMessage{perr.Error()}
	} else {
		msg = projects
	}
	EncodeAndSend(w, r, qs, msg)
}

func CorsHandler(w http.ResponseWriter, r *http.Request) {
	// caching stuff is handler specific
	w.Header().Set("Cache-Control", "no-store")

	// allow authorization to be send for CORS requests
	if Conf.Cors.AllowCrossOrigin {
		w.Header().Set("Access-Control-Allow-Headers", "authorization")
	}

	qs := ParseQueryStrings(r)
	msg := ""
	EncodeAndSend(w, r, qs, msg)
}
