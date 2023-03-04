package txt

import (
	"log"
	"testing"
	"time"
	"want-read/core/db"
)

func TestReadTxt(t *testing.T) {
	bt, err := ReadTxt("../../upload/zqqsq.txt")
	db.InitLevelDb("../../data")
	log.Println(err, len(bt))
	arr := PageSlicing([]rune(string(bt)), 64)
	log.Println("arr", len(arr), string(arr[0]), string(arr[1]))
	book := db.Book{
		Name:     "123123",
		FileName: "zqqsq.txt",
		UpdateAt: time.Now(),
	}
	err = db.InsertOne(db.T_BOOKS, book)
	result := db.QueryList[db.Book](db.T_BOOKS)
	log.Println("===========>", err, result.Err, result.Data[0])
}
