package file

import (
	"time"
	"want-read/core/db"
	"want-read/server/check"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	c.SaveUploadedFile(file, "./upload/"+file.Filename)
	id := db.NextId()
	book := db.Book{
		IdKey:     id,
		Name:      file.Filename,
		FileName:  file.Filename,
		CreatedAt: time.Now(),
	}
	db.InsertOne(db.T_BOOKS, book)
	c.JSON(200, check.ErrSuccessNull())
}
