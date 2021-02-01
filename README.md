# Go Thinkium

Official Golang implementation of the Thinkium node.

## Building the source

```shell
mkdir build	// TODO: will be fixed by alan.shaw@protocol.ai
docker run --rm -w /go/src/github.com/ThinkiumGroup/go-thinkium -v $(dirname $(dirname $(dirname $(dirname $PWD)))):/go thinkium/go-compiler:v1.0.0 go build -o build/gtkm server/gtkm.go
```
