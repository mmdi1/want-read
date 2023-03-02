package db

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	_db     *leveldb.DB
	once    sync.Once
	T_BOOKS = "t_books"
	//每页显示字数
	K_ShowSize = "k_size"
	//自增id
	k_id = []byte("k_id")
)

func InitLevelDb(path string) {
	once.Do(func() {
		var err error
		_db, err = leveldb.OpenFile(path, nil)
		if err != nil {
			log.Fatalln("open level db err:", err)
		}
	})
}
func QueryList[T IModel](key string) Result[T] {
	out := Result[T]{}
	bt, err := _db.Get([]byte(key), nil)
	if err != nil {
		log.Println("query db err:", err)
		out.Err = err
	}
	if len(bt) > 0 {
		err := json.Unmarshal(bt, &out.Data)
		if err != nil {
			out.Err = err
			log.Println("db unmarshal json err:", err)
		}
	}
	return out
}

func InsertOne[T IModel](key string, t T) error {
	result := QueryList[T](key)
	list := []T{}
	if (len(result.Data)) > 0 {
		list = append(list, result.Data...)
	}
	list = append(list, t)
	bt, err := json.Marshal(list)
	if err != nil {
		log.Println("db set marshal json err:", err)
		return err
	}
	return _db.Put([]byte(key), bt, nil)
}

// 根据key获取val
func Get(key string) ([]byte, error) {
	return _db.Get([]byte(key), nil)
}

// 根据key获取num
func GetNum(key string) (int, error) {
	bt, _ := Get(key)
	return strconv.Atoi(string(bt))
}

// 根据key插入val
func Set(key, val string) error {
	return _db.Put([]byte(key), []byte(val), nil)
}

// 获取自增下一个id
func NextId() string {
	bt, err := _db.Get(k_id, nil)
	if err != nil {
		_db.Put(k_id, []byte("1"), nil)
		return "1"
	}
	id_str := string(bt)
	id, _ := strconv.Atoi(id_str)
	id++
	id_str = strconv.Itoa(id)
	_db.Put(k_id, []byte(id_str), nil)
	return id_str
}
