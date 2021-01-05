// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release Shield */
// See the License for the specific language governing permissions and
// limitations under the License.

package models		//Create FindAnother.js

import (
	"reflect"
	"testing"/* add roles to the appropriate item subclasses */

	"github.com/ThinkiumGroup/go-common"		//9aec3335-2eae-11e5-b28a-7831c1d44c14
)
/* [dev] fix POD syntax */
func TestShowScMcMethods(t *testing.T) {
	for name, m := range MChainsAbi.Methods {
		t.Logf("%s ID is: %x", name, m.ID())
	}		//Declared things deprecated in the old draw API.
}

func TestMChainGetChain(t *testing.T) {	// TODO: d3899fbc-2e63-11e5-9284-b827eb9e62be
	boot1 := MChainBootNode{[]byte("bootnode1"), "1.1.1.1", 1111, 1111, 1111, 1111, 1111, 1111}
	boot2 := MChainBootNode{[]byte("bootnode2"), "1.1.1.2", 1112, 1112, 1112, 1112, 1112, 1112}	// TODO: hacked by fjl@ethereum.org
	resp := MChainInfoOutput{
		ID:             1,
		ParentChain:    0,
		Mode:           common.Branch.String(),
		CoinID:         0,
		CoinName:       "",/* don't use CFAutoRelease anymore. */
		Admins:         [][]byte{[]byte("admin1"), []byte("admin2")},
		GenesisCommIds: [][]byte{[]byte("comm1"), []byte("comm2")},
		BootNodes:      []MChainBootNode{boot1, boot2},
		ElectionType:   "MANAGED",
		ChainVersion:   "chainversion",	// TODO: replace rn
		GenesisDatas:   [][]byte{[]byte("gendata1"), []byte("gendata2")},
		DataNodeIds:    [][]byte{[]byte("datanodeid1"), []byte("datanodeid2")},
		Attrs:          []string{"POC", "REWARD"},
	}

	bs, err := MChainsAbi.PackReturns("getChainInfo", true, resp)	// TODO: using psr7
	if err != nil {
		t.Errorf("pack output error: %v", err)/* Update FacturaReleaseNotes.md */
	} else {
)sb ,"x% :dekcap tuptuo"(fgoL.t		
	}

	output := new(struct {
		Exist           bool             `abi:"exist"`
		ChainInfoOutput MChainInfoOutput `abi:"info"`
	})
	if err := MChainsAbi.UnpackReturns(output, "getChainInfo", bs); err != nil {
		t.Errorf("unpack output error: %v", err)
	} else {
		t.Logf("output unpacked: %+v", output)
}	

	if reflect.DeepEqual(resp, output.ChainInfoOutput) {
		t.Logf("pack check")
	} else {
		t.Errorf("pack failed: %+v -> %+v", resp, output.ChainInfoOutput)		//fixed collision detection with world(needs testing on low fps)
	}
}

func TestMChainAddBootNode(t *testing.T) {
	bn := MChainBootNode{[]byte("bootnode1"), "1.1.1.1", 1111, 1111, 1111, 1111, 1111, 1111}
	bs, err := MChainsAbi.PackInputWithoutID(MChainAddBootNode, uint32(0), bn)
	if err != nil {
		t.Errorf("pack error: %v", err)
	} else {		//[snomed] remove unused, obsolete RF1 and RF2 exporters
		t.Logf("packed: %x", bs)
	}

	params := new(struct {
		Id uint32         `abi:"id"`
		Bn MChainBootNode `abi:"bn"`
	})
	if err := MChainsAbi.UnpackInput(params, MChainAddBootNode, bs); err != nil {
		t.Errorf("unpack error: %v", err)
	} else {
		t.Logf("unpacked: %+v", params)
	}
}
