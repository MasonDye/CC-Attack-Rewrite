![CC Attack 重写](https://github.com/MasonDye/CC-Attack-Rewrite/blob/95728b220dc90c762ce9034904ec9489037854b7/img/CCAttack%2B%2BGo128.png)
# CC Attack ++ 重写
![](https://img.shields.io/badge/Build-success-green) ![](https://img.shields.io/badge/Version-2.4.0-orange) ![](https://img.shields.io/badge/Author-MasonDye-blue)
![CC Attack 重写 预览](https://github.com/MasonDye/CC-Attack-Rewrite/blob/main/img/Preview.png)
:-:
下一代 CC 压力测试工具 ✨
✨ 多线程 ✨ HTTP 代理 ✨ 异步 ✨

## 什么是 CC Attack ++ 重写？
CC Attack ++ 重写是在其前身项目 CC Attack ++ 的基础上用 GoLang 重写的 CC 攻击程序。

## 它可以做什么？
测试网站防火墙, DDoS CC 防护; 测试网络性能, 最大网络请求负载.

## 如何使用 CC Attack ++ 重写？
完整命令:
<pre><code>./CC-Attack-Rewrite -url=http://localhost -speed=100 -thread=8 -timeout=2500 -ua_pool=ua-list.txt -ip_pool=ip-list.txt -time=300 -http_version=1.1 -http_methods 'GET' -cookie='test=cookievalue;'</code></pre>

### 参数说明:

[URL] 必须为 HTTP/HTTPS.
<pre><code>-url string</code></pre>

[IP Pool] IP 池路径 (相对路径) (.txt).
<pre><code>-ip_pool string</code></pre>

[Thread] 线程数 (默认 2)
<pre><code>-thread int</code></pre>

[Speed] 攻击速度(ms) (默认 100)
<pre><code>-speed int</code></pre>

[UA Pool] 设置 User-Agent 池路径 (相对路径) (.txt)
<pre><code>-ua_pool string</code></pre>

[Timeout] 超时时间(ms) (默认 1000)
<pre><code>-timeout int</code></pre>

[Cookie] 附加在请求中的Cookie (默认 NULL)
<pre><code>-cookie string</code></pre>

[Time] 攻击时间 (seconds) (默认 NULL)
<pre><code>-time int</code></pre>

[Http-Version] HTTP 版本 (1.1 或 2.0) (默认 1.1)
<pre><code>-http_version string</code></pre>

[Http-Method] HTTP 请求方法 (GET, POST, PUT, DELETE, HEAD, OPTIONS, PATCH, TRACE, CONNECT) (默认 GET)
<pre><code>-http_methods string</code></pre>

### 格式: 
ip 池 (eg. ip-list.txt)
<pre><code>address:port
address:port
address:port
......</code></pre>

ua 池 (eg. ua-list.txt)
<pre><code>Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/22.0.1207.1 Safari/537.1
Mozilla/5.0 (X11; CrOS i686 2268.111.0) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.57 Safari/536.11
Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6
......</code></pre>

### 使用:
<pre><code>使用:
  -cookie
        附加在请求中的Cookie
  -http_methods
        HTTP 请求方法 (GET, POST, PUT, DELETE, HEAD, OPTIONS, PATCH, TRACE, CONNECT)
  -http_version
        HTTP 版本 (1.1 or 2.0)
  -ip_pool
        IP 池路径 (txt)
  -speed
        攻击速度(ms)
  -thread
        线程数
  -time
        攻击时间 (seconds)
  -timeout
        请求超时时间 (ms)
  -ua_pool
        User-Agent 池路径 (txt)
  -url
        攻击地址</pre></code>

## 如何编译?
<pre><code>go build CC-Attack-Rewrite.go</code></pre>

## 为什么不更新？
出于我的学业和个人原因，更新频率会慢很多，如果您急迫需要更新，可以提交pull request，或者创建一个issuse等待我有时间的时候更新。

## Stargazers over time
[![Stargazers over time](https://starchart.cc/MasonDye/CC-Attack-Rewrite.svg)](https://starchart.cc/MasonDye/CC-Attack-Rewrite)
