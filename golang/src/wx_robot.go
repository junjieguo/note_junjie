package main

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)
type WXMSG struct {
	Encrypt string `form:"Encrypt" json:"Encrypt" xml:"Encrypt" binding:"required"`
	Name string `form:"name" json:"name" xml:"name" binding:"required"`	
}
var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})
	r.GET("/sky/:name/*action", func(c *gin.Context){
		name:=c.Param("name")
		action:=c.Param("action")
		message:=name + " is " + action
		c.String(http.StatusOK,"Hello %s",message)
	})

	r.POST("/junjie",func(c *gin.Context){
		var param  WXMSG
		//id :=c.Query("id")
		//info_ms:=c.BindXML(&param )
		//<xml> 
  		//<name>skyguo</name>
  		// <Encrypt>1212121212121</Encrypt>
		//	</xml>
		//
		
		c.BindXML(&param )
		name:=c.PostForm("name")
		password:=c.PostForm("password")
		fmt.Printf("name is %s,message is %s Encrypt %v ",name,password,param.Name)
		//c.String(http.StatusOK,"id %s Hello %s ,password %s ",id,name,password)
		//c.JSON(200,gin.H{"result",param.Encrypt })
		c.JSON(200, gin.H{"result": param})
	})
	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":80")
}
