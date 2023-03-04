package file

import (
	"log"
	"time"
	"want-read/core/db"
	"want-read/server/check"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("from file err:", err)
	}
	log.Println("upload file:", file.Filename)
	c.SaveUploadedFile(file, "./upload/"+file.Filename)
	id := db.NextId()
	book := db.Book{
		IdKey:    id,
		Name:     file.Filename,
		FileName: file.Filename,
		UpdateAt: time.Now(),
	}
	db.InsertOne(db.T_BOOKS, book)
	c.JSON(200, check.ErrSuccessNull())
}
