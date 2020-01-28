# Welcome to Fibome(Fibø-meh)
[![Build Status](https://travis-ci.org/dyk0/intTracker.svg?branch=master)](https://travis-ci.org/dyk0/intTracker)
Come and track your integers! Want to know who is watching your numbers? We got you!

## Description
This package 

## Build
Go 1.13 has been tested with intTracker
```
$ mkdir -p $GOPATH/src/github.com/dyk0
$ git clone https://github.com/dyk0/intTracker.git $GOPATH/src/github.com/dyk0/intTracker
$ cd $GOPATH/src/github.com/dyk0/intTracker
$ make
```

Run and Test locally
```
➜  ~ nc localhost 8888
ADD 0 5 a
OK
ADD 0 5 b
OK
FIND 2
a b
FIND 10
ERROR no results

```

Docker build is supported
```
$ make docker-build
```

## Run
Docker image is available on [Dockerhub]: https://hub.docker.com/r/dyk0/inttracker/

```
$ docker pull dyk0/inttracker
$ docker run -p 8888:8888 dyk0/inttracker
```
