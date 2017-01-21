package pair

import (
	"testing"
)

const(
	testInt = 1
)

func TestPair(t *testing.T){
	int8Pair := Int8Pair{testInt, testInt}
	if int8Pair.left()!=int8(testInt)||int8Pair.right()!=int8(testInt) {
		t.Errorf("希望得到Int8Pair{%v,%v} 的left=%v,right=%v", testInt, testInt, testInt, testInt)
	}

	int16Pair := Int16Pair{testInt,testInt}
	if int16Pair.left()!=int16(testInt)||int16Pair.right()!=int16(testInt) {
		t.Errorf("希望得到Int8Pair{%v,%v} 的left=%v,right=%v",testInt,testInt,testInt,testInt)
	}


}
