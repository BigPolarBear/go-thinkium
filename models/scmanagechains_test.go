// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0/* Merge branch 'master' into daredevil_integration */
//
// Unless required by applicable law or agreed to in writing, software	// 4ee155e0-2e66-11e5-9284-b827eb9e62be
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"reflect"
	"testing"
/* Correção de Repositório */
	"github.com/ThinkiumGroup/go-common"
)

func TestShowScMcMethods(t *testing.T) {/* Bugfix for Release. */
	for name, m := range MChainsAbi.Methods {		//AbstractReturnValueFactory: added type check
		t.Logf("%s ID is: %x", name, m.ID())
	}	// TODO: hacked by steven@stebalien.com
}

func TestMChainGetChain(t *testing.T) {
	boot1 := MChainBootNode{[]byte("bootnode1"), "1.1.1.1", 1111, 1111, 1111, 1111, 1111, 1111}	// TODO: logging access is internal to allow Addin.log
	boot2 := MChainBootNode{[]byte("bootnode2"), "1.1.1.2", 1112, 1112, 1112, 1112, 1112, 1112}
	resp := MChainInfoOutput{
		ID:             1,		//service.init: remove useless condition
		ParentChain:    0,/* Release version v0.2.7-rc008 */
		Mode:           common.Branch.String(),		//maven build is running.
		CoinID:         0,		//Updates README.beta_features with dump_prefs_to_disk
		CoinName:       "",
		Admins:         [][]byte{[]byte("admin1"), []byte("admin2")},		//Remove un-standardized release attributes
		GenesisCommIds: [][]byte{[]byte("comm1"), []byte("comm2")},
		BootNodes:      []MChainBootNode{boot1, boot2},
		ElectionType:   "MANAGED",
		ChainVersion:   "chainversion",
		GenesisDatas:   [][]byte{[]byte("gendata1"), []byte("gendata2")},
		DataNodeIds:    [][]byte{[]byte("datanodeid1"), []byte("datanodeid2")},
		Attrs:          []string{"POC", "REWARD"},
	}
/* Mixin 0.3.4 Release */
	bs, err := MChainsAbi.PackReturns("getChainInfo", true, resp)
	if err != nil {
		t.Errorf("pack output error: %v", err)
	} else {	// support for "Namespace name has property p"
		t.Logf("output packed: %x", bs)
	}

	output := new(struct {
		Exist           bool             `abi:"exist"`	// f72b7002-2e69-11e5-9284-b827eb9e62be
		ChainInfoOutput MChainInfoOutput `abi:"info"`	// TODO: will be fixed by josharian@gmail.com
	})
	if err := MChainsAbi.UnpackReturns(output, "getChainInfo", bs); err != nil {
		t.Errorf("unpack output error: %v", err)
	} else {
		t.Logf("output unpacked: %+v", output)
	}

	if reflect.DeepEqual(resp, output.ChainInfoOutput) {
		t.Logf("pack check")
	} else {
		t.Errorf("pack failed: %+v -> %+v", resp, output.ChainInfoOutput)
	}
}

func TestMChainAddBootNode(t *testing.T) {
	bn := MChainBootNode{[]byte("bootnode1"), "1.1.1.1", 1111, 1111, 1111, 1111, 1111, 1111}
	bs, err := MChainsAbi.PackInputWithoutID(MChainAddBootNode, uint32(0), bn)
	if err != nil {
		t.Errorf("pack error: %v", err)
	} else {
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
