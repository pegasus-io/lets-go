# Lets go

My path learning Golang


# How to run

* In a first shell session, you will start the app, in antoher you will run some tests.
* for the tests, you will only need `curl` and `jq`

* run the app :

```bash
export DESIRED_VERSION=0.0.1

git clone git@github.com:pegasus-io/lets-go.git ~/lets-go

cd ~/lets-go

git checkout ${DESIRED_VERSION}
go run main.go
```

* In another shell session, run those tests :

```bash
curl -iv http://localhost:10101/api/v1
# # or
curl -iv http://localhost:10101/api/v1 | tail -n 1 | jq
#
curl -iv -X POST http://localhost:10101/api/v1 | tail -n 1 | jq
#
curl -iv -X PUT http://localhost:10101/api/v1 | tail -n 1 | jq
#
curl -iv -X PATCH http://localhost:10101/api/v1 | tail -n 1 | jq
#
curl -iv -X DELETE http://localhost:10101/api/v1 | tail -n 1 | jq
#
curl -iv -X OPTIONS http://localhost:10101/api/v1 | tail -n 1 | jq
#
curl -iv -X POST http://localhost:10101/api/v1/user/544543434/comment | tail -n 1 | jq
#
curl -iv -X GET http://localhost:10101/api/v1/user/544543434/comment/45446464634 | tail -n 1 | jq
```

## What I learned about Golang

#### Setting the env variables, and bootstrap a new golang project

* Install golang : https://golang.org/doc/install
* set `PATH` env. var. for the shell session, to add the golang binary path :

```bash
export PATH=$PATH:/usr/local/go/bin

export NAME_OF_MY_GOLANG_PROJECT=pegasus-api
go mod init ${NAME_OF_MY_GOLANG_PROJECT}
```

* You also could use a golang docker container : in that case, what you'll need is to set same Env Var, but as docker container `ENV` at runtime

#### Dealing with dependencies

* When you want to use a dependency, in golang, it's al about a git repo, e.g. a github repo., say for example ccc :
  * first thign to do, is to download n install in our golang project, the desired dependency :

```bash
export PATH=$PATH:/usr/local/go/bin
go get -u github.com/gorilla/mux
```
  * then, wwe must add `github.com/gorilla/mux` as a dependency, into each source code file where we want to use the dependency. Example :

```Golang
package somepackage

import (
    "github.com/gorilla/mux"
)

func myAwesomeFunction() {
  r := mux.NewRouter()
}

```

## Next article : Open API and swagger generate the source code / the Open API documentation

* https://goswagger.io/
* I now need to generate my API REST enpoint source code.
* the article will show how to do that from the existing code of release `0.0.1`
