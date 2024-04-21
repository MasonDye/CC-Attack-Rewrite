![CC Attack Rewrite](https://github.com/MasonDye/CC-Attack-Rewrite/blob/95728b220dc90c762ce9034904ec9489037854b7/img/CCAttack%2B%2BGo128.png)
# CC Attack ++ Rewrite
![](https://img.shields.io/badge/Build-success-green) ![](https://img.shields.io/badge/Version-2.3.0-orange) ![](https://img.shields.io/badge/Author-MasonDye-blue)
![CC Attack Rewrite preview](https://github.com/MasonDye/CC-Attack-Rewrite/blob/main/img/Preview.png)
:-:
Next Generation CC Attack Tool ✨
✨ Threads ✨ HTTP Proxy ✨ Asynchronous ✨ Multi-threaded ✨

## What is CC Attack ++ Rewrite?
CC Attack ++ Rewrite is a CC attack program that has been rewritten in GoLang based on its predecessor project, CC Attack ++.

## What can it do?
Test website firewall, DDoS CC protection; test network performance, maximum network request load.

## How to use CC Attack ++ Rewrite?
Full command:
<pre><code>./CC-Attack-Rewrite -url=http://localhost -speed=100 -thread=8 -timeout=2500 -ua=ua-list.txt -ip=ip-list.txt -time=300 -http=1.1 -cookie='test=cookievule;'</code></pre>

### Parameter description:

[URL] need to specify HTTP/HTTPS.
<pre><code>-url string</code></pre>

[IP Pool] IP pool path (relative path) (.txt).
<pre><code>-ip string</code></pre>

[Thread] Number of threads (default 2)
<pre><code>-thread int</code></pre>

[Speed] Attack Speed(ms) (default 100)
<pre><code>-speed int</code></pre>

[UA] Set User-Agent pool Path (relative path) (.txt)
<pre><code>-ua string</code></pre>

[Timeout] Timeout(ms) (default 1000)
<pre><code>-timeout int</code></pre>

[Cookie] Cookie to include in request (default NULL)
<pre><code>-cookie string</code></pre>

[Time] Attack Time (seconds) (default NULL)
<pre><code>-time int</code></pre>

[Http] HTTP version (1.1 or 2.0) (default 1.1)
<pre><code>-http string</code></pre>

### Format: 
ip pool (eg. ip-list.txt)
<pre><code>address:port
address:port
address:port
......</code></pre>

ua pool (eg. ua-list.txt)
<pre><code>Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/22.0.1207.1 Safari/537.1
Mozilla/5.0 (X11; CrOS i686 2268.111.0) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.57 Safari/536.11
Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6
......</code></pre>

### Usage:
<pre><code>Usage:
  -cookie
        Cookie to include in request
  -http
        HTTP version (1.1 or 2.0)
  -ip
        IP Pool Path (txt)
  -speed
        Attack Speed(ms)
  -thread
        thread
  -time
        Attack Time (seconds)
  -timeout
        Request Timeout (ms)
  -ua
        User-Agent Pool Path (txt)
  -url
        Attack URL</code></pre>

## How to Build?
<pre><code>go build CC-Attack-Rewrite.go</code></pre>

## Stargazers over time
[![Stargazers over time](https://starchart.cc/MasonDye/CC-Attack-Rewrite.svg)](https://starchart.cc/MasonDye/CC-Attack-Rewrite)
