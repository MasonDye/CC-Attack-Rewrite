package proxy

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"
)

// Manager manages a list of proxies.
type Manager struct {
	proxies []string
	rand    *rand.Rand
}

// NewManager creates a new ProxyManager and loads proxies from the given file.
func NewManager(filePath string) (*Manager, error) {
	pm := &Manager{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	if filePath != "" {
		err := pm.loadProxies(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read IP Pool file: %w", err)
		}
	}

	return pm, nil
}

// loadProxies reads proxy list from a file.
func (pm *Manager) loadProxies(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		proxy := strings.TrimSpace(scanner.Text())
		if proxy != "" {
			pm.proxies = append(pm.proxies, proxy)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// GetRandomProxy returns a random proxy URL from the loaded list.
func (pm *Manager) GetRandomProxy() (*url.URL, error) {
	if len(pm.proxies) == 0 {
		return nil, nil // No proxies available
	}
	proxyStr := pm.proxies[pm.rand.Intn(len(pm.proxies))]
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse proxy URL '%s': %w", proxyStr, err)
	}
	return proxyURL, nil
}

// HasProxies checks if there are any proxies loaded.
func (pm *Manager) HasProxies() bool {
	return len(pm.proxies) > 0
}


