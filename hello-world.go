package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/:name", indexHandler)
	router.Run()
}

func indexHandler(ctx *gin.Context) {
	name := ctx.Params.ByName("name")
	ctx.JSON(200, gin.H{"message": name})
}
