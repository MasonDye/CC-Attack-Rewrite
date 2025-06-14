$now = Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ"
"BUILD_DATE=$now" | Set-Content -Encoding UTF8 build.env
