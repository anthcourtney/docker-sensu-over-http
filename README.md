# docker sensu-over-http

## Background

This is a quick and dirty PoC of providing a HTTP interface to sensu-plugins for the purpose of performing health monitoring of an EC2 instance.

In other words:

* We have deployed some internal NTP servers.
* We haven't setup any custom metric -> CloudWatch for NTP time sync.
* We like sensu, but we don't use it elsewhere yet.
* We like sensu's plugins, because we know they work.
* Sensu has a nice 'check-ntp' plugin which will check the offset of the local ntpd process against upstream servers.
* We want to expose this plugin over HTTP, so we can query it (i.e. remotely) and perform some sort of action.

This is a little contrived, and the example is very hacky, but it's a quick and dirty PoC.

## Build and Run

```
$ docker build -t ac:sensu-over-http ./
$ docker run -it -p 127.0.0.1:8080:8080 ac:sensu-over-http
```
## Test

Connect to the docker container on port 8080 and make a HTTP request.

If ntpd is in sync, a 200 OK will be returned .

i.e.

```
$ curl -v http://localhost:8080/ntp
> GET /ntp HTTP/1.1
> User-Agent: curl/7.38.0
> Host: localhost:8080
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Wed, 07 Sep 2016 06:08:03 GMT
< Content-Length: 35
< Content-Type: text/plain; charset=utf-8
< 
CheckNTP OK: NTP offset by 2.085ms
```

If ntpd is not in sync or there's an unknown status, a 503 Service Unavailable will be returned.

```
$ curl -v http://localhost:8080/ntp
> GET /ntp HTTP/1.1
> User-Agent: curl/7.38.0
> Host: localhost:8080
> Accept: */*
> 
< HTTP/1.1 503 Service Unavailable
< Content-Type: text/plain; charset=utf-8
< X-Content-Type-Options: nosniff
< Date: Wed, 07 Sep 2016 06:07:58 GMT
< Content-Length: 75
< 
CheckNTP UNKNOWN: NTP command Failed
```

In both cases, the body of the response will contain the output of the sensu check.

## Notes

* CentOS was used because alpine linux doesn't include ntpq in any packages. It's kind of essential to the sensu plugin.
* We could add more checks, but again, it's a PoC. 
* Similarly, supported HTTP methods/schemes could be tightened up. 
* Stuff is hardcoded.
* This is all a little contrived.
