# ProxyCheck - Check your Reverse Proxy
A simple Docker image containing a demo API to test your HTTP reverse proxy servers and check how they behave.

## Run Container
```
docker run --name=proxycheck -p 3001:8080 tobyguelly/proxycheck
```

## Built-In Routes

`/` Returns Hello World with HTTP status code 200

`/status?code=500` Returns an empty response with the HTTP status code in the code parameter

`/headers` Returns all headers that got passed including the headers passed by the proxy

`/ignore` Does not return anything and keeps the request idle
