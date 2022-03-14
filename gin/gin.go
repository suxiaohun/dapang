package gin

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"xiaoming": "123",
			"xiaomu":   "456",
		})
	})

	r.GET("/user/:name/:action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")

		result := name + " is " + action

		context.JSON(200, gin.H{
			"result": result,
		})
	})
	r.GET("/welcome", func(context *gin.Context) {
		name := context.Query("name")
		age := context.Query("age")

		context.String(200, "hello %s this year is %s", name, age)
	})

	r.POST("/form", func(context *gin.Context) {
		name := context.PostForm("name")
		age := context.DefaultPostForm("age", "11")

		context.String(200, "hello %s age is %s", name, age)
	})

	r.MaxMultipartMemory = 8 << 20

	r.POST("/upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		log.Println(file.Filename)

		context.SaveUploadedFile(file, "upload")
		context.String(200, "get the upload file:  %s", file.Filename)
	})

	r.GET("/su", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"func2": "222",
		})
	})

	r.Run(":8080")
}

