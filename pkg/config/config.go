package config

import (
	"flag"
	"fmt"
	"strings"

	"github.com/MasonDye/CC-Attack-Rewrite/pkg/version"
)

const usage = "\x1b[31mCC Attack ++ Rewrite \x1b[34mVersion: %s (Releases %s)\x1b[0m\n" +
	"\x1b[32mAuthor: MasonDye\x1b[0m\n" +
	"\x1b[32mGitHub: https://github.com/MasonDye/CC-Attack-Rewrite\x1b[0m\n\n" +
	"\x1b[31mUsage:\x1b[0m\n" +
	"  -url string\n" +
	"        Attack URL\n" +
	"  -cookie string\n" +
	"        Cookie to include in request\n" +
	"  -data string\n" +
	"        Data to send in the request body (for POST/PUT requests)\n" +
	"  -http_methods string\n" +
	"        HTTP Request Method (GET, POST, PUT, DELETE, HEAD, OPTIONS, PATCH, TRACE, CONNECT) (default \"GET\")\n" +
	"  -http_version float\n" +
	"        HTTP version (1.1 or 2.0) (default 1.1)\n" +
	"  -ip_pool string\n" +
	"        IP Pool Path (txt)\n" +
	"  -speed int\n" +
	"        Attack Speed(ms) (default 100)\n" +
	"  -thread int\n" +
	"        Thread count (default 2)\n" +
	"  -time int\n" +
	"        Attack Time (seconds) (default 0)\n" +
	"  -timeout int\n" +
	"        Request Timeout (ms) (default 2500)\n" +
	"  -ua_pool string\n" +
	"        User-Agent Pool Path (txt)"

type Config struct {
	TargetURL       string  `flag:"url" description:"Attack URL"`
	RequestInterval int     `flag:"speed" default:"100" description:"Attack Speed(ms)"`
	Timeout         int     `flag:"timeout" default:"2500" description:"Request Timeout (ms)"`
	UserAgentFile   string  `flag:"ua_pool" description:"User-Agent Pool Path (txt)"`
	ProxyListFile   string  `flag:"ip_pool" description:"IP Pool Path (txt)"`
	ThreadCount     int     `flag:"thread" default:"2" description:"Thread count"`
	HTTPVersion     float64 `flag:"http_version" default:"1.1" description:"HTTP version (1.1 or 2.0)"`
	AttackTime      int     `flag:"time" default:"0" description:"Attack Time (seconds)"`
	Cookie          string  `flag:"cookie" description:"Cookie to include in request"`
	Method          string  `flag:"http_methods" default:"GET" description:"HTTP Request Method"`
	Data            string  `flag:"data" description:"Data to send in the request body"`
}

func (c *Config) validate() error {
	if c.TargetURL == "" {
		return fmt.Errorf(usage, version.Version, version.BuildDate)
	}

	if c.HTTPVersion != 1.1 && c.HTTPVersion != 2.0 {
		return fmt.Errorf("\033[31mInvalid HTTP version. Please use 1.1 or 2.0.\033[0m")
	}

	validMethods := map[string]bool{
		"GET":     true,
		"POST":    true,
		"PUT":     true,
		"DELETE":  true,
		"HEAD":    true,
		"OPTIONS": true,
		"PATCH":   true,
		"TRACE":   true,
		"CONNECT": true,
	}

	if !validMethods[strings.ToUpper(c.Method)] {
		return fmt.Errorf("\033[31mInvalid HTTP method. Please use GET, POST, PUT, DELETE, HEAD, OPTIONS, PATCH, TRACE, or CONNECT.\033[0m")
	}

	return nil
}

func ParseConfig() (*Config, error) {
	cfg := &Config{}

	flag.StringVar(&cfg.TargetURL, "url", "", "Attack URL")
	flag.IntVar(&cfg.RequestInterval, "speed", 100, "Attack Speed(ms)")
	flag.IntVar(&cfg.Timeout, "timeout", 2500, "Request Timeout (ms)")
	flag.StringVar(&cfg.UserAgentFile, "ua_pool", "", "User-Agent Pool Path (txt)")
	flag.StringVar(&cfg.ProxyListFile, "ip_pool", "", "IP Pool Path (txt)")
	flag.IntVar(&cfg.ThreadCount, "thread", 2, "thread")
	flag.Float64Var(&cfg.HTTPVersion, "http_version", 1.1, "HTTP version (1.1 or 2.0)")
	flag.IntVar(&cfg.AttackTime, "time", 0, "Attack Time (seconds)")
	flag.StringVar(&cfg.Cookie, "cookie", "", "Cookie to include in request")
	flag.StringVar(&cfg.Method, "http_methods", "GET", "HTTP Request Method")
	flag.StringVar(&cfg.Data, "data", "", "Data to send in the request body")

	flag.Parse()

	return cfg, cfg.validate()
}
