package txt

import (
	"os"
)

func ReadTxt(path string) ([]byte, error) {
	return os.ReadFile(path)

}
func PageSlicing(bt string, size int) []string {
	arr := []string{}
	i := 0
	bt_len := len(bt)
	for {
		start := i * size
		end := start + size
		if end < bt_len {
			if start != 0 {
				arr = append(arr, bt[start-1:start+size-1])
			} else {
				arr = append(arr, bt[start:start+size-1])
			}
			i++
			continue
		}
		arr = append(arr, bt[start:])
		break
	}
	return arr
}
