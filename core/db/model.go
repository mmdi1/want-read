package db

import "time"

type IModel interface {
	Book |
		string
}

type Result[T IModel] struct {
	Data []T
	Err  error
}
type Book struct {
	//自动生成的id
	IdKey     string    `json:"id"`
	Name      string    `json:"name"`
	FileName  string    `json:"file_name"`
	CreatedAt time.Time `json:"created_at"`
	//总字数
	TotalSize int `json:"total_size"`
	//阅读到的字数
	ReadSize int `json:"read_size"`
}
