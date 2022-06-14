package aliasmap

import (
	"bytes"
	"encoding/binary"
)

// TODO --M1.4

const (
	// creatorLength is the expected length of the hash
	CreatorLength = 850
	// AliasLength is the expected length of the alias
	// AliasLength = 5
	// 3个长度，指示映射的值中对应的长度
	BlockNoLen    = 8
	TxNoLen       = 2
	EndorserNoLen = 2
)

type FixedLenCreatorBytes [CreatorLength]byte

// type FixedLenAliasBytes [AliasLength]byte

type CertInfo struct {
	Cert       FixedLenCreatorBytes
	BlockNo    uint64
	TxNo       int
	EndorserNo int
}

// type AliasValue struct {
// 	BlockNo    uint64
// 	TxNo       int
// 	EndorserNo int
// }

type AliasValue [BlockNoLen + TxNoLen + EndorserNoLen]byte

// TODO: 目前把长度写死，后续应当设置为可变
func SetAliasValue(BlockNO uint64, TxNO, EndorserNO uint16) AliasValue {
	bytesBuffer := bytes.NewBuffer([]byte{})
	// TODO: M1.4 debug
	binary.Write(bytesBuffer, binary.BigEndian, BlockNO)
	binary.Write(bytesBuffer, binary.BigEndian, TxNO)
	binary.Write(bytesBuffer, binary.BigEndian, EndorserNO)

	var alias AliasValue

	binary.Read(bytesBuffer, binary.BigEndian, &alias)

	return alias
}

func ToFixedLenCreatorBytes(creator []byte) FixedLenCreatorBytes {
	var fixedCreator FixedLenCreatorBytes
	fixedCreator.SetBytes(creator)
	return fixedCreator
}

// func ToFixedLenAliasBytes(alias []byte) FixedLenAliasBytes {
// 	var fixedAlias FixedLenAliasBytes
// 	fixedAlias.SetBytes(alias)
// 	return fixedAlias
// }

func ToFixedLenAliasValue(alias []byte) AliasValue {
	var fixedAlias AliasValue
	fixedAlias.SetBytes(alias)
	return fixedAlias
}

// SetBytes sets the FixedLenCreatorBytes to the value of b.
// If b is larger than len(h), b will be cropped from the left.
func (h *FixedLenCreatorBytes) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-CreatorLength:]
	}
	// TODO  内存拷贝
	copy(h[CreatorLength-len(b):], b)
}

// func (h *FixedLenAliasBytes) SetBytes(b []byte) {
// 	if len(b) > len(h) {
// 		b = b[len(b)-AliasLength:]
// 	}
// 	// TODO  内存拷贝
// 	copy(h[AliasLength-len(b):], b)
// }

func (h *AliasValue) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-(BlockNoLen+TxNoLen+EndorserNoLen):]
	}
	// TODO  内存拷贝
	copy(h[(BlockNoLen+TxNoLen+EndorserNoLen)-len(b):], b)
}

// Bytes gets the byte representation of the underlying hash.
func (h FixedLenCreatorBytes) Bytes() []byte { return h[:] }

// func (h FixedLenAliasBytes) Bytes() []byte { return h[:] }

func (h AliasValue) Bytes() []byte { return h[:] }

// 在这里建立映射map，并实现处理的相关函数
// 哈希映射为定长的creator字节数组→短字节数组
// var AliasForCreator = make(map[FixedLenCreatorBytes][]byte)
// var CreatorForAlias = make(map[FixedLenAliasBytes][]byte)
var CreatorForAlias = make(map[AliasValue][]byte)

var AliasForCreator = make(map[FixedLenCreatorBytes]AliasValue)

// TODO 使用哈夫曼编码
var CurEncode = 1

// var CreatorsChan = make(chan FixedLenCreatorBytes, 500)
// 修改chan的类型为结构体，在传输证书内容的同时传输位置信息
var CreatorsChan = make(chan *CertInfo, 500)

func (h FixedLenCreatorBytes) RecoverCreatorBytesLen() []byte {
	// M1.4 去除定长的前导0
	// 计算前导0的数量
	zeroNum := 0
	for i := 0; i < len(h) && h[i] == 0; i++ {
		zeroNum++
	}
	return h[zeroNum:]
}
