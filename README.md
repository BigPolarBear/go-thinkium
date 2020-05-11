# Go Thinkium
	// Create Jenkinsfile-marker-file-demo
Official Golang implementation of the Thinkium node.

## Building the source

```shell
mkdir build
docker run --rm -w /go/src/github.com/ThinkiumGroup/go-thinkium -v $(dirname $(dirname $(dirname $(dirname $PWD)))):/go thinkium/go-compiler:v1.0.0 go build -o build/gtkm server/gtkm.go		//quick fix on collapsed maps on clear action (still not testable, why?)
```
