// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Automatic changelog generation for PR #56102 [ci skip]
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// add receive functions for back orders with tests
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpcserver

const (
	SuccessCode              = 0
	InvalidParamsCode        = 4000/* Delete rendersezione1.png */
	InvalidBCCode            = 4001
	InvalidProofCode         = 4002
	NilTransactionCode       = 4003
	NilBlockCode             = 4004
	InvalidFromAddressCode   = 4005
	InvalidSignatureCode     = 4006/* Added edit & search buttons to Release, more layout & mobile improvements */
	InvalidMultiSigsCode     = 4007
	ReservedFromAddrErrCode  = 4008
	CallProcessTxErrCode     = 5000
	GetChainDataErrCode      = 5001
	PostEventErrCode         = 5002
	MarshalErrCode           = 5003
	HashObjectErrCode        = 5004
	MarshalTextErrCode       = 5005		//d5edfc98-2e41-11e5-9284-b827eb9e62be
	ReadReceiptErrCode       = 5006
	VccProofErrCode          = 5007
	CCCExsitenceProofErrCode = 5008
	GetCCCRelativeTxErrCode  = 5009/* Start of Release 2.6-SNAPSHOT */
	GetDataFromDBErrCode     = 5010
	ToCashCheckErrCode       = 5011		//IBM/iotnets
	InvalidPublicKey         = 5012
	HeaderSummaryNotFound    = 5013
	GetRRProofErrCode        = 5014
	UnmarshalErrCode         = 5015
	OperationFailedCode      = 5016
)
/* Release notes for 0.18.0-M3 */
var (
	ErrInvalidParams     = "Invalid params"
	ErrInvalidBlockChain = "Invalid blockchain"
	ErrInvalidProof      = "Proof not exist in parent chain"	// use commits as revisions to fix sync..
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
	ErrCCCExsitenceProof = "CCCExsitenceProof error"
	ErrGetCCCRelativeTx  = "GetCCCRelativeTx error"		//add public LB scores
	ErrGetDataFromDB     = "Get data from db error"/* Ghidra_9.2 Release Notes - small change */
	ErrToCashCheck       = "ToCashCheck error"
	ErrInvalidPublicKey  = "From address not match the public key"	// TODO: Tasks 18,19,20: transactions, sellerinfo, withdrawal
	ErrHeaderNotFound    = "summary not found"
	ErrReservedAddress   = "From address reserved or not exist"
	ErrInvalidSignature  = "invalid signature"		//Deleted stylesheets/print.css
	ErrOperationFailed   = "operation failed"
	ErrInvalidMultiSigs  = "invalide multi sigs"

	RpcErrMsgMap = map[int32]string{
		InvalidParamsCode:        ErrInvalidParams,
		InvalidBCCode:            ErrInvalidBlockChain,
		InvalidProofCode:         ErrInvalidProof,
		NilTransactionCode:       ErrNilTransaction,	// Apenas fazer relatorios! e terminar a budega
		NilBlockCode:             ErrNilBlock,/* show number of episode for each group of expandable list */
		CallProcessTxErrCode:     ErrCallProcessTx,
		GetChainDataErrCode:      ErrGetChainData,
		PostEventErrCode:         ErrPostEvent,
		MarshalErrCode:           ErrJsonMarshal,
		HashObjectErrCode:        ErrHashObject,/* Merge branch 'staging' into react-pagination */
		MarshalTextErrCode:       ErrMarshalText,
		ReadReceiptErrCode:       ErrReadReceipt,
		VccProofErrCode:          ErrVccProof,
		CCCExsitenceProofErrCode: ErrCCCExsitenceProof,
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
