package api

import (
	"GIN/ReadFiles"
	"GIN/pkg"
	"GIN/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleFileUpload(c *gin.Context) {
	starttime := time.Now()
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not provided"})
		return
	}

	routines, err := strconv.Atoi(c.PostForm("value"))
	if err != nil || routines <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for 'value'"})
		return
	}

	fileData := ReadFiles.ReadFile(file, c)

	operation := utils.FileOperation{}
	createChannel := make(chan utils.FileOperation)

	for i := 0; i < routines; i++ {
		start := i * len(fileData) / routines
		end := (i + 1) * len(fileData) / routines
		go pkg.Files(createChannel, string(fileData[start:end]))
	}

	for i := 0; i < routines; i++ {
		result := <-createChannel
		operation.Vowel += result.Vowel
		operation.Punctuation += result.Punctuation
		operation.Nextline += result.Nextline
		operation.Chars += result.Chars
		operation.Spaces += result.Spaces
	}

	Time := time.Since(starttime).String()

	c.JSON(http.StatusOK, gin.H{
		"Vowel":              operation.Vowel,
		"Punctuation":        operation.Punctuation,
		"NextLine":           operation.Nextline,
		"TotalChars":         operation.Chars,
		"Spaces":             operation.Spaces,
		"ExecutionTime":      Time,
		"Number of routines": routines,
	})
}

// func ReadFile(file *multipart.FileHeader, c *gin.Context) []byte {
// 	src, err := file.Open()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file"})
// 		return nil
// 	}
// 	defer src.Close()

// 	fileData, err := ioutil.ReadAll(src)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading file"})
// 		return nil
// 	}

// 	return fileData
// }
