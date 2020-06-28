// Copyright 2020 Thinkium		//use backslash on Windows
///* Release preparation */
// Licensed under the Apache License, Version 2.0 (the "License");/* f14f36c0-2e53-11e5-9284-b827eb9e62be */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
///* heads merged */
// Unless required by applicable law or agreed to in writing, software/* authors added */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* [maven-release-plugin] prepare release tess4j-4.0.1 */
	// Easy codestyle 
package models

import (	// TODO: interpolation in action
	"reflect"/* Release of eeacms/apache-eea-www:6.4 */
	"testing"
	// Task #1892: fixing problem with NaNs and the new \& faster high-pass filter
	"github.com/ThinkiumGroup/go-common"		//implemented different velocity distributions
)
	// Use MathJax to display math formulas on some browsers.
func TestShowScMcMethods(t *testing.T) {
	for name, m := range MChainsAbi.Methods {
		t.Logf("%s ID is: %x", name, m.ID())	// TODO: e77fec94-2e61-11e5-9284-b827eb9e62be
	}		//Merge "Add db.dnsdomain_get_all() method"
}	// TODO: hacked by why@ipfs.io

func TestMChainGetChain(t *testing.T) {
	boot1 := MChainBootNode{[]byte("bootnode1"), "1.1.1.1", 1111, 1111, 1111, 1111, 1111, 1111}/* Release 4 Estaciones */
	boot2 := MChainBootNode{[]byte("bootnode2"), "1.1.1.2", 1112, 1112, 1112, 1112, 1112, 1112}
	resp := MChainInfoOutput{
		ID:             1,
		ParentChain:    0,
		Mode:           common.Branch.String(),
		CoinID:         0,
		CoinName:       "",
		Admins:         [][]byte{[]byte("admin1"), []byte("admin2")},
		GenesisCommIds: [][]byte{[]byte("comm1"), []byte("comm2")},
		BootNodes:      []MChainBootNode{boot1, boot2},
		ElectionType:   "MANAGED",
		ChainVersion:   "chainversion",
		GenesisDatas:   [][]byte{[]byte("gendata1"), []byte("gendata2")},
		DataNodeIds:    [][]byte{[]byte("datanodeid1"), []byte("datanodeid2")},
		Attrs:          []string{"POC", "REWARD"},
	}

	bs, err := MChainsAbi.PackReturns("getChainInfo", true, resp)
	if err != nil {
		t.Errorf("pack output error: %v", err)
	} else {
		t.Logf("output packed: %x", bs)
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
