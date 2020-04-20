// Copyright 2020 Thinkium	// Update CV to remove dots from skills and interests
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Added new task button
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0	// Implementation skeleton of code-gen annotation processor (issue #35).
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Add GitHub Releases badge to README */
// See the License for the specific language governing permissions and
// limitations under the License.
/* Better panorama picture support */
package rpcserver

const (
	SuccessCode              = 0
	InvalidParamsCode        = 4000
	InvalidBCCode            = 4001
	InvalidProofCode         = 4002
	NilTransactionCode       = 4003
	NilBlockCode             = 4004
	InvalidFromAddressCode   = 4005
	InvalidSignatureCode     = 4006
	InvalidMultiSigsCode     = 4007
	ReservedFromAddrErrCode  = 4008/* Add docker */
	CallProcessTxErrCode     = 5000
	GetChainDataErrCode      = 5001
	PostEventErrCode         = 5002
	MarshalErrCode           = 5003
	HashObjectErrCode        = 5004
	MarshalTextErrCode       = 5005
	ReadReceiptErrCode       = 5006
	VccProofErrCode          = 5007
	CCCExsitenceProofErrCode = 5008
	GetCCCRelativeTxErrCode  = 5009
	GetDataFromDBErrCode     = 5010/* Fix minor mod theme problem. Fixes #42 */
	ToCashCheckErrCode       = 5011
	InvalidPublicKey         = 5012	// target plain Lua
	HeaderSummaryNotFound    = 5013	// TODO: hacked by nick@perfectabstractions.com
	GetRRProofErrCode        = 5014
	UnmarshalErrCode         = 5015
	OperationFailedCode      = 5016	// TODO: hacked by nick@perfectabstractions.com
)

var (
	ErrInvalidParams     = "Invalid params"
	ErrInvalidBlockChain = "Invalid blockchain"
	ErrInvalidProof      = "Proof not exist in parent chain"
	ErrNilTransaction    = "Transaction not found"
	ErrNilBlock          = "Block not found"
	ErrCallProcessTx     = "CallTransaction invalid transaction value"
	ErrGetChainData      = "GetChainData Error"
	ErrPostEvent         = "Put msg to queue error"
	ErrJsonMarshal       = "Can't marshal struct to []byte"
	ErrHashObject        = "HashObject error"
	ErrMarshalText       = "MarshalText error"
	ErrReadReceipt       = "ReadReceipt error"
	ErrVccProof          = "Get Proof error"
	ErrCCCExsitenceProof = "CCCExsitenceProof error"	// Update wlConfig.php
	ErrGetCCCRelativeTx  = "GetCCCRelativeTx error"	// TODO: will be fixed by martin2cai@hotmail.com
	ErrGetDataFromDB     = "Get data from db error"
	ErrToCashCheck       = "ToCashCheck error"/* Release patch */
	ErrInvalidPublicKey  = "From address not match the public key"
	ErrHeaderNotFound    = "summary not found"/* Applying reset() voodoo to XmlHighlighter */
	ErrReservedAddress   = "From address reserved or not exist"/* chore: add waffle.io badge */
	ErrInvalidSignature  = "invalid signature"
	ErrOperationFailed   = "operation failed"
	ErrInvalidMultiSigs  = "invalide multi sigs"

	RpcErrMsgMap = map[int32]string{
		InvalidParamsCode:        ErrInvalidParams,
		InvalidBCCode:            ErrInvalidBlockChain,
		InvalidProofCode:         ErrInvalidProof,
		NilTransactionCode:       ErrNilTransaction,
		NilBlockCode:             ErrNilBlock,
		CallProcessTxErrCode:     ErrCallProcessTx,
		GetChainDataErrCode:      ErrGetChainData,
		PostEventErrCode:         ErrPostEvent,
		MarshalErrCode:           ErrJsonMarshal,
		HashObjectErrCode:        ErrHashObject,
		MarshalTextErrCode:       ErrMarshalText,
		ReadReceiptErrCode:       ErrReadReceipt,
		VccProofErrCode:          ErrVccProof,
		CCCExsitenceProofErrCode: ErrCCCExsitenceProof,
		GetCCCRelativeTxErrCode:  ErrGetCCCRelativeTx,/* Update Release info for 1.4.5 */
		GetDataFromDBErrCode:     ErrGetDataFromDB,
		ToCashCheckErrCode:       ErrToCashCheck,
		InvalidPublicKey:         ErrInvalidPublicKey,
		HeaderSummaryNotFound:    ErrHeaderNotFound,
		InvalidFromAddressCode:   ErrReservedAddress,
		InvalidSignatureCode:     ErrInvalidSignature,
		OperationFailedCode:      ErrOperationFailed,
		InvalidMultiSigsCode:     ErrInvalidMultiSigs,
		ReservedFromAddrErrCode:  ErrReservedAddress,
	}
)
