package restfool

// config contains all config information
type Config struct {
	General   General   `toml:"general"`
	Certs     Certs     `toml:"certs"`
	TLS       TlsConf   `toml:"tls"`
	Cors      Cors      `toml:"cors"`
	Logging   Logging   `toml:"logging"`
	RateLimit RateLimit `toml:"ratelimit"`
	Users     []User    `toml:"user"`
}

type General struct {
	Listen    string
	Port      string
	BasicAuth bool
}

type Certs struct {
	Public  string
	Private string
}

type TlsConf struct {
	Encryption          bool
	Minversion          string
	CurvePrefs          []string
	Ciphers             []string
	PreferServerCiphers bool
	Hsts                bool
	HstsMaxAge          int
}

type Cors struct {
	AllowCrossOrigin bool
	CorsMethods      []string
	AllowFrom        string
}

type Logging struct {
	Type     string
	Loglevel string
	Output   string
	Logfile  string
}

type RateLimit struct {
	Limit int
}

type User struct {
	Username string
	Password string
}
