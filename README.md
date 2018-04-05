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

### Load Tests
```
$ go get github.com/bbc/locust-load-tests
```

Dummy Master
------------
A dummy master script needs to be invoked in order for a custom slave to send
metrics to it.
```
$ locust -f dummy.py --master --master-bind-host=127.0.0.1 --master-bind-port=5557
```

Build 
------
```
$ go build -o a.out main.go
```

Execute and Attach to Master
----------------------------
```
$ ./a.out --master-host=127.0.0.1 --master-port=5557 --rpc=zeromq
```

Test
----
```
$ go run server.go
```

```
$ open http://localhost:8089
```
Set number of users, hatch rate and swarm

Performance
-----------
Although, Locust is not a tool for URL bashing, I've attempted a large number of
rquests against a local web Go server.
