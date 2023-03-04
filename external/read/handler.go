package read

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"
	"want-read/configs"
	"want-read/core/db"
	"want-read/core/txt"
	"want-read/external/setting"
	"want-read/server/api/ws"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	sync.Mutex
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
			configs.ReadBook = &result.Data[i]
			break
		}
	}
	if configs.ReadBook == nil {
		return "没有书籍！"
	}
	bts, err := db.Get(db.K_Read + configs.ReadBook.IdKey)
	if err == nil {
		out := db.Book{}
		err := json.Unmarshal(bts, &out)
		if err == nil {
			configs.ReadBook.ReadSize = out.ReadSize
		}
	}
	configs.SettingModel, err = setting.GetSetting()
	if configs.SettingModel.ShowSize == 0 {
		configs.SettingModel.ShowSize = 50
	}
	if err != nil {
		log.Println("get setting err:", err)
	}
	bt, err := txt.ReadTxt("./upload/" + configs.ReadBook.FileName)
	if err != nil {
		log.Println("reload page read txt err:", err, configs.ReadBook.FileName)
		return err.Error()
	}
	configs.ReadBook.Conetent = txt.PageSlicing([]rune(string(bt)), configs.SettingModel.ShowSize)
	page := configs.ReadBook.ReadSize / configs.SettingModel.ShowSize
	if page > len(configs.ReadBook.Conetent) {
		return "last page"
	}
	configs.CurrentPage = page
	ws.Conn.SendMsg(db.WsMsgModel{Id: db.WsID_Update, Data: configs.SettingModel})
	return string(configs.ReadBook.Conetent[page])
}

func (sef *App) SelectFile() string {
	path, err := runtime.OpenFileDialog(configs.APP_CTX, runtime.OpenDialogOptions{})
	if err != nil {
		log.Println("select file err:", err)
		return err.Error()
	}
	fileName := path[strings.LastIndex(path, "\\")+1:]
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	// 转码成utf-8格式
	utf8Data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(data),
		simplifiedchinese.GBK.NewDecoder()))
	if err != nil {
		log.Fatal(err)
	}
	dirPath := "./upload/"
	// 判断目录是否存在，不存在则创建
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			panic(err)
		}
	}
	// 重写写入成utf-8格式文本
	file, err := os.Create(dirPath + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bufferOut := bufio.NewWriter(file)
	_, err = bufferOut.Write(utf8Data)
	if err != nil {
		log.Fatal(err)
	}
	bufferOut.Flush()
	id := db.NextId()
	book := db.Book{
		IdKey:    id,
		Name:     fileName,
		FileName: fileName,
		UpdateAt: time.Now(),
	}
	db.InsertOne(db.T_BOOKS, book)
	return path
}
