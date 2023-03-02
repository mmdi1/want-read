package read

import (
	"log"
	"sync"
	"want-read/configs"
	"want-read/core/db"
	"want-read/core/txt"
)

type App struct {
	sync.Mutex
	book *db.Book
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

func (sef *App) RemoveBook(id string) bool {
	result := db.QueryList[db.Book](db.T_BOOKS)
	new_data := []db.Book{}
	for i := 0; i < len(result.Data); i++ {
		if result.Data[i].IdKey != id {
			new_data = append(new_data, result.Data[i])
		}
	}
	db.CoverList(db.T_BOOKS, new_data)
	return true
}

// 重新分页
func (sef *App) ReloadPage(id string) string {
	lock := sef.TryLock()
	if !lock {
		return "正在读取中!"
	}
	defer sef.Unlock()
	result := db.QueryList[db.Book](db.T_BOOKS)
	for i := 0; i < len(result.Data); i++ {
		if result.Data[i].IdKey == id {
			sef.book = &result.Data[i]
			break
		}
	}
	if sef.book == nil {
		return "没有书籍！"
	}
	size, err := db.GetNum(db.K_ShowSize)
	if err != nil {
		log.Println("reload page get num err:", err, sef.book.FileName)
		return err.Error()
	}
	bt, err := txt.ReadTxt("./upload/" + sef.book.FileName)
	if err != nil {
		log.Println("reload page read txt err:", err, sef.book.FileName)
		return err.Error()
	}
	configs.CurrentReadBook = txt.PageSlicing([]rune(string(bt)), size)
	page := sef.book.ReadSize / size
	if page > len(configs.CurrentReadBook) {
		return "last page"
	}
	configs.CurrentPage = page
	return string(configs.CurrentReadBook[page])
}
