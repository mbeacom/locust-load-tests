Locust-Load-Tests
================

Purpose
-------
Spike using [Locust](https://locust.io/) with a [Go driver](https://github.com/myzhan/boomer)

Setup
-----

### ZeroMQ
```
$ pip install pyzmq 
```

### Locust
```
$ pip install locustio
```

### Go
```
$ brew update
$ brew install go
```

Explicitly set $GOPATH in a shell initialisation file.
```
$ export GOPATH=$HOME/go
$ export PATH=$PATH:$HOME/go/bin
```

### Boomer
```
$ go get github.com/myzhan/boomer
```

Dummy Master
------------
A dummy master script needs to be invoked in order for a custom slave to send
metrics to it.
```
$ locust -f dummy.py --master --master-bind-host=127.0.0.1 --master-bind-port=5557
```

Example
------
We are going to run a local Go server and bash it with the Go driver

Change directory to **server**.
```
$ go run server.go
```

In **examples/urlbash**
```
$ go build -o a.out main.go
```

Execute and Attach to Master
----------------------------
```
$ ./a.out --master-host=127.0.0.1 --master-port=5557 --rpc=zeromq
```

Test
```
$ open http://localhost:8089
```
Set **10** users with a hatch rate of *1* and watch the requests swarm.

Requests per second
-------------------
If you would like to set requests per second add a one second delay.

```go
import "time"

time.Sleep(1 * time.Second)

```
