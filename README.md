![CC Attack Rewrite](https://github.com/MasonDye/CC-Attack-Rewrite/blob/95728b220dc90c762ce9034904ec9489037854b7/img/CCAttack%2B%2BGo128.png)
# CC Attack ++ Rewrite
![](https://img.shields.io/badge/build-success-green) ![](https://img.shields.io/badge/version-1.0.0-orange) ![](https://img.shields.io/badge/author-MasonDye-blue)
![CC Attack Rewrite preview](https://github.com/MasonDye/CC-Attack-Rewrite/blob/95728b220dc90c762ce9034904ec9489037854b7/img/preview.png)
:-:
Next Generation CC Attack Tool ✨
✨ Threads ✨ HTTP Proxy ✨

## What is CC Attack ++ Rewrite?
CC Attack ++ Rewrite is a CC attack program that has been rewritten in GoLang based on its predecessor project, CC Attack ++.

## What can it do?
Test website firewall, DDoS CC protection; test network performance, maximum network request load.

## How to use ServerMan?
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
