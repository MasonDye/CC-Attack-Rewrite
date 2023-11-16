![CC Attack Rewrite](https://github.com/MasonDye/CC-Attack-Rewrite/blob/95728b220dc90c762ce9034904ec9489037854b7/img/CCAttack%2B%2BGo128.png)
# CC Attack ++ Rewrite
![](https://img.shields.io/badge/Build-success-green) ![](https://img.shields.io/badge/Version-2.2.0-orange) ![](https://img.shields.io/badge/Author-MasonDye-blue)
![CC Attack Rewrite preview](https://github.com/MasonDye/CC-Attack-Rewrite/blob/main/img/Preview.png)
:-:
Next Generation CC Attack Tool ✨
✨ Threads ✨ HTTP Proxy ✨ Asynchronous ✨ Multi-threaded✨

## What is CC Attack ++ Rewrite?
CC Attack ++ Rewrite is a CC attack program that has been rewritten in GoLang based on its predecessor project, CC Attack ++.

## What can it do?
Test website firewall, DDoS CC protection; test network performance, maximum network request load.

</div>

## How to use CC Attack ++ Rewrite?
Full command:
<pre><code>./cca -url=http://localhost -speed=100 -thread=8 -timeout=2500 -ua=ua.txt -ip=ip-pool.txt</code></pre>

### Parameter description:

[URL] need to specify HTTP/HTTPS.
<pre><code>-url string</code></pre>

[IP Pool] IP pool path (relative path) (.txt).
<pre><code>-ip string</code></pre>

[Thread] Number of threads (default 2)
<pre><code>-thread int</code></pre>

[Speed] Attack Speed(ms) (default 100)
<pre><code>-speed int</code></pre>

[UA] Set User-Agent Config Path (relative path) (.txt)
<pre><code>-ua string</code></pre>

[Timeout] Timeout(ms) (default 1000)
<pre><code>-timeout int</code></pre>

### Format: 
ip pool (eg. ip-list.txt)
<pre><code>http://address:port
http://address:port
http://address:port
......</code></pre>

ua pool (eg. ua-list.txt)
<pre><code>Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/22.0.1207.1 Safari/537.1
Mozilla/5.0 (X11; CrOS i686 2268.111.0) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.57 Safari/536.11
Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6
......</code></pre>

### Usage:
<pre><code>Usage:
  -ip string
        IP Pool Path (txt)
  -thread int
        thread (default 2)
  -time int
        Attack Speed(ms) (default 100)
  -ua string
        User-Agent Pool Path (txt)
  -url string
        Attack URL
  -timeout int
        Timeout(ms) (default 2500)</code></pre>

## How to Build?
<pre><code>go build CC-Attack-Rewrite.go</code></pre>

## Stargazers over time
[![Stargazers over time](https://starchart.cc/MasonDye/CC-Attack-Rewrite.svg)](https://starchart.cc/MasonDye/CC-Attack-Rewrite)
