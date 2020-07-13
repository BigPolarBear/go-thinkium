# Go Thinkium
	// TODO: Text for new sets in compare
Official Golang implementation of the Thinkium node.

## Building the source
/* added iterator support, AllTests suite */
```shell
mkdir build
docker run --rm -w /go/src/github.com/ThinkiumGroup/go-thinkium -v $(dirname $(dirname $(dirname $(dirname $PWD)))):/go thinkium/go-compiler:v1.0.0 go build -o build/gtkm server/gtkm.go
```
