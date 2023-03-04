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
	IdKey    string    `json:"id"`
	Name     string    `json:"name"`
	FileName string    `json:"file_name"`
	UpdateAt time.Time `json:"update_at"`
	//总字数
	TotalSize int `json:"total_size"`
	//阅读到的字数
	ReadSize int `json:"read_size"`
	//分好页的数组
	Conetent [][]rune `json:"-"`
}

const (
	WsID_PrevPage = iota
	WsID_NextPage
	WsID_HidePanel
	//配置界面
	WsID_Setting
	//退出
	WsID_Exit
	//更新
	WsID_Update
)

type WsMsgModel struct {
	Id   int    `json:"id"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}
type Setting struct {
	PrevGroup  []uint32 `json:"prev_group"`
	NextGroup  []uint32 `json:"next_group"`
	HideGroup  []uint32 `json:"hide_group"`
	FontSize   int      `json:"font_size"`
	FontColor  string   `json:"font_color"`
	ShowSize   int      `json:"show_size"`
	ReadWidth  int      `json:"read_width"`
	ReadHeight int      `json:"read_height"`
}
type Operationkey struct {
	Name    string
	Handler int
	Group   []uint32
}
