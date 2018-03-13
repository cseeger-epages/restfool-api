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

// help reference for all routes
func (a RestAPI) help(w http.ResponseWriter, r *http.Request) {
	// never cache help commands
	w.Header().Set("Cache-Control", "no-store")

	qs := ParseQueryStrings(r)

	var msg []pathList

	for _, m := range routes {
		msg = append(msg, pathList{m.Method, m.Pattern, m.Description})
	}

	EncodeAndSend(w, r, qs, msg)

}

func (a RestAPI) corsHandler(w http.ResponseWriter, r *http.Request) {
	// caching stuff is handler specific
	w.Header().Set("Cache-Control", "no-store")

	// allow authorization to be send for CORS requests
	if a.Conf.Cors.AllowCrossOrigin {
		w.Header().Set("Access-Control-Allow-Headers", "authorization")
	}

	qs := ParseQueryStrings(r)
	msg := ""
	EncodeAndSend(w, r, qs, msg)
}
