package txt

import (
	"fmt"
	"testing"
	"time"
	"want-read/core/db"
)

func TestReadTxt(t *testing.T) {
	bt, err := ReadTxt("../../upload/zqqsq.txt")
	db.InitLevelDb("../../data")
	fmt.Println(err, len(bt))
	arr := PageSlicing([]rune(string(bt)), 64)
	fmt.Println("arr", len(arr), string(arr[0]), string(arr[1]))
	book := db.Book{
		Name:      "123123",
		FileName:  "zqqsq.txt",
		CreatedAt: time.Now(),
	}
	err = db.InsertOne(db.T_BOOKS, book)
	result := db.QueryList[db.Book](db.T_BOOKS)
	fmt.Println("===========>", err, result.Err, result.Data[0])
}
