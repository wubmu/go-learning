package main

/*
主要表明了获取参数的三种方法：
1. 获取querystring的参数，例localhost:8080/user/search?name=admin&address=123456
2. 获取form表单参数
3. 获取path中的参数，例如localhost:8080/user/search/admin
*/
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type User struct {
	ID string
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/login", loginHandler)
	router.POST("/login1", loginHandler1)

	// localhost:8080/user/search?name=admin&address=123456
	router.GET("/user/search", userSearchHandler)

	//localhost:8080/user/search/admin
	router.GET("/user/search2/:username/:address", userSearchHandler1)
	return router
}

// 1. 获取querystring的参数
func userSearchHandler(c *gin.Context) {
	// 获取url携带的query String 参数的3种方式：
	//name := c.Query("name")
	//name :=c.DefaultQuery("name", "默认name")  // 取不到使用默认值

	//name, ok := c.GetQuery("name")	// 取不到参数，第二个参数返回false
	//if !ok {
	//	// 取不到值
	//	name = "默认name"
	//}
	name := c.Query("name")
	address := c.Query("address")
	c.JSON(http.StatusOK, gin.H{
		"name":    name,
		"address": address,
	})
}

func main() {
	router := setupRouter() // 启动路由
	router.Run(":8080")     // 监听端口8080
}

// 3. 获取path中的参数，例如localhost:8080/user/search2/admin
func userSearchHandler1(context *gin.Context) {
	username := context.Param("username")
	address := context.Param("address")
	context.JSON(http.StatusOK, gin.H{
		"username": username,
		"address":  address,
	})
}

// 2. 获取form表单参数
func loginHandler1(context *gin.Context) {
	var form LoginForm
	// 获取form表单的数据
	form.Username = context.PostForm("username")
	form.Password = context.PostForm("password")
	//form.Username = context.DefaultPostForm("username","默认")
	//form.Username, ok = context.GetPostForm("username")

	if form.Username == "admin" && form.Password == "123456" {
		// 这里gin.H 底层就是一个简单的map
		context.JSON(http.StatusOK, gin.H{"status": "登录成功"}) // StatusOk状态码 200
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"status": "账号或密码错误"}) // StatusUnauthorized状态码 401
	}

}

// 4. post请求，form表单绑定,参数绑定
func loginHandler(context *gin.Context) {
	var form LoginForm
	// 自动绑定，可以显示绑定声明绑定
	//context.ShouldBindWith(&form, binding.Form)	这俩个参数分别表示：绑定到哪个对象，bing.Form指定是Form类型数据（还可以是:json,ProtoBuf,xml,YAM等）
	// 也可以使用ShouldBind自动绑定
	if context.ShouldBind(&form) == nil {
		if form.Username == "admin" && form.Password == "123456" {
			// 这里gin.H 底层就是一个简单的map
			context.JSON(http.StatusOK, gin.H{"status": "登录成功"}) // StatusOk状态码 200
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "账号或密码错误"}) // StatusUnauthorized状态码 401
		}
	}

}
