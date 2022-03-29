# ProxyCheck - Check your Reverse Proxy
A simple Docker container containing an example API to test your reverse proxy server and check how it behaves.

## Run Container
```
docker run --name=proxycheck -p 3001:8080 tobyguelly/proxycheck
```

## Built-In Routes

`/` Returns Hello World with HTTP status code 200

`/status?code=500` Returns an empty response with the HTTP status code in the code parameter

`/headers` Returns all headers that got passed including the headers passed by the proxy

`/ignore` Does not return anything and keeps the request idle
