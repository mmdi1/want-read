package txt

import (
	"os"
)

func ReadTxt(path string) ([]byte, error) {
	return os.ReadFile(path)

}
func PageSlicing(bt []rune, size int) [][]rune {
	arr := [][]rune{}
	i := 0
	bt_len := len(bt)
	for {
		start := i * size
		end := start + size
		if end < bt_len {
			arr = append(arr, bt[start:end])
			i++
			continue
		}
		arr = append(arr, bt[start:])
		break
	}
	return arr
}
