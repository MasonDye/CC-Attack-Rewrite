package useragent

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Manager manages a list of user agents.
type Manager struct {
	userAgents []string
	rand       *rand.Rand
}

// NewManager creates a new UserAgentManager and loads user agents from the given file.
func NewManager(filePath string) (*Manager, error) {
	um := &Manager{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	if filePath != "" {
		err := um.loadUserAgents(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read User-Agent file: %w", err)
		}
	}

	return um, nil
}

// loadUserAgents reads user agent list from a file.
func (um *Manager) loadUserAgents(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ua := strings.TrimSpace(scanner.Text())
		if ua != "" {
			um.userAgents = append(um.userAgents, ua)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// GetRandomUserAgent returns a random user agent from the loaded list.
func (um *Manager) GetRandomUserAgent() string {
	if len(um.userAgents) == 0 {
		return ""
	}
	return um.userAgents[um.rand.Intn(len(um.userAgents))]
}


