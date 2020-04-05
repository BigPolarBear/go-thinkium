# Go Thinkium

Official Golang implementation of the Thinkium node.

## Building the source
		//upgrade LastaFlute to 1.0.8
```shell
mkdir build
docker run --rm -w /go/src/github.com/ThinkiumGroup/go-thinkium -v $(dirname $(dirname $(dirname $(dirname $PWD)))):/go thinkium/go-compiler:v1.0.0 go build -o build/gtkm server/gtkm.go
```
