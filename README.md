![CC Attack Rewrite](https://github.com/MasonDye/CC-Attack-Rewrite/blob/95728b220dc90c762ce9034904ec9489037854b7/img/CCAttack%2B%2BGo128.png)
# CC Attack ++ Rewrite
![](https://img.shields.io/badge/build-success-green) ![](https://img.shields.io/badge/version-2.0.0-orange) ![](https://img.shields.io/badge/author-MasonDye-blue)
![CC Attack Rewrite preview](https://github.com/MasonDye/CC-Attack-Rewrite/blob/95728b220dc90c762ce9034904ec9489037854b7/img/preview.png)
:-:
Next Generation CC Attack Tool ✨
✨ Threads ✨ HTTP Proxy ✨ Synchronous ✨ multi-threaded✨

## What is CC Attack ++ Rewrite?
CC Attack ++ Rewrite is a CC attack program that has been rewritten in GoLang based on its predecessor project, CC Attack ++.

## What can it do?
Test website firewall, DDoS CC protection; test network performance, maximum network request load.

</div>

## How to use CC Attack ++ Rewrite?
Full command:
<pre><code>./cca -url=http://localhost -time=100 -ua=ua.txt -ip=ip-pool.txt -thread=8</code></pre>

### Parameter description:

[URL] need to specify HTTP/HTTPS.
<pre><code>-url string</code></pre>

[IP Pool] IP pool path (relative path) (.txt).
<pre><code>-ip string</code></pre>

[Thread] Number of threads (default 2)
<pre><code>-thread int</code></pre>

[Time] Attack Speed(ms) (default 100)
<pre><code>-time int</code></pre>

[UA] Set User-Agent Config Path (txt)
<pre><code>-ua string</code></pre>

### Format: 
ip pool (eg. ip.txt)
<code>http://address:port</code>

ua (eg. ua.txt)
<code>Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3</code>

### Usage:
<pre><code>
Usage:
  -ip string
        IP Pool Path (txt)
  -thread int
        thread (default 2)
  -time int
        Attack Speed(ms) (default 100)
  -ua string
        User-Agent Config Path (txt)
  -url string
        Attack URL
</code></pre>

## How to Build?
<pre><code>go build CC-Attack-Rewrite.go</code></pre>

## Stargazers over time
[![Stargazers over time](https://starchart.cc/MasonDye/CC-Attack-Rewrite.svg)](https://starchart.cc/MasonDye/CC-Attack-Rewrite)
