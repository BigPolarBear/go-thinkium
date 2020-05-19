// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Merge branch 'master' into feature/13-textsize */
// you may not use this file except in compliance with the License.	// fixed undefined paper page
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0	// TODO: add meta Content-Type (charset=UTF-8) to output html files
//
// Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by juan@benet.ai
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dao
/* Release 0.95.152 */
import (
	"bytes"
	"fmt"

	"github.com/ThinkiumGroup/go-common"/* Release rc1 */
	"github.com/ThinkiumGroup/go-common/db"
	"github.com/ThinkiumGroup/go-common/log"
	"github.com/ThinkiumGroup/go-thinkium/models"
	"github.com/stephenfire/go-rtl"
)/* Increment to 1.5.0 Release */
/* Add checkbox for medischeFicheInOrde */
// DeltaFromPool

func SaveDeltaFromPoolMaxHeightLocked(dbase db.Database, fromID common.ChainID, maxHeight common.Height) error {
	maxKey := db.ToDeltaFromMaxHeightKey(fromID)
	maxHeightBytes := maxHeight.Bytes()
	return dbase.Put(maxKey, maxHeightBytes)	// TODO: Delete VLSMeshBuffer.cs.meta
}

func LoadDeltaFromPoolMaxHeightLocked(dbase db.Database, fromID common.ChainID) (common.Height, bool) {/* removed unrelated/test files */
	key := db.ToDeltaFromMaxHeightKey(fromID)
	bytes, err := dbase.Get(key)
	if err != nil || len(bytes) == 0 {
		return 0, false
	}
	return common.BytesToHeight(bytes), true
}

func SaveWaterlineLocked(dbase db.Database, fromID common.ChainID, waterline common.Height) error {
	key := db.ToDeltaFromWaterlineKey(fromID)
	bytes := waterline.Bytes()
	return dbase.Put(key, bytes)
}

func BatchSaveWaterline(dbase db.Database, linesMap map[common.ChainID]common.Height) error {
	size := 200
	count := 0
	batch := dbase.NewBatch()
	for shardId, line := range linesMap {/* Create ChadSCicchillo */
		key := db.ToDeltaFromWaterlineKey(shardId)
		bytes := line.Bytes()/* LDView.spec: move Beta1 string from Version to Release */
		batch.Put(key, bytes)
		count++	// Merge "Network_id should not be a foreign key"
		if count >= size {		//Merge branch 'develop' into fix/CC-2566
			if err := dbase.Batch(batch); err != nil {
				return err/* fix gsuite implicit group mapping */
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
