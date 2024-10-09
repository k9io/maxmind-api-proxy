# maxmind-api-proxy

The "maxmind-api-proxy" is a small Golang program that acts as a "proxy" between a "client" and the Maxmind GeoIP
API (https://www.maxmind.com/en/solutions/ip-geolocation-databases-api-services).

As requests are made to the "proxy", results are stored (cached) to a Redis database.  The database will automatically
"expire" entire "cached" entires after a specified amount of time (default: 24 hours).

The idea is to query Maxmind less,  thus saving you money on queries. 

Configuring the "maxmind-api-proxy"
-----------------------------------

In the "etc" directory is the "config.json" file.  This holds the settings that are used by the proxy.  

<pre>
{
  "api_key": "YOUR_PROXY_SERVICE_API_KEY",

  "http_listen": ":8443",
  "http_tls": true,
  "http_cert": "/etc/letsencrypt/live/YOURSITE/fullchain.pem",
  "http_key": "/etc/letsencrypt/live/YOURSITE/privkey.pem",
  "http_mode": "release",

  "maxmind_username":"MAXMIND_USERNAME",
  "maxmind_password":"MAXMIND_PASSWORD",
  "maxmind_url":"https://geoip.maxmind.com/geoip/v2.1/city/",

  "redis_host": "127.0.0.1",
  "redis_port": 6379,
  "redis_password":"YOUR_REDIS_PASSWORD",
  "redis_database": "0",
  "redis_cache_time": 24

}
</pre>

Building "maxmind-api-proxy" and executing the proxy
----------------------------------------------------

<pre>
$ go mod init maxmind-api-proxy
$ go mod tidy
$ go build
$ ./maxmind-api-proxy etc/config.json   # Running the proxy
</pre>

Example query to the proxy
--------------------------

<pre>
curl -H 'API_KEY: YOUR_PROXY_SERVICE_API_KEY' https://your.site:8444/8.8.8.8
</pre>

The proxy will return either a cached or non-cached version of JSON from Maxmind. 


Prebuild Maxmind-API-Proxy binaries
-----------------------------------

If you are unable to access a Golang compiler, you can download pre-built/pre-compiled binaries. These binaries are available for various architectures (i386, amd64, arm64, etc) and multiple operating systems (Linux, Solaris, NetBSD, etc).

You can find those binaries at: https://github.com/k9io/k9-binaries/tree/main/maxmind-api-proxy

You will need a copy of the 'maxmind-api-proxy' configuation file.  That is located at:

https://github.com/k9io/maxmind-api-proxy/blob/main/etc/config.json

