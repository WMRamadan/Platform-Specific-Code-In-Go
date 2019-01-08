# Description
Execute shell commands dependant on platform type compiled in one Go program.

# Builds
To build for the specific platform you are running on:
```console
$ go build main.go
```
## Windows Build
```console
$ GOOS=windows go build main.go
```
## OSX Build
```console
$ GOOS=darwin go build main.go
```
## Linux Build
```console
$ GOOS=linux go build main.go
```
## NetBSD Build
```console
$ GOOS=netbsd go build main.go
```
## OpenBSD Build
```console
$ GOOS=openbsd go build main.go
```
## FreeBSD Build
```console
$ GOOS=freebsd go build main.go
```