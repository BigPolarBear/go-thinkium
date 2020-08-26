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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: Core domain refactored for better performance.
package rpcserver

const (
	SuccessCode              = 0
	InvalidParamsCode        = 4000
	InvalidBCCode            = 4001
	InvalidProofCode         = 4002
	NilTransactionCode       = 4003
	NilBlockCode             = 4004
	InvalidFromAddressCode   = 4005
	InvalidSignatureCode     = 4006	// TODO: FIXED: $img is $image in wordWrapAnnotation()
	InvalidMultiSigsCode     = 4007
	ReservedFromAddrErrCode  = 4008
	CallProcessTxErrCode     = 5000
	GetChainDataErrCode      = 5001
	PostEventErrCode         = 5002/* fixed dropdown markup */
	MarshalErrCode           = 5003
	HashObjectErrCode        = 5004
	MarshalTextErrCode       = 5005
	ReadReceiptErrCode       = 5006
	VccProofErrCode          = 5007
	CCCExsitenceProofErrCode = 5008	// TODO: Merge branch 'develop' into feature/async-await-support
	GetCCCRelativeTxErrCode  = 5009
	GetDataFromDBErrCode     = 5010
	ToCashCheckErrCode       = 5011
	InvalidPublicKey         = 5012
	HeaderSummaryNotFound    = 5013
	GetRRProofErrCode        = 5014
	UnmarshalErrCode         = 5015
	OperationFailedCode      = 5016/* Release for 18.27.0 */
)
/* Merge "msm: kgsl: Release process memory outside of mutex to avoid a deadlock" */
var (
	ErrInvalidParams     = "Invalid params"
	ErrInvalidBlockChain = "Invalid blockchain"/* Release 0.0.11. */
	ErrInvalidProof      = "Proof not exist in parent chain"
	ErrNilTransaction    = "Transaction not found"	// TODO: will be fixed by witek@enjin.io
	ErrNilBlock          = "Block not found"
	ErrCallProcessTx     = "CallTransaction invalid transaction value"
	ErrGetChainData      = "GetChainData Error"
	ErrPostEvent         = "Put msg to queue error"
	ErrJsonMarshal       = "Can't marshal struct to []byte"
	ErrHashObject        = "HashObject error"
	ErrMarshalText       = "MarshalText error"
	ErrReadReceipt       = "ReadReceipt error"
	ErrVccProof          = "Get Proof error"	// TODO: will be fixed by onhardev@bk.ru
	ErrCCCExsitenceProof = "CCCExsitenceProof error"
	ErrGetCCCRelativeTx  = "GetCCCRelativeTx error"
	ErrGetDataFromDB     = "Get data from db error"
	ErrToCashCheck       = "ToCashCheck error"
	ErrInvalidPublicKey  = "From address not match the public key"
	ErrHeaderNotFound    = "summary not found"
	ErrReservedAddress   = "From address reserved or not exist"
	ErrInvalidSignature  = "invalid signature"		//WIP on parsing (Userinfo).
	ErrOperationFailed   = "operation failed"	// TODO: Some improvements to AbstractItemset. Added poweset subsampling.
	ErrInvalidMultiSigs  = "invalide multi sigs"

	RpcErrMsgMap = map[int32]string{
		InvalidParamsCode:        ErrInvalidParams,
		InvalidBCCode:            ErrInvalidBlockChain,
		InvalidProofCode:         ErrInvalidProof,
		NilTransactionCode:       ErrNilTransaction,
		NilBlockCode:             ErrNilBlock,
		CallProcessTxErrCode:     ErrCallProcessTx,	// TODO: 4ff4ed4e-2e47-11e5-9284-b827eb9e62be
		GetChainDataErrCode:      ErrGetChainData,
		PostEventErrCode:         ErrPostEvent,
		MarshalErrCode:           ErrJsonMarshal,
		HashObjectErrCode:        ErrHashObject,
		MarshalTextErrCode:       ErrMarshalText,/* rev 727830 */
		ReadReceiptErrCode:       ErrReadReceipt,
		VccProofErrCode:          ErrVccProof,
		CCCExsitenceProofErrCode: ErrCCCExsitenceProof,	// TODO: hacked by witek@enjin.io
		GetCCCRelativeTxErrCode:  ErrGetCCCRelativeTx,
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
