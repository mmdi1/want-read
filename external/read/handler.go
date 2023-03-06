package read

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
	"want-read/configs"
	"want-read/core/db"
	"want-read/core/txt"
	"want-read/external/setting"
	"want-read/server/api/ws"

	"golang.org/x/text/transform"

	"golang.org/x/text/encoding/simplifiedchinese"

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
	for i := 0; i < len(result.Data); i++ {
		bt, err := db.Get(db.K_Read + result.Data[i].IdKey)
		if err != nil {
			log.Println("err:", err)
			continue
		}
		book := db.Book{}
		json.Unmarshal(bt, &book)
		result.Data[i] = book
	}
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
func (sef *App) ReloadPage(id string, readSize int) string {
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
	fmt.Println("=================>", configs.ReadBook.ReadSize, readSize)
	if readSize > 0 {
		configs.ReadBook.ReadSize = readSize
	}
	configs.SettingModel, err = setting.GetSetting()
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
	configs.ReadBook.TotalSize = len(configs.ReadBook.Conetent) * configs.SettingModel.ShowSize
	configs.CurrentPage = page
	ws.Conn.SendMsg(db.WsMsgModel{Id: db.WsID_Update, Data: configs.SettingModel})
	return string(configs.ReadBook.Conetent[page])
}

func (sef *App) SelectFile() string {
	lock := sef.TryLock()
	if !lock {
		return "正在上传中!"
	}
	defer sef.Unlock()
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
	dirPath := "./upload/"
	// 判断目录是否存在，不存在则创建
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			panic(err)
		}
	}
	if (len(data) >= 3 && data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF) || utf8.Valid(data) {
		log.Printf("%s is UTF-8 encoded with BOM\n", fileName)
		err := os.WriteFile(dirPath+fileName, data, 0644)
		if err != nil {
			log.Println("write error: ", err)
			return err.Error()
		}
	} else {
		log.Printf("%s is not UTF-8 encoded\n", fileName)
		// 转码成utf-8格式
		utf8Data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(data),
			simplifiedchinese.GBK.NewDecoder()))
		if err != nil {
			log.Println("ioutil readall file err:", err)
			return err.Error()
		}
		// 重写写入成utf-8格式文本
		file, err := os.Create(dirPath + fileName)
		if err != nil {
			log.Println("create path file err:", err)
			return err.Error()
		}
		defer file.Close()
		bufferOut := bufio.NewWriter(file)
		_, err = bufferOut.Write(utf8Data)
		if err != nil {
			log.Println("buffer write err:", err)
			return err.Error()
		}
		bufferOut.Flush()
	}
	bt, err := txt.ReadTxt(dirPath + fileName)
	if err != nil {
		log.Println("read txt err:", err)
		return err.Error()
	}
	configs.SettingModel, err = setting.GetSetting()
	if err != nil {
		log.Println("get setting err:", err)
	}
	id := db.NextId()
	book := db.Book{
		IdKey:     id,
		Name:      fileName,
		FileName:  fileName,
		UpdateAt:  time.Now(),
		TotalSize: len(txt.PageSlicing([]rune(string(bt)), configs.SettingModel.ShowSize)) * configs.SettingModel.ShowSize,
		ReadSize:  0,
	}
	db.InsertOne(db.T_BOOKS, book)
	return path
}
