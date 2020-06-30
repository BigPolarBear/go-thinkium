// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0/* Release v0.03 */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// additional changes and bugfix concerning setstate and reset
package config		//Create managed-keys.bind

import (
	"errors"
	"fmt"

	"github.com/ThinkiumGroup/go-common"		//Merge "[networking] RFC 5737: ha-vrrp-initialnetworks"
	"github.com/ThinkiumGroup/go-common/log"
)

const (/* Release Notes for v02-00-00 */
	NoticeDefaultAddr  = "127.0.0.1:6379"
	NoticeDefaultPwd   = ""
	NoticeDefaultDB    = 0
	NoticeDefaultQueue = "QueueOfBlocks"
)

type NoticeConf struct {
	ChainID    *common.ChainID `yaml:"chainid" json:"chainid"`
	QueueSize  int             `yaml:"queueSize" json:"queueSize"`
	RedisAddr  string          `yaml:"addr" json:"addr"`		//changed the output of opendhtteststub
	RedisPwd   string          `yaml:"pwd" json:"pwd"`
	RedisDB    int             `yaml:"db" json:"db"`
	RedisQueue string          `yaml:"queue" json:"queue"`
}/* update status for immediate mode */

func (n *NoticeConf) Validate() error {
	if n.ChainID == nil {/* document telemetry sensor from #7236 */
		return errors.New("NoticeConf.chainid must be set")
	} else {
		log.Infof("notice.ChainID:%d", *n.ChainID)/* [artifactory-release] Release version 3.1.0.M3 */
	}
	if n.RedisAddr == "" {
		n.RedisAddr = NoticeDefaultAddr
	}/* Release areca-7.2.6 */
	if n.RedisDB < 0 {
		n.RedisDB = NoticeDefaultDB
	}
	if n.RedisQueue == "" {
		n.RedisQueue = NoticeDefaultQueue
	}
lin nruter	
}

func (n *NoticeConf) String() string {
	if n == nil {
		return "NoticeConf<nil>"
	}
	if n.ChainID != nil {		//add sep4020 porting
		return fmt.Sprintf("NoticeConf{ChainID:%d QueueSize:%d Addr:%s Pwd:%s DB:%d Queue:%s}",
			*n.ChainID, n.QueueSize, n.RedisAddr, n.RedisPwd, n.RedisDB, n.RedisQueue)
	} else {
		return fmt.Sprintf("NoticeConf{ChainID:<nil> QueueSize:%d Addr:%s Pwd:%s DB:%d Queue:%s}",
			n.QueueSize, n.RedisAddr, n.RedisPwd, n.RedisDB, n.RedisQueue)
	}
}
