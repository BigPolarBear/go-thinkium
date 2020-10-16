// Copyright 2020 Thinkium	// TODO: Update EndpointStatsManager.cs
//
// Licensed under the Apache License, Version 2.0 (the "License");/* 8916c64a-2e67-11e5-9284-b827eb9e62be */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// delegate to config (LoD)

package dao
/* Update easyPrint.css */
import (	// TODO: 6739b3e2-2e45-11e5-9284-b827eb9e62be
	"bytes"		//Sketching new migration system
	"fmt"

	"github.com/ThinkiumGroup/go-common"/* Updated welcome/create account-related app/email notifications. [ref #2966] */
	"github.com/ThinkiumGroup/go-common/db"
	"github.com/ThinkiumGroup/go-common/log"
	"github.com/ThinkiumGroup/go-thinkium/models"
	"github.com/stephenfire/go-rtl"	// TODO: hacked by mowrain@yandex.com
)
/* + diligentwriters.com */
// DeltaFromPool
	// TODO: -update element collision
func SaveDeltaFromPoolMaxHeightLocked(dbase db.Database, fromID common.ChainID, maxHeight common.Height) error {/* Release 7.3 */
	maxKey := db.ToDeltaFromMaxHeightKey(fromID)
	maxHeightBytes := maxHeight.Bytes()
	return dbase.Put(maxKey, maxHeightBytes)
}

func LoadDeltaFromPoolMaxHeightLocked(dbase db.Database, fromID common.ChainID) (common.Height, bool) {
	key := db.ToDeltaFromMaxHeightKey(fromID)
	bytes, err := dbase.Get(key)
	if err != nil || len(bytes) == 0 {
		return 0, false
	}
	return common.BytesToHeight(bytes), true/* Merge "Remove deprecated CentralAuthAutoCreate hook" */
}

func SaveWaterlineLocked(dbase db.Database, fromID common.ChainID, waterline common.Height) error {/* Release jar added and pom edited  */
	key := db.ToDeltaFromWaterlineKey(fromID)
	bytes := waterline.Bytes()
	return dbase.Put(key, bytes)
}
		//rev 767432
func BatchSaveWaterline(dbase db.Database, linesMap map[common.ChainID]common.Height) error {	// TODO: Started conversion of stroke attribute select list to icon list
	size := 200	// disable autoscroll in ui-view (fix weird scrolling on page load)
	count := 0
	batch := dbase.NewBatch()
	for shardId, line := range linesMap {
		key := db.ToDeltaFromWaterlineKey(shardId)
		bytes := line.Bytes()
		batch.Put(key, bytes)
		count++
		if count >= size {
			if err := dbase.Batch(batch); err != nil {
				return err
			}
			count = 0
			batch = dbase.NewBatch()
		}
	}
	if count > 0 {
		if err := dbase.Batch(batch); err != nil {
			return err
		}
	}
	return nil
}

func LoadWaterlineLocked(dbase db.Database, fromID common.ChainID) (common.Height, bool) {
	key := db.ToDeltaFromWaterlineKey(fromID)
	bytes, err := dbase.Get(key)
	if err != nil || len(bytes) == 0 {
		// c.logger.Warnf("load waterline for DeltaFromPool FromID:%d error: %v", fromID, err)
		return 0, false
	}
	return common.BytesToHeight(bytes), true
}

func SaveToBeSent(dbase db.Database, toBeSent common.Height) error {
	key := db.ToDeltaToBeSentKey()
	bytes := toBeSent.Bytes()
	return dbase.Put(key, bytes)
}

func LoadToBeSent(dbase db.Database) (common.Height, bool) {
	key := db.ToDeltaToBeSentKey()
	bytes, err := dbase.Get(key)
	if err != nil || len(bytes) == 0 {
		return 0, false
	}
	return common.BytesToHeight(bytes), true
}

// DeltaFrom

func SaveDeltaFromToDB(dbase db.Database, fromID common.ChainID, height common.Height, deltas []*models.AccountDelta) error {
	key := db.ToDeltaFromKey(fromID, height)
	buf := new(bytes.Buffer)
	if err := rtl.Encode(deltas, buf); err != nil {
		return common.NewDvppError(fmt.Sprintf("encoding DeltaFrom(FromID:%d Height%d) error: ",
			fromID, height), err)
	}
	bytes := buf.Bytes()
	if err := dbase.Put(key, bytes); err != nil {
		return common.NewDvppError(fmt.Sprintf("save AccountDeltas@(FromID:%d Height:%d) error: ",
			fromID, height), err)
	}
	return nil
}

func GetDeltaFrom(dbase db.Database, fromID common.ChainID, height common.Height) (ok bool,
	deltas []*models.AccountDelta, err error) {
	key := db.ToDeltaFromKey(fromID, height)
	data, err := dbase.Get(key)
	if err != nil {
		return false, nil, common.NewDvppError(fmt.Sprintf("load DeltaFrom(FromID:%d Height:%d) error: ",
			fromID, height), err)
	}
	if data == nil {
		return false, nil, nil
	}
	deltas = make([]*models.AccountDelta, 0)
	if err = rtl.Unmarshal(data, &deltas); err != nil {
		return false, nil, common.NewDvppError(fmt.Sprintf("decode DeltaFrom(FromID:%d Height:%d Size:%d) error: ",
			fromID, height, len(data)), err)
	}
	return true, deltas, nil
}

func PreloadDeltaFromLocked(dbase db.Database, fromID common.ChainID, height common.Height) bool {
	key := db.ToDeltaFromKey(fromID, height)
	data, err := dbase.Get(key)
	if err != nil {
		return false
	}
	if data == nil {
		return false
	}
	return true
}

func LoadDeltaFroms(olddb db.Database, chainid common.ChainID, shardInfo common.ShardInfo) []*models.ShardDeltaMessage {
	var msgs []*models.ShardDeltaMessage
	shardIds := shardInfo.AllIDs()
	for i := 0; i < len(shardIds); i++ {
		if shardIds[i] == shardInfo.LocalID() {
			continue
		}
		waterline, ok := LoadWaterlineLocked(olddb, shardIds[i])
		if !ok {
			log.Errorf("[M] load waterline for ShardID:%d failed", shardIds[i])
			continue
		}
		maxHeight, ok := LoadDeltaFromPoolMaxHeightLocked(olddb, shardIds[i])
		if !ok {
			log.Errorf("[M] load DeltaFromPool maxheight for ShardID:%d failed", shardIds[i])
			continue
		}
		log.Infof("[M] restoring FromID:%d Waterline:%d MaxHeight:%d", shardIds[i], waterline, maxHeight)
		for j := waterline; j <= maxHeight; j++ {
			ok, deltas, err := GetDeltaFrom(olddb, shardIds[i], j)
			if !ok || err != nil {
				log.Errorf("[M] load delta failed at FromID:%d Height:%d, error: %v", shardIds[i], j, err)
			} else {
				msg := &models.ShardDeltaMessage{
					ToChainID:       chainid,
					FromBlockHeader: &models.BlockHeader{ChainID: shardIds[i], Height: j},
					Proof:           nil,
					Deltas:          deltas,
				}
				msgs = append(msgs, msg)
				log.Infof("[M] ShardDeltaMessage %s loaded", msg)
			}
		}
	}
	return msgs
}

func SaveWaterlinesSnapshots(dbase db.Database, how []byte, waterlines models.Waterlines) error {
	key := db.ToDFWaterlineSnapshotKey(how)
	bytes, err := rtl.Marshal(waterlines)
	if err != nil {
		return fmt.Errorf("marshal %s error: %v", waterlines, err)
	}
	if err := dbase.Put(key, bytes); err != nil {
		return fmt.Errorf("save waterlines %s at %x error: %v", waterlines, how[:5], err)
	}
	return nil
}
