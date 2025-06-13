package config

import (
	"flag"
	"fmt"
	"strings"
)

type Config struct {
	TargetURL       string
	RequestInterval int
	Timeout         int
	UserAgentFile   string
	ProxyListFile   string
	ThreadCount     int
	HTTPVersion     float64
	AttackTime      int
	Cookie          string
	Method          string
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
	flag.StringVar(&cfg.Method, "http_methods", "GET", "HTTP Request Method (GET, POST, PUT, DELETE, HEAD, OPTIONS, PATCH, TRACE, CONNECT)")

	flag.Parse()

	if cfg.TargetURL == "" {
		return nil, fmt.Errorf("\033[31mCC Attack ++ Rewrite \033[34mVersion: 2.4.0 (Releases 2024/05/29 9:15 AM)\033[0m\n\033[32mAuthor: MasonDye\033[0m\n\033[32mGitHub: https://github.com/MasonDye/CC-Attack-Rewrite\033[0m\n\n\033[31mUsage:\033[0m\n  -url string\n        Attack URL\n  -cookie string\n        Cookie to include in request\n  -http_methods string\n        HTTP Request Method (GET, POST, PUT, DELETE, HEAD, OPTIONS, PATCH, TRACE, CONNECT) (default \"GET\")\n  -http_version float\n        HTTP version (1.1 or 2.0) (default 1.1)\n  -ip_pool string\n        IP Pool Path (txt)\n  -speed int\n        Attack Speed(ms) (default 100)\n  -thread int\n        thread (default 2)\n  -time int\n        Attack Time (seconds) (default 0)\n  -timeout int\n        Request Timeout (ms) (default 2500)\n  -ua_pool string\n        User-Agent Pool Path (txt)")
	}

	if cfg.HTTPVersion != 1.1 && cfg.HTTPVersion != 2.0 {
		return nil, fmt.Errorf("\033[31mInvalid HTTP version. Please use 1.1 or 2.0.\033[0m")
	}

	validMethods := map[string]bool{"GET": true, "POST": true, "PUT": true, "DELETE": true, "HEAD": true, "OPTIONS": true, "PATCH": true, "TRACE": true, "CONNECT": true}
	if !validMethods[strings.ToUpper(cfg.Method)] {
		return nil, fmt.Errorf("\033[31mInvalid HTTP method. Please use GET, POST, PUT, DELETE, HEAD, OPTIONS, PATCH, TRACE, or CONNECT.\033[0m")
	}

	return cfg, nil
}

