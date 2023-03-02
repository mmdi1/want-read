package read

import (
	"log"
	"want-read/core/db"
	"want-read/core/txt"
)

type App struct {
	book *db.Book
	//当前阅读的书
	current [][]rune
	//当前页数
	readPage int
}

func NewApp() *App {
	return &App{}
}

func (sef *App) GetBookshelf() any {
	result := db.QueryList[db.Book](db.T_BOOKS)
	return result.Data
}

func (sef *App) GetBook() db.Book {
	out := db.Book{}
	result := db.QueryList[db.Book](db.T_BOOKS)
	if len(result.Data) > 0 {
		out = result.Data[0]
	}
	return out
}

// 上一页
func (sef *App) PrevPage(page int) string {
	sef.readPage--
	if sef.readPage < 0 {
		sef.readPage = 0
		return "fisrt page"
	}
	return string(sef.current[page])
}

// 下一页
func (sef *App) NextPage(page int) string {
	sef.readPage++
	if sef.readPage > len(sef.current) {
		sef.readPage--
		return "last page"
	}
	return string(sef.current[page])
}

// 重新分页
func (sef *App) ReloadPage(id string) string {
	result := db.QueryList[db.Book](db.T_BOOKS)
	for i := 0; i < len(result.Data); i++ {
		if result.Data[i].IdKey == id {
			sef.book = &result.Data[i]
			break
		}
	}
	size, err := db.GetNum(db.K_ShowSize)
	if err != nil {
		log.Println("reload page get num err:", err)
		return err.Error()
	}
	bt, err := txt.ReadTxt("../../upload/" + sef.book.FileName)
	if err != nil {
		log.Println("reload page read txt err:", err)
		return err.Error()
	}
	sef.current = txt.PageSlicing([]rune(string(bt)), size)
	page := sef.book.ReadSize / size
	if page > len(sef.current) {
		return "last page"
	}
	sef.readPage = page
	return string(sef.current[page])
}
