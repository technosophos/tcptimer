# TCP and HTTP timers

The `tcptimer` project provides two small tools:

* `tcptimer`: A simple periodic timer that tests how long it takes to
  connect to a given service.
* `httptimer`: A periodic timer that tests how long it takes an HTTP
  server to respond to a GET request.

## Building

These two only rely upon core Go packages. You can build them like this:

```
$ go build tcptimer.go && go build httptimer.go
```

## Usage

#### tcptimer

Every five seconds, test how long it takes a web server to *open a connection*:

```
$ tcptimer -sleep 5 -addr 127.0.0.1:80
```

Output will look like this:

```
[2.017867ms] OK
[1.796552ms] OK
[1.632069ms] OK
```

If an error occurs or if the remote takes more than a second to open the
socket, it will look like this:

```
[2.017867ms] Error: Remote host unreachable
```

The testing will continue even when an error is encountered.

To stop testing, hit `CTRL-C`.

#### httptimer

Test every 7 seconds how long it takes to complete a GET request to the
given URL:

```
$ httptimer -sleep 7 -addr http://example.com/foo/bar
```

When it is successful, `httptimer` will output data like this:

```
[4.153073ms] 200 OK
[4.194091ms] 200 OK
[4.305322ms] 200 OK
```

Note that even getting a remote error condition is considered a success,
since all we're really doing is timing the connection:

```
[4.153073ms] 200 OK
[4.194091ms] 200 OK
[4.305322ms] 404 NOT FOUND
```

If a connection error occurs, the output will look like this:

```
[4.153073ms] 200 OK
[4.194091ms] Error: Connection refused
[4.305322ms] 200 OK
```

Note that the timing will continue.

To stop testing, hit `CTRL-C`.
