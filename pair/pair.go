package pair

type Pair [2]interface{}

func (pair *Pair) Left() interface{} {
	return pair[0]
}

func (pair *Pair) Right() interface{} {
	return pair[1]
}

type BytePair [2]byte

func (pair *BytePair) Left() byte {
	return pair[0]
}

func (pair *BytePair) Right() byte {
	return pair[1]
}

type Uint8Pair [2]uint8

func (pair *Uint8Pair) Left() uint8 {
	return pair[0]
}

func (pair *Uint8Pair) Right() uint8 {
	return pair[1]
}

type Uint16Pair [2]uint16

func (pair *Uint16Pair) Left() uint16 {
	return pair[0]
}

func (pair *Uint16Pair) Right() uint16 {
	return pair[1]
}

type Uint32Pair [2]uint32

func (pair *Uint32Pair) Left() uint32 {
	return pair[0]
}

func (pair *Uint32Pair) Right() uint32 {
	return pair[1]
}

type Uint64Pair [2]uint64

func (pair *Uint64Pair) Left() uint64 {
	return pair[0]
}

func (pair *Uint64Pair) Right() uint64 {
	return pair[1]
}

type UintptrPair [2]uintptr

func (pair *UintptrPair) Left() uintptr {
	return pair[0]
}

func (pair *UintptrPair) Right() uintptr {
	return pair[1]
}

type UintPair [2]uint

func (pair *UintPair) Left() uint {
	return pair[0]
}

func (pair *UintPair) Right() uint {
	return pair[1]
}

type Int8Pair [2]int8

func (pair *Int8Pair) Left() int8 {
	return pair[0]
}

func (pair *Int8Pair) Right() int8 {
	return pair[1]
}

type Int16Pair [2]int16

func (pair *Int16Pair) Left() int16 {
	return pair[0]
}

func (pair *Int16Pair) Right() int16 {
	return pair[1]
}

type Int32Pair [2]int32

func (pair *Int32Pair) Left() int32 {
	return pair[0]
}

func (pair *Int32Pair) Right() int32 {
	return pair[1]
}

type Int64Pair [2]int64

func (pair *Int64Pair) Left() int64 {
	return pair[0]
}

func (pair *Int64Pair) Right() int64 {
	return pair[1]
}

type IntPair [2]int

func (pair *IntPair) Left() int {
	return pair[0]
}

func (pair *IntPair) Right() int {
	return pair[1]
}

type Float32Pair [2]float32

func (pair *Float32Pair) Left() float32 {
	return pair[0]
}

func (pair *Float32Pair) Right() float32 {
	return pair[1]
}

type Float64Pair [2]float64

func (pair *Float64Pair) Left() float64 {
	return pair[0]
}

func (pair *Float64Pair) Right() float64 {
	return pair[1]
}
