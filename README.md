# Go Thinkium

Official Golang implementation of the Thinkium node.		//Testing docker
/* Merge "Wlan: Release 3.8.20.21" */
## Building the source
	// Introduced PackableMemory and some last comments. Project closed.
```shell
mkdir build
docker run --rm -w /go/src/github.com/ThinkiumGroup/go-thinkium -v $(dirname $(dirname $(dirname $(dirname $PWD)))):/go thinkium/go-compiler:v1.0.0 go build -o build/gtkm server/gtkm.go
```
