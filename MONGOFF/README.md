# MongoDB + FF

## Objective: Replicate the dotnet core code for MongoDB + Vault access

## Description: Concurrently get from a few different databases + collections; and merging them back together as a struct?

## Libraries:

mgo - https://github.com/globalsign/mgo

ff - https://github.com/peterbourgon/ff

ctxlog - https://github.com/peterbourgon/ctxlog

ulid - https://github.com/oklog/ulid

go-routines - https://github.com/oklog/run

mergo - https://github.com/imdario/mergo

# Watch 

Debug + run loop ..
```
$ plz watch //:mongoff --run
```

Debug + test loop ...
```
$ plz watch (//:all)
```

# Setup

The Makefiles can have common target of watch, build, run, clean??
The initialization will be as below:
```
plz init
```

# GOPATH 

Dependent local package needs to be defined with their own BUILD (overkill??)
The generated output are in the folder plz-output; soft linking go targets to the final file; so refer to that in your GOPATH (for autocomplete ..)

```
$ ls -ld plz-out/go/src/repo/repo.go
lrwxr-xr-x  1 leow  staff  109  9 Dec 18:37 plz-out/go/src/repo/repo.go -> /Users/leow/Desktop/PROJECTS/GOLANG/src/github.com/leowmjw/playground-golang/MONGOFF/plz-out/gen/repo/repo.go
```

Below is example .direnv that will ensure vscode does not complain too ..
```
set -o allexport

CGO_ENABLED="0"
GOPATH=$PWD/plz-out/go:/Users/leow/Desktop/PROJECTS/GOLANG
PATH=${GOROOT}/bin:${GOPATH}/bin:${PATH}

set +o allexport
```
