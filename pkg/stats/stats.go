package stats

import (
	"fmt"
	"sync"
	"time"
)

// Stats collects and reports attack statistics.
type Stats struct {
	successCount    int64
	errorCount      int64
	startTime       time.Time
	mu              sync.Mutex
	lastReportTime  time.Time
	lastSuccessCount int64
}

// NewStats creates a new Stats instance.
func NewStats() *Stats {
	return &Stats{
		startTime: time.Now(),
		lastReportTime: time.Now(),
	}
}

// IncrementSuccess increments the success count.
func (s *Stats) IncrementSuccess() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.successCount++
}

// IncrementError increments the error count.
func (s *Stats) IncrementError() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.errorCount++
}

// Report prints the current attack statistics.
func (s *Stats) Report(targetURL string, threadCount, requestInterval, timeout int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	elapsed := time.Since(s.startTime).Seconds()
	if elapsed == 0 {
		elapsed = 1 // Avoid division by zero
	}

	// Calculate requests per second based on the last report interval
	intervalElapsed := time.Since(s.lastReportTime).Seconds()
	currentSuccessCount := s.successCount - s.lastSuccessCount
	requestsPerSecond := float64(currentSuccessCount) / intervalElapsed

	fmt.Printf("\r\033[31mRequested:%d \033[0m|\033[31m %.1f p/s \033[0m|\033[31m URL:%s \033[0m|\033[31m Thread:%d \033[0m|\033[31m Speed:%d\033[31m ms \033[0m|\033[31m Timeout:%d\033[31m ms \033[0m|\033[31m Error:%d \033[0m",
		s.successCount, requestsPerSecond, targetURL, threadCount, requestInterval, timeout, s.errorCount)

	s.lastReportTime = time.Now()
	s.lastSuccessCount = s.successCount
}

// GetErrorCount returns the current error count.
func (s *Stats) GetErrorCount() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.errorCount
}

// GetSuccessCount returns the current success count.
func (s *Stats) GetSuccessCount() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.successCount
}

// GetElapsedTime returns the elapsed time since the attack started.
func (s *Stats) GetElapsedTime() float64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return time.Since(s.startTime).Seconds()
}


