package test

import (
	"bufio"
	"io"
	"net/http"
	"os"
)

func DownFile(path string) (string, error) {
	res, httpErr := http.Get(path)
	if httpErr != nil {
		return "Http Error", httpErr
	}

	defer res.Body.Close()
	reader := bufio.NewReaderSize(res.Body, 64*1024) // 获得reader对象
	file, createErr := os.Create("./ttttt")
	if createErr != nil {
		return "Create Error", createErr
	}

	writer := bufio.NewWriter(file) // 获得writer对象
	_, copyErr := io.Copy(writer, reader)
	if copyErr != nil {
		return "Copy Error", copyErr
	}

	return "111", nil
}

func main() {

}
