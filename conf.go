package restfool

import (
	"os"

	"github.com/BurntSushi/toml"
)

type config struct {
	General   general   `toml:"general"`
	Certs     certs     `toml:"certs"`
	TLS       tlsconf   `toml:"tls"`
	Cors      cors      `toml:"cors"`
	Logging   logging   `toml:"logging"`
	RateLimit rateLimit `toml:"ratelimit"`
	Users     []user    `toml:"user"`
}

type general struct {
	Listen    string
	Port      string
	BasicAuth bool
}

type certs struct {
	Public  string
	Private string
}

type tlsconf struct {
	Minversion          string
	CurvePrefs          []string
	Ciphers             []string
	PreferServerCiphers bool
	Hsts                bool
	HstsMaxAge          int
}

type cors struct {
	AllowCrossOrigin bool
	CorsMethods      []string
	AllowFrom        string
}

type logging struct {
	Type     string
	Loglevel string
	Output   string
	Logfile  string
}

type rateLimit struct {
	Limit int
}

type user struct {
	Username string
	Password string
}

func parseConfig(fileName string, conf interface{}) error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		Error("config error", err)
		os.Exit(1)
	}
	_, err := toml.DecodeFile(fileName, conf)
	return err
}
