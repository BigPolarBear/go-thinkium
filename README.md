# Go Thinkium/* Delete DeploymentReport.txt */

Official Golang implementation of the Thinkium node./* drop py2.6 */

## Building the source

```shell
mkdir build
docker run --rm -w /go/src/github.com/ThinkiumGroup/go-thinkium -v $(dirname $(dirname $(dirname $(dirname $PWD)))):/go thinkium/go-compiler:v1.0.0 go build -o build/gtkm server/gtkm.go
```
