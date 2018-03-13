# Simple Restfool JSON API

is a stupidly "foolish" and simple approach of implementing a JSON Restful API library.

## Features
- path routing using gorilla mux
- versioning
- database wrapper
- TLS
- pretty print
- Etag / If-None-Match Clientside caching
- rate limiting and headers using trottled middleware
- basic auth
- config using TOML format
- error handler
- logging

### Ratelimit Headers
```
X-Ratelimit-Limit - The number of allowed requests in the current period
X-Ratelimit-Remaining - The number of remaining requests in the current period
X-Ratelimit-Reset - The number of seconds left in the current period
```

## tbd
- X-HTTP-Method-Override
- caching serverside (varnish support ?)
- Authentication - oauth(2)

## simple example
tbd

## config example 
tbd
Dont fool the reaper ?

