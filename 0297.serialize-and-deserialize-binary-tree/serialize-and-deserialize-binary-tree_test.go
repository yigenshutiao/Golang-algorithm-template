package _297_serialize_and_deserialize_binary_tree

import (
	"fmt"
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"testing"
)

func TestCodec_deserialize(t *testing.T) {

	tcs := struct {
		pre, in []int
	}{
		[]int{1, 2, 3, 4, 5},
		[]int{2, 1, 4, 3, 5},
	}

	ser := Constructor()

	deser := Constructor()

	data := ser.serialize(util.PreIn2Tree(tcs.pre, tcs.in))

	ans := deser.deserialize(data)

	fmt.Println(ans)
}
