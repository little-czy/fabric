/*
	新建blockCache包用于缓存unmarshal出的block的内容，减少性能开销
*/
package blockCache

import (
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/rwsetutil"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/peer"
)

type BlockCache struct {
	Num      uint64
	TxsCache []*TransactionCache
}

type TransactionCache struct {
	IndexInBlock            int
	ID                      string
	Env                     *common.Envelope
	Chdr                    *common.ChannelHeader
	Payl                    *common.Payload
	Tx                      *peer.Transaction
	CcPayload               *peer.ChaincodeActionPayload
	PRespPayload            *peer.ProposalResponsePayload
	RespPayload             *peer.ChaincodeAction
	Rwset                   *rwsetutil.TxRwSet
	ValidationCode          peer.TxValidationCode
	ContainsPostOrderWrites bool
}

// ------mytest 声明全局变量，对chdr和tx信息进行缓存
var BCache *BlockCache
