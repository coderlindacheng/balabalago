package pair

type Pair interface {
	Left() *interface{}
	Right() *interface{}
}

type ObjPair [2]interface{}

func (pair *ObjPair)left() interface{} {
	return pair[0]
}

func (pair *ObjPair)right() interface{} {
	return pair[1]
}

type BytePair [2]byte

func (pair *BytePair)left() byte {
	return pair[0]
}

func (pair *BytePair)right() byte {
	return pair[1]
}

type Uint8Pair [2]uint8

func (pair *Uint8Pair)left() uint8 {
	return pair[0]
}

func (pair *Uint8Pair)right() uint8 {
	return pair[1]
}

type Uint16Pair [2]uint16

func (pair *Uint16Pair)left() uint16 {
	return pair[0]
}

func (pair *Uint16Pair)right() uint16 {
	return pair[1]
}

type Uint32Pair [2]uint32

func (pair *Uint32Pair)left() uint32 {
	return pair[0]
}

func (pair *Uint32Pair)right() uint32 {
	return pair[1]
}

type Uint64Pair [2]uint64

func (pair *Uint64Pair)left() uint64 {
	return pair[0]
}

func (pair *Uint64Pair)right() uint64 {
	return pair[1]
}

type UintptrPair [2]uintptr

func (pair *UintptrPair)left() uintptr {
	return pair[0]
}

func (pair *UintptrPair)right() uintptr {
	return pair[1]
}

type UintPair [2]uint

func (pair *UintPair)left() uint {
	return pair[0]
}

func (pair *UintPair)right() uint {
	return pair[1]
}

type Int8Pair [2]int8

func (pair *Int8Pair)left() int8 {
	return pair[0]
}

func (pair *Int8Pair)right() int8 {
	return pair[1]
}

type Int16Pair [2]int16

func (pair *Int16Pair)left() int16 {
	return pair[0]
}

func (pair *Int16Pair)right() int16 {
	return pair[1]
}

type Int32Pair [2]int32

func (pair *Int32Pair)left() int32 {
	return pair[0]
}

func (pair *Int32Pair)right() int32 {
	return pair[1]
}

type Int64Pair [2]int64

func (pair *Int64Pair)left() int64 {
	return pair[0]
}

func (pair *Int64Pair)right() int64 {
	return pair[1]
}

type IntPair [2]int

func (pair *IntPair)left() int {
	return pair[0]
}

func (pair *IntPair)right() int {
	return pair[1]
}

type Float32Pair [2]float32

func (pair *Float32Pair)left() float32 {
	return pair[0]
}

func (pair *Float32Pair)right() float32 {
	return pair[1]
}

type Float64Pair [2]float64

func (pair *Float64Pair)left() float64 {
	return pair[0]
}

func (pair *Float64Pair)right() float64 {
	return pair[1]
}
