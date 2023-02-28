package txt

import (
	"fmt"
	"testing"
)

func TestReadTxt(t *testing.T) {
	bt, err := ReadTxt("../../upload/最强弃少.txt")
	fmt.Println(err, len(bt))
	arr := PageSlicing(string(bt), 64)
	fmt.Println("arr", len(arr), arr[0], "11", arr[1], "22", arr[2], "33", arr[3])
}
