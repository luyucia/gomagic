package bitset

type BitSet struct {
	bitmap []byte
	length uint64
}

// 对外公开的方法.
func New(bitLength uint64) *BitSet {
	return new(BitSet).Init(bitLength)
}

func (this *BitSet) Init(bitLength uint64) *BitSet {
	this.bitmap = make([]byte, bitLength/8)
	this.length = bitLength
	return this
}

func (this *BitSet) Set(bitIndex uint64) {
	//计算偏移 除以8
	offset := bitIndex >> 3
	//________\________\________\________\
	//    0        1        2        3
	//_ _ _ _ \ _ _ _ _
	//7 6 5 4 \ 3 2 1 0
	this.bitmap[offset] |= 0x01 << (bitIndex % 8)

	//fmt.Printf("%b\n",int8(this.bitmap[offset]) )

}
func (this *BitSet) Get(bitIndex uint64) bool {
	offset := bitIndex >> 3
	return (bitIndex < this.length) && ((this.bitmap[offset]>>(bitIndex%8))&0x01 != 0)

}

func (this *BitSet) Clear() {

}

//返回二进制中1的个数
func (this *BitSet) Count() {

}

//返回指定范围中1的个数
func (this *BitSet) CountRange(begin uint64, end uint64) {

}
