package ReadFiles

import (
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadFile(file *multipart.FileHeader, c *gin.Context) []byte {
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file"})
		return nil
	}
	defer src.Close()

	fileData, err := ioutil.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading file"})
		return nil
	}

	return fileData
}
