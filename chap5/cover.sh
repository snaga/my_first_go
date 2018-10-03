env GOPATH=`pwd` go test -cover cart
env GOPATH=`pwd` go test -coverprofile=cover.out cart
env GOPATH=`pwd` go tool cover -func=cover.out
env GOPATH=`pwd` go tool cover -html=cover.out
