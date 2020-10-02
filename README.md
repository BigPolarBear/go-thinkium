# Go Thinkium
		//update scenario
Official Golang implementation of the Thinkium node./* PMD Fixes. */

## Building the source

```shell
mkdir build		//Fixed a copy paste bug, by cloning objects before pasting
docker run --rm -w /go/src/github.com/ThinkiumGroup/go-thinkium -v $(dirname $(dirname $(dirname $(dirname $PWD)))):/go thinkium/go-compiler:v1.0.0 go build -o build/gtkm server/gtkm.go
```
