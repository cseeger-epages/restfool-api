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
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"

	goji "goji.io"
)

func (a RestAPI) createTLSConf() *tls.Config {

	var minversion uint16 = tls.VersionTLS12
	switch a.Conf.TLS.Minversion {
	case SSL30:
		minversion = tls.VersionSSL30
	case TLS10:
		minversion = tls.VersionTLS10
	case TLS11:
		minversion = tls.VersionTLS11
	case TLS12:
		minversion = tls.VersionTLS12
	default:
		Debug("no tls minversion found using default", map[string]interface{}{"default": "tls12"})
	}

	var curves []tls.CurveID
	for _, v := range a.Conf.TLS.CurvePrefs {
		switch v {
		case CURVEP256:
			curves = append(curves, tls.CurveP256)
		case CURVEP384:
			curves = append(curves, tls.CurveP384)
		case CURVEP521:
			curves = append(curves, tls.CurveP521)
		case X25519:
			curves = append(curves, tls.X25519)
		}
	}

	if curves == nil {
		Debug("no tls curvepref found using default", map[string]interface{}{"default": "p256, p384, p521"})
		curves = []tls.CurveID{tls.CurveP256, tls.CurveP384, tls.CurveP521}
	}

	var ciphers []uint16
	for _, v := range a.Conf.TLS.Ciphers {
		ciphers = append(ciphers, CipherMap[v])
	}

	if ciphers == nil {
		Debug("no tls ciphers found using default", map[string]interface{}{
			"default": "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_GCM_SHA384,",
		})
		ciphers = []uint16{tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, tls.TLS_RSA_WITH_AES_256_GCM_SHA384}
	}

	tlsCfg := &tls.Config{
		MinVersion:               minversion,
		CurvePreferences:         curves,
		PreferServerCipherSuites: a.Conf.TLS.PreferServerCiphers,
		CipherSuites:             ciphers,
	}
	return tlsCfg
}

func (a RestAPI) createServerAndListener(router *goji.Mux, ip string, port string) (*http.Server, net.Listener, error) {

	if port == "" || router == nil {
		return nil, nil, fmt.Errorf("router or port is empty")
	}

	if ip != "" && !isIPAddr(ip) {
		return nil, nil, fmt.Errorf("unknown IPAddress format")
	}

	// default initilization
	listen := fmt.Sprintf("%s:%s", "", port)

	if isV4Addr(ip) {
		listen = fmt.Sprintf("%s:%s", ip, port)
	}
	if isV6Addr(ip) {
		listen = fmt.Sprintf("[%s]:%s", ip, port)
	}
	l, err := net.Listen("tcp", listen)
	if err != nil {
		return nil, nil, err
	}

	tlsCfg := a.createTLSConf()
	srv := &http.Server{
		Addr:         listen,
		Handler:      router,
		TLSConfig:    tlsCfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	return srv, l, nil
}

func isIPAddr(ip string) bool {
	if isV4Addr(ip) {
		return true
	} else if isV6Addr(ip) {
		return true
	}
	return false
}

func isV4Addr(ip string) bool {
	trial := net.ParseIP(ip)
	if trial.To4() != nil {
		return true
	}
	return false
}

func isV6Addr(ip string) bool {
	trial := net.ParseIP(ip)
	if trial.To16() != nil {
		return trial != nil && strings.Contains(ip, ":")
	}
	return false
}
