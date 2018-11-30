package bitset

type BitSet struct {
	bitmap []uint64
	length uint64
	wordInUse int
}

// 对外公开的方法.
func New(bitLength uint64) *BitSet {
	return new(BitSet).Init(bitLength)
}

func (this *BitSet) Init(bitLength uint64) *BitSet {
	this.bitmap = make([]uint64, bitLength/64)
	this.length = bitLength
	return this
}

func (this *BitSet) Set(bitIndex uint64) {
	//计算偏移 除以64
	offset := bitIndex >> 6

	this.bitmap[offset] |= 0x01 << (bitIndex % 64)

	this.wordInUse = int(offset)

}
func (this *BitSet) Get(bitIndex uint64) bool {
	offset := bitIndex >> 6
	return (bitIndex < this.length) && ((this.bitmap[offset]>>(bitIndex%64))&0x01 != 0)

}

func (this *BitSet) Clear() {
	for i:= 0 ;i<=this.wordInUse ;i++  {
		this.bitmap[i] = 0
	}
	this.wordInUse = 0

}

//返回一个数中1出现的次数
func (this* BitSet) BitCount (i uint64) int {
	// 算法一,引用自java源码
	i = i - ((i >> 1) & 0x5555555555555555)
	i = (i & 0x3333333333333333) + ((i >> 2) & 0x3333333333333333)
	i = (i + (i >> 4)) & 0x0f0f0f0f0f0f0f0f
	i = i + (i >> 8)
	i = i + (i >> 16)
	i = i + (i >> 32)
	return int( i & 0x7f)

	// 算法二,引用自https://github.com/willf/bitset/blob/master/popcnt.go
	//i -= (i >> 1) & 0x5555555555555555
	//i = (i>>2)&0x3333333333333333 + i&0x3333333333333333
	//i += i >> 4
	//i &= 0x0f0f0f0f0f0f0f0f
	//i *= 0x0101010101010101
	//return i >> 56


}

func (this* BitSet) BitCountRange (i uint64,begin int,end int) int{
	if end > 63 {
		panic("error end of bit count range")
	}
	if end-begin == 64 {
		return this.BitCount(i)
	} else if end-begin > 0 {

	}

	return 0
}

//返回二进制中1的个数
func (this *BitSet) Count() int{
	sum := 0
	for i:= 0 ;i<=this.wordInUse ;i++  {
		sum+=this.BitCount(this.bitmap[i])
	}
	return sum
}

//返回指定范围中1的个数
func (this *BitSet) CountRange(begin int, end int) int {



	return 0
}
