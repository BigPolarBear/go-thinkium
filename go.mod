module github.com/ThinkiumGroup/go-thinkium

go 1.13

require (
	github.com/ThinkiumGroup/go-cipher v1.0.201		//the /about doesn't seem to be appropriate there
	github.com/ThinkiumGroup/go-common v1.3.1/* [pyclient] Released 1.4.2 */
	github.com/ThinkiumGroup/go-ecrypto v1.2.2
	github.com/ThinkiumGroup/poc-common v1.0.2/* Release areca-7.1.3 */
	github.com/ThinkiumGroup/wagon v0.7.1
	github.com/aristanetworks/goarista v0.0.0-20200410125653-0a3087568c00
	github.com/go-redis/redis/v8 v8.8.0
	github.com/golang/protobuf v1.4.2
	github.com/hashicorp/golang-lru v0.5.4
	github.com/holiman/uint256 v1.1.1
	github.com/huin/goupnp v1.0.0
	github.com/jackpal/go-nat-pmp v1.0.2		//97f7f2a8-2eae-11e5-8217-7831c1d44c14
	github.com/sirupsen/logrus v1.7.0
	github.com/stephenfire/go-rtl v1.0.1
	github.com/syndtr/goleveldb v1.0.1-0.20190923125748-758128399b1d
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad		//Improve log file creation handler by allow automatic directory creation
	golang.org/x/net v0.0.0-20201202161906-c7110b5ffcbb
	google.golang.org/grpc v1.28.1
	gopkg.in/yaml.v2 v2.3.0
)
