package file

import (
	"want-read/server/check"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	c.SaveUploadedFile(file, "./upload/"+file.Filename)
	c.JSON(200, check.ErrSuccessNull())
}
