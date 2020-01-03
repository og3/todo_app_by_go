package main

import (
	// ginのインポート
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*.html")

    data := "Hello Go/Gin!!"

    router.GET("/", func(ctx *gin.Context){
        // 変数dataをviewに持っていく
        ctx.HTML(200, "index.html", gin.H{"data": data})
    })

    router.Run()
}
