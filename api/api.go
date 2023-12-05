package api

import (
	"convert/convert"
	"fmt"

	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/download/:filename", downloadImage)

	api := r.Group("/api")
	api.POST("/convert", convertImage)

	return r
}

func Run() {
	r := setupRouter()
	r.Run(":8080")
}

func convertImage(c *gin.Context) {
	image, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"result": nil,
				"ok":     false,
				"error":  "Failed to get image file from the request.",
			})
		return
	}

	format := c.PostForm("format")
	if format == "" {
		format = "jpg"
	}

	// Читаем содержимое файла в байты
	fileBytes, err := convert.ReadFile(image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"ok":     false,
			"error":  fmt.Sprintf("Failed to read image file. Error: %v", err),
		})
		return
	}

	convertedImage, err := convert.ConvertImage(fileBytes, format)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"result": nil,
				"ok":     false,
				"error":  fmt.Sprintf("Failed to convert image to %s. Error: %v", format, err),
			})
		return
	}

	imagename :=  uuid.NewString() + "." + format
	filename := "static\\uploads\\" + imagename

	err = convertedImage.WriteImageBytes(filename)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"result": nil,
				"ok":     false,
				"error":  fmt.Sprintf("Failed to save image to %s. Error: %v", format, err),
			})
		return
	}

	fullPath := "/download/" + imagename

	c.JSON(http.StatusOK,
		gin.H{
			"result": fullPath,
			"ok":     true,
			"error":  nil,
		})

}

func downloadImage(c *gin.Context) {
	filename := c.Param("filename")

	// Замените на фактический путь к вашему каталогу с изображениями
	imagePath := filepath.Join("static/uploads", filename)

	// Передача файла для скачивания
	c.FileAttachment(imagePath, filename)
}
