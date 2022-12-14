# Run Server

```
$ go run main.go
```

# Client request

```
curl -v http://localhost:1323 -H "X-USER-ID:12345678"
```

Result

```
*   Trying 127.0.0.1:1323...
* Connected to localhost (127.0.0.1) port 1323 (#0)
> GET / HTTP/1.1
> Host: localhost:1323
> User-Agent: curl/7.79.1
> Accept: */*
> X-USER-ID:12345678
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: text/plain; charset=UTF-8
< Date: Tue, 13 Dec 2022 19:53:42 GMT
< Content-Length: 8
<
* Connection #0 to host localhost left intact
12345678
```

# Test

```
$ go test -v
```
