package aliasmap

// TODO --M1.4

const (
	// creatorLength is the expected length of the hash
	CreatorLength = 850
	// AliasLength is the expected length of the alias
	AliasLength = 5
)

type FixedLenCreatorBytes [CreatorLength]byte
type FixedLenAliasBytes [AliasLength]byte

func ToFixedLenCreatorBytes(creator []byte) FixedLenCreatorBytes {
	var fixedCreator FixedLenCreatorBytes
	fixedCreator.SetBytes(creator)
	return fixedCreator
}

func ToFixedLenAliasBytes(alias []byte) FixedLenAliasBytes {
	var fixedAlias FixedLenAliasBytes
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

func (h *FixedLenAliasBytes) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-AliasLength:]
	}
	// TODO  内存拷贝
	copy(h[AliasLength-len(b):], b)
}

// Bytes gets the byte representation of the underlying hash.
func (h FixedLenCreatorBytes) Bytes() []byte { return h[:] }

func (h FixedLenAliasBytes) Bytes() []byte { return h[:] }

// 在这里建立映射map，并实现处理的相关函数
// 哈希映射为定长的creator字节数组→短字节数组
var AliasForCreator = make(map[FixedLenCreatorBytes][]byte)

var CreaterForAlias = make(map[FixedLenAliasBytes][]byte)

// TODO 使用哈夫曼编码
var CurEncode = 1

var CreatorsChan = make(chan FixedLenCreatorBytes, 500)

func (h FixedLenCreatorBytes) RecoverCreatorBytesLen() []byte {
	// M1.4 去除定长的前导0
	// 计算前导0的数量
	zeroNum := 0
	for i := 0; i < len(h) && h[i] == 0; i++ {
		zeroNum++
	}
	return h[zeroNum:]
}
