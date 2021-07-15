package main

import (
	//"net/http"

	"github.com/gin-gonic/gin"
)

// var db = make(map[string]string)

// func setupRouter() *gin.Engine {
// 	// Disable Console Color
// 	// gin.DisableConsoleColor()
// 	r := gin.Default()

// 	// Ping test
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.String(http.StatusOK, "pong")
// 	})

// 	// Get user value
// 	r.GET("/user/:name", func(c *gin.Context) {
// 		user := c.Params.ByName("name")
// 		value, ok := db[user]
// 		if ok {
// 			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
// 		} else {
// 			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
// 		}
// 	})

// 	// Authorized group (uses gin.BasicAuth() middleware)
// 	// Same than:
// 	// authorized := r.Group("/")
// 	// authorized.Use(gin.BasicAuth(gin.Credentials{
// 	//	  "foo":  "bar",
// 	//	  "manu": "123",
// 	//}))
// 	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
// 		"foo":  "bar", // user:foo password:bar
// 		"manu": "123", // user:manu password:123
// 	}))

// 	/* example curl for /admin with basicauth header
// 	   Zm9vOmJhcg== is base64("foo:bar")

// 		curl -X POST \
// 	  	http://localhost:8080/admin \
// 	  	-H 'authorization: Basic Zm9vOmJhcg==' \
// 	  	-H 'content-type: application/json' \
// 	  	-d '{"value":"bar"}'
// 	*/
// 	authorized.POST("admin", func(c *gin.Context) {
// 		user := c.MustGet(gin.AuthUserKey).(string)

// 		// Parse JSON
// 		var json struct {
// 			Value string `json:"value" binding:"required"`
// 		}

// 		if c.Bind(&json) == nil {
// 			db[user] = json.Value
// 			c.JSON(http.StatusOK, gin.H{"status": "ok"})
// 		}
// 	})

// 	return r
// }

// func main() {
// 	r := setupRouter()
// 	// Listen and Server in 0.0.0.0:8080
// 	r.Run(":8080")
// }


// func main() {
// 	r := gin.Default()

// 	r.GET("/someJSON", func(c *gin.Context) {
// 		data := map[string]interface{}{
// 			"lang": "GO语言",
// 			"tag":  "<br>",
// 		}

// 		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
// 		c.AsciiJSON(http.StatusOK, data)
// 	})

// 	// Listen and serve on 0.0.0.0:8080
// 	r.Run(":8080")
// }

// type StructA struct {
//     FieldA string `form:"field_a"`
// }

// type StructB struct {
//     NestedStruct StructA
//     FieldB string `form:"field_b"`
// }

// type StructC struct {
//     NestedStructPointer *StructA
//     FieldC string `form:"field_c"`
// }

// type StructD struct {
//     NestedAnonyStruct struct {
//         FieldX string `form:"field_x"`
//     }
//     FieldD string `form:"field_d"`
// }

// func GetDataB(c *gin.Context) {
//     var b StructB
//     c.Bind(&b)
//     c.JSON(200, gin.H{
//         "a": b.NestedStruct,
//         "b": b.FieldB,
//     })
// }

// func GetDataC(c *gin.Context) {
//     var b StructC
//     c.Bind(&b)
//     c.JSON(200, gin.H{
//         "a": b.NestedStructPointer,
//         "c": b.FieldC,
//     })
// }

// func GetDataD(c *gin.Context) {
//     var b StructD
//     c.Bind(&b)
//     c.JSON(200, gin.H{
//         "x": b.NestedAnonyStruct,
//         "d": b.FieldD,
//     })
// }

// func main() {
//     r := gin.Default()
//     r.GET("/getb", GetDataB)
//     r.GET("/getc", GetDataC)
//     r.GET("/getd", GetDataD)

//     r.Run()
// }



type myForm struct {
    Colors []string `form:"colors[]"`
}

func formHandler(c *gin.Context) {
    var fakeForm myForm
    c.ShouldBind(&fakeForm)
    c.JSON(200, gin.H{"color": fakeForm.Colors})
}