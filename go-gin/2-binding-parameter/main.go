package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

/*
本小节主要展示参数绑定（通俗点讲，就是把前端的参数自动填充到我们自己定义的结构体中）
1. 绑定query string和 post data
2. 绑定html checkboxes
3. 绑定url参数

*/

func main() {
	router := setupRouter() // 启动路由
	router.Run(":8080")     // 监听端口8080
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	//localhost:8080/user/search/admin/123
	router.GET("/user/search/:name/:id", userSearchHandler)

	router.GET("/testing", testingHandler)
	return router
}

type PersonInfo struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func testingHandler(c *gin.Context) {
	var personInfo PersonInfo
	if c.ShouldBind(&personInfo) == nil {
		log.Println(personInfo.Name)
		log.Println(personInfo.Address)
	}
}

// 1. 绑定URL参数

type Person struct {
	ID   string `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func userSearchHandler(c *gin.Context) {
	var person Person

	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, gin.H{"msg": err})
	}
	c.JSON(200, gin.H{
		"name": person.Name, "id": person.ID})
}

// 2. 绑定数组（form中的多选）

type myForm struct {
	Colors []string `form:"colors[]"`
}
