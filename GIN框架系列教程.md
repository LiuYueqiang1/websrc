# net./http协议

## GET请求示例

使用`net/http`包编写一个简单的发送HTTP请求的Client端，代码如下

```go
resp,err:=http.Get("https://www.liwenzhou.com/")
defer resp.Body.Close()
body,err:=ioutil.ReadAll(resp.Body)//[]byte
fmt.Println(string(body))//在终端打印liwenzhou.com的全部内容
```

## server示例

```go
// http server

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello 沙河！")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
```

# GIN框架系列教程

![image-20230428203023506](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230428203023506.png)



浏览器、APP发送请求给服务器

服务器把响应回复回来

第一个GIN示例：

```go
package main
import (
	"github.com/gin-gonic/gin"
)
//创建路由引擎
//GET方法
//启动服务
func main() {
	// 创建一个默认的路由引擎,引导别人访问什么地址，执行什么函数
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{ //json格式的
			"message": "Hello world!",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}

```

```go
r:=gin.Default()
r.GET("/hello",func(c *gin.Context){
    c.JSON(200,gin.H{
        "message","Helloworld!",
    })
})
r.POST("/hello",func(c *gin.Context){
    s.JSON(http.StatusOK,gin.H{
        "message","post"
    })
})
r.Run()
```

# HTML渲染

http/template

## 初始化

创建模板、解析模板、渲染模板

我们按照Go模板语法定义一个`hello.tmpl`的模板文件，内容如下：

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Hello</title>
</head>
<body>
    <p>Hello {{.}}</p>
</body>
</html>
```

然后我们创建一个`main.go`文件，在其中写下HTTP server端代码如下

```go
package main

import (
   "fmt"
   "html/template"
   "net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
   //解析模板
   tmpl, err := template.ParseFiles("F:\\goland\\go_project\\go_web\\websrc\\web_04\\hello.tmpl")
   if err != nil {
      fmt.Println("create tmplate failed,err:", err)
      return
   }
   //渲染模板
   tmpl.Execute(w, "沙河")
}
func main() {
   http.HandleFunc("/", sayHello)
   err := http.ListenAndServe(":9090", nil)
   if err != nil {
      fmt.Println("HTTP server is failed", err)
      return
   }
}
```

## 输入结构体：

```go
package main

import (
   "fmt"
   "html/template"
   "net/http"
)

type Student struct {
   Name    string
   Age     int
   Address string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
   //解析模板
   tmpl, err := template.ParseFiles("F:\\goland\\go_project\\go_web\\websrc\\web_05\\hello.tmpl")
   if err != nil {
      fmt.Println("create tmplate failed,err:", err)
      return
   }
   student := Student{
      Name:    "沙河",
      Age:     18,
      Address: "beijing",
   }
   stu := map[string]string{
      "name": "ss",
      "addr": "newyork",
   }
   hobbyLise := []string{
      "闪电五连鞭",
      "偷袭",
      "大意了啊",
   }
   //渲染模板
   tmpl.Execute(w, map[string]interface{}{
      "u1":    student,
      "s1":    stu,
      "hobby": hobbyLise,
   })
}
func main() {
   http.HandleFunc("/", sayHello)
   err := http.ListenAndServe(":9090", nil)
   if err != nil {
      fmt.Println("HTTP server is failed", err)
      return
   }
}
```

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Hello</title>
</head>
<body>
    <hr>
    <p>u1</p>
    <p>Hello {{- .u1.Name -}}</p>
    <p>年龄 {{ .u1.Age }}</p>
    <p>地址 {{.u1.Address}}</p>
    <hr>
    {{with .s1 }}
    <p>s1</p>
    <p>Hello {{ .name }}</p>
    <p>地址 {{ .addr}}</p>
    {{end }}
    <hr>
    {{if lt .u1.Age 22}}
    好好上学
    {{else}}
    好好工作
    {{end}}
    <hr>
    {{ range $idx,$hobby := .hobby}}
    <p>{{$idx}} - {{$hobby}}</p>
    {{end}}
</body>
</html>
```

## 创建模板、解析模板、渲染模板

```go
package main

import (
   "fmt"
   "html/template"
   "net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
   k := func(name string) (string, error) {
      return name + "年轻又帅气", nil
   }
   //创建模板
   t := template.New("f.tmpl")
   t.Funcs(template.FuncMap{
      "kua": k,
   })
   //解析模板
   tmpl, err := t.ParseFiles("F:\\goland\\go_project\\go_web\\websrc\\web_06\\f.tmpl")
   if err != nil {
      fmt.Println("create tmplate failed,err:", err)
      return
   }
   //渲染模板
   t.Execute(w, "帅气")    //渲染一次  这两个都可以
   tmpl.Execute(w, "沙河") //渲染一次
}
func demo1(w http.ResponseWriter, r *http.Request) {
   //*******加载文件的先后顺序十分重要*******
   t, err := template.ParseFiles("F:\\goland\\go_project\\go_web\\websrc\\web_06\\t.tmpl", "F:\\goland\\go_project\\go_web\\websrc\\web_06\\u1.tmpl")
   if err != nil {
      fmt.Printf("parse tmplate failed,err:%v\n", err)
      return
   }
   name := "小王子"
   t.Execute(w, name)
}
func main() {
   http.HandleFunc("/", f1)
   http.HandleFunc("/tmplDemo", demo1)
   err := http.ListenAndServe(":9000", nil)
   if err != nil {
      fmt.Println("HTTP server failed,err:%v", err)
      return
   }
}
```

# lesson 9 将网站中输入的变为HTML

## 加载文件、自定义函数



```go
//加载文件：
r.LoadHTMLFiles("./websrc/web_09/templates/index.tmpl")
r.LoadHTMLGlob("./websrc/web_09/templates/**/*")

//自定义函数
	//gin 框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
        //函数名：匿名函数
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
```



```go
func main() {
   r := gin.Default()
   //r.LoadHTMLFiles("./websrc/web_09/templates/index.tmpl") //模板解析
   // 静态文件
   r.Static("/static", "./websrc/web_09/statics")
   //gin 框架中给模板添加自定义函数
   r.SetFuncMap(template.FuncMap{
      "safe": func(str string) template.HTML {
         return template.HTML(str)
      },
   })
   r.LoadHTMLGlob("./websrc/web_09/templates/**/*")
   r.GET("posts/index", func(c *gin.Context) {
      c.HTML(http.StatusOK, "index.tmpl", gin.H{ //渲染模板
         "title": "liwenzhou.com",
      })
   })
   r.GET("users/posts", func(c *gin.Context) {
      c.HTML(http.StatusOK, "posts.tmpl", gin.H{ //渲染模板
         "title": "<a href='http://liwenzhou.com'>liwenzhou的博客</a>",
      })
   })
   r.Run(":9090")
}
```

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    {{/*css文件*/}}
    <link rel="stylesheet" href="/static/index.css">
    <title>users/posts</title>
</head>
<body>
<div>{{ .title | safe }}</div>
{{/*js文件*/}}
<script src = "/static/index.js"></script>
</body>
</html>
```

# lesson 10 利用前端文件加载后台

```go
func main() {
   r := gin.Default()
   // 静态文件
   r.Static("/assets", "F:\\goland\\go_project\\go_web\\modle\\assets")
   //gin 框架中给模板添加自定义函数
   
   //r.LoadHTMLFiles("F:\\goland\\go_project\\go_web\\modle\\index.html") //模板解析
   r.LoadHTMLGlob("F:\\goland\\go_project\\go_web\\modle\\templates/*")

   r.GET("index", func(c *gin.Context) {
      c.HTML(http.StatusOK, "index.html", nil)
   })
   r.GET("widgets.html", func(c *gin.Context) {
      c.HTML(http.StatusOK, "widgets.html", nil)
   })
   r.GET("index.html", func(c *gin.Context) {
      c.HTML(http.StatusOK, "index.html", nil)
   })
   r.GET("index2.html", func(c *gin.Context) {
      c.HTML(http.StatusOK, "index2.html", nil)
   })
   r.Run(":9090")
}
```

# lesson 10 map和结构体

```go
gin.H就是map类型
```



```go
func main() {
   r := gin.Default()
   r.GET("/json", func(c *gin.Context) {
      // 方法1：使用map
      //data := map[string]interface{}{
      // "name":    "小王子",
      // "age":     18,
      // "message": "hello world!",
      //}
      data := gin.H{"name": "小王子", "age": 18, "message": "hello world!"}
      //H is a shortcut for map[string]interface{}
      c.JSON(http.StatusOK, data)
   })

   // 方法2：结构体，灵活使用tag对结构体字段做定制化操作
   type msg struct {
      Name    string `json:"name"`
      Message string `json:"message"`
      Age     int    `json:"age"`
   }
   r.GET("/newjson", func(c *gin.Context) {
      data := msg{
         "马老师",
         "我打一个连五鞭,发生甚摸事了",
         69,
      }
      c.JSON(http.StatusOK, data)
   })
   r.Run()
}
```

lesson 11 从网址获取内容打印到网页中

```go
func main() {
   r := gin.Default()
   r.GET("/web", func(c *gin.Context) {
      name := c.Query("query")
      age := c.Query("age")
      //name := c.DefaultQuery("query", "somebody") //取不到就用指定的默认值

      //name, ok := c.GetQuery("query")
      //if !ok {
      // //取不到
      // name = "somebody"
      //}

      c.JSON(http.StatusOK, gin.H{
         "name": name,
         "age":  age,
      })
   })
   r.Run()
}
```

```
//http://127.0.0.1:8080/web?query=筱往&age=18
```

# lesson 12 表单输入用户名密码后打印

```go
c.PostForm()
```



```go
// 将 网页中输入框中的文件密码取出，打印到index.html网页中
func main() {
   r := gin.Default()
   r.LoadHTMLFiles("F:\\goland\\go_project\\go_web\\websrc\\web_12\\login.html", "F:\\goland\\go_project\\go_web\\websrc\\web_12\\index.html")
   r.GET("/login", func(c *gin.Context) {
      c.HTML(http.StatusOK, "login.html", nil)
   })
   // /login post
   r.POST("/login", func(c *gin.Context) {
      username := c.PostForm("username")         //根据html文件得到用户名
      password := c.PostForm("password")         //根据html文件得到密码
      c.HTML(http.StatusOK, "index.html", gin.H{ //加载文件
         "Name":     username,
         "Password": password,
      })
   })
   r.Run(":9090")
}
```

```html
login.html
<!DOCTYPE html>  
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>login</title>
</head>
<body>
<form action="/login" method="post" novalidate autocomplete="off">
    <label for="username">username:</label>
  <div><input type="text" name="username" id ="username"></div>
    <label for="password">password:</label>
   <div><input type="password" name="password" id ="password"></div>

    <input type="submit" value="登录">
</form>

</body>
</html>
```

```html
index.html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>index</title>
</head>
<body>
<h1>Hello,{{.Name}}!</h1>
<p>你的密码是：{{.Password}}</p>
</body>
</html>
```

# lesson 13 从网页中加载年月信息

```go
// 从网页中读取年月信息，或者名字，年龄等信息
func main() {
   r := gin.Default()
   r.GET("/user/:name/:age", func(c *gin.Context) {
      // 获得路径参数
      name := c.Param("name")
      age := c.Param("age")

      c.JSON(http.StatusOK, gin.H{
         "name": name,
         "age":  age,
      })
   })
   //http://127.0.0.1:9090/blog/2022/12
   //{"month":"12","year":"2022"}
   r.GET("blog/:year/:month", func(c *gin.Context) {
      year := c.Param("year")
      month := c.Param("month")
      c.JSON(http.StatusOK, gin.H{
         "year":  year,
         "month": month,
      })
   })
   r.Run(":9090")
}
```

# lesson 14

## 1、读取网页上的username和password

```go
type UserInfo struct {
   Username string `form:"username" json:"username"`
   Password string `form:"password" json:"password"`
}

func main() {
   r := gin.Default()
   r.LoadHTMLFiles("F:\\goland\\go_project\\go_web\\websrc\\web_14\\index.html")
   r.GET("/user", func(c *gin.Context) {
      //username := c.Query("username")
      //password := c.Query("password")
      //u := UserInfo{
      // username: username,
      // password: password,
      //}
      var u UserInfo
      err := c.ShouldBind(&u)
      if err != nil {
         c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
         })
      } else {
         fmt.Printf("%#v\n", u)
         c.JSON(http.StatusOK, gin.H{
            "status": "ok",
         })
      }
   })
   r.Run()
```

```
//http://127.0.0.1:8080/user?username=qimi&password=111
```

## 2、访问http://127.0.0.1:8080/index

写入form表单中username和password后给post请求

```html
index.html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>login</title>
</head>
<body>
{{/*发送post请求，发送到gin框架的form表单中*/}}
<form action="/form" method="post" novalidate autocomplete="off">
    <label for="username">username:</label>
    <div><input type="text" name="username" id ="username"></div>
    <label for="password">password:</label>
    <div><input type="password" name="password" id ="password"></div>

    <input type="submit" value="登录">
</form>
</body>
</html>
```

```go
func main() {
r := gin.Default()
	r.LoadHTMLFiles("F:\\goland\\go_project\\go_web\\websrc\\web_14\\index.html")
r.GET("/index", func(c *gin.Context) {
   c.HTML(http.StatusOK, "index.html", nil)
})

r.POST("/form", func(c *gin.Context) {
   var u UserInfo
   err := c.ShouldBind(&u)
   if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{
         "error": err.Error(),
      })
   } else {
      fmt.Printf("%#v\n", u)
      c.JSON(http.StatusOK, gin.H{
         "status": "ok",
      })
   }
})
r.Run()
}
```

## 3、使用postman发送 在终端可以接收到数据

# lesson 15 文件上传

```go
c.FormFile("f1")
//读取成功的话，把文件加载到
 dst := path.Join("F:\\goland\\go_project\\go_web\\websrc\\web_15", f.Filename)
         c.SaveUploadedFile(f, dst)
```



```go
func main() {
   r := gin.Default()
   r.LoadHTMLFiles("F:\\goland\\go_project\\go_web\\websrc\\web_15\\index.html")
   r.GET("/index", func(c *gin.Context) {
      c.HTML(http.StatusOK, "index.html", nil)
   })
   r.POST("/upload", func(c *gin.Context) {
       //FormFile返回所提供表单键的第一个文件
      f, err := c.FormFile("f1")
      if err != nil {
         c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
         })
         return
      } else {
         // dst := fmt.Sprintf("F:\\goland\\go_project\\go_web\\websrc\\web_15\\%s", f.Filename)
         dst := path.Join("F:\\goland\\go_project\\go_web\\websrc\\web_15", f.Filename)
         c.SaveUploadedFile(f, dst)
         c.JSON(http.StatusBadRequest, gin.H{
            "status": "OK",
         })
      }
   })
   r.Run()
}
```

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>上传文件示例</title>
</head>
<body>
{{/*发送post请求，发送到gin框架的upload表单中*/}}
{{/*加后面的enctype enctype="multipart/form-data" 才能上传成功，不加也行，别加novalidate autocomplete="off" 就行*/}}
<form action="/upload" method="post"  >
    <input type="file" name="f1">
    <input type="submit" value="上传">
</form>
</body>
</html>
```

# lesson 16 网页重定向

```go
c.Redirect(http.StatusMovePermanently,"http://www.baidu.com")
c.Request.URL.Path="/b"
c.HandleContext(c)
```

```go
func main() {
   r := gin.Default()
   r.GET("/index", func(c *gin.Context) {
      //c.JSON(http.StatusOK, gin.H{
      // "message": "ok",
      //})
      c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
   })

   r.GET("/a", func(c *gin.Context) {
      c.Request.URL.Path = "/b"
      r.HandleContext(c)
   })
   r.GET("/b", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
         "message": "ok",
      })
   })
   r.Run()
}
```

# lesson 17 路由组

```go
//加载任何方法
r.Any("/home",func(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{
        "method":"PUT"
    })
})
//路由组的嵌套
shopGroup:=r.Group("/shop")
{
    shopGroup.GET("/index",func(c *gin.Context),nil){
	        
    }
}
```



```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

func main() {
   r := gin.Default()
   r.GET("/index", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
         "method": "GET",
      })
   }) //获取信息
    
   r.POST("/index", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
         "method": "POST",
      })
   }) //创建操作
   
    r.DELETE("/index", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
         "method": "DELETE",
      })
   }) //
   
    r.PUT("/index", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
         "method": "PUT",
      })
   }) //更新操作

   r.Any("/home", func(c *gin.Context) {
      switch c.Request.Method {
      case "GET":
         c.JSON(http.StatusOK, gin.H{"method": "GET"})
      case http.MethodPost:
         c.JSON(http.StatusOK, gin.H{"method": "POST"})
      }
   })

   //404
   r.NoRoute(func(c *gin.Context) {
      c.JSON(http.StatusNotFound, gin.H{"mes": "404 not found"})
   })
   //路由组
   videoGroup := r.Group("/video")
   {
      videoGroup.GET("/index", func(c *gin.Context) {
         c.JSON(http.StatusOK, gin.H{"message": "/video/index"})
      })
      videoGroup.GET("/xx", func(c *gin.Context) {
         c.JSON(http.StatusOK, gin.H{"message": "/video/xx"})
      })
      videoGroup.GET("/ee", func(c *gin.Context) {
         c.JSON(http.StatusOK, gin.H{"message": "/video/ee"})
      })
   }
   //嵌套路由组
   shopGroup := r.Group("/shop")
   {
      shopGroup.GET("/index", func(c *gin.Context) {
         c.JSON(http.StatusOK, gin.H{"messsage": "shop/index"})
      })
      //嵌套
      xx := shopGroup.Group("/xx")
      {
         xx.GET("/oo", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"message": "shop/xx/oo"})
         })
      }
   }
   r.Run()
}
```

# lesson 18 中间件

浏览器 -路由（处理函数）

认证了就登陆，否则打回去

![1.png](https://github.com/LiuYueqiang1/websrc/blob/main/web_18/1.png?raw=true)

![2.png](https://github.com/LiuYueqiang1/websrc/blob/main/web_18/2.png?raw=true)

![3.png](https://github.com/LiuYueqiang1/websrc/blob/main/web_18/3.png?raw=true)

```go
package main

// gin中间件
import (
   "fmt"
   "github.com/gin-gonic/gin"
   "net/http"
   "time"
)

func indexHandler(c *gin.Context) {
   fmt.Println("index")
   c.JSON(http.StatusOK, gin.H{
      "msg": "index",
   })
}

// 定义一个gin中间件 ，统计请求处理函数的耗时
func m1(c *gin.Context) {
   fmt.Println("m1 in ...")
   //计时
   start := time.Now()
   c.Next()                  //调用后续的处理函数
   cost := time.Since(start) //统计所有处理函数花费的时间
   fmt.Println("cost:", cost)
   fmt.Println("m1 out ...")
}
func m2(c *gin.Context) {
   fmt.Println("m2 in ...")
   //c.Next() //调用后续的处理函数
   c.Abort() //阻止调用后续函数
   fmt.Println("m2 out ...")
   // m1 in ...
   // m2 in ...
   // m2 out ...
   //cost: 711.6µs
   // m1 out ...

   //return
   //m1 in ...
   //m2 in ...
   //cost: 1.2252ms
   //m1 out ...
}
func main() {
   r := gin.Default()
   //r.GET("/index", m1, indexHandler)
   r.Use(m1, m2) //全局注册中间件m1,m2
   r.GET("/index", indexHandler)
   r.GET("/shop", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
         "msg": "shop",
      })
   })
   r.GET("/user", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
         "msg": "user",
      })
   })
   r.Run(":9090")
}
```



# lesson 19 Gorm

导包

```go
   "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/mysql"
```

创建数据库db1

```go
// 连接mysql数据库 ，数据库的用户名、密码
db, err := gorm.Open("mysql", "root:961024@tcp(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
if err != nil {
	fmt.Println("init db failed!,err:", err)
	return
}
defer db.Close()
```

创建表

```go
//创建了一个名字为user_infos的表
	db.AutoMigrate(&UserInfo{})
```

创建记录

```go
	//在 user_infos的表 创建记录
	db.Create(&u1)
```

查询、更新、删除

```go
	//更新
	db.Model(&u).Update("hobby", "唱跳rap")
	//删除
	db.Delete(&u) //故此处删掉的u为 ID：01的
```



```go
import (
   "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
    // 连接mysql数据库 ，数据库的用户名、密码
	db, err := gorm.Open("mysql", "root:961024@tcp(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("init db failed!,err:", err)
		return
	}
	defer db.Close()

	//自动迁移
	//创建了一个名字为user_infos的表
	db.AutoMigrate(&UserInfo{}) //自动迁移为给定的模型运行自动迁移，只会添加缺失的字段，不会删除/更改当前数据

	u1 := UserInfo{1, "qimi", "男", "篮球"}
	u2 := UserInfo{2, "欧阳修", "男", "混元功法"}

	//在 user_infos的表 创建记录
	db.Create(&u1)
	db.Create(&u2)
	//查询
	var u = new(UserInfo)  //将结构体的地址赋给 u
	db.First(u)            //First find first record that match given conditions, order by primary key
	fmt.Printf("%#v\n", u) //**** &main.UserInfo{ID:0x1, Name:"qimi", Gender:"男", Hobby:"篮球"}

	var uu UserInfo //将结构体的数值赋给 uu
	db.Find(&uu, "hobby=?", "混元功法")
	fmt.Printf("%#v\n", uu) //**** main.UserInfo{ID:0x2, Name:"欧阳修", Gender:"男", Hobby:"混元功法"}

	//更新
	db.Model(&u).Update("hobby", "唱跳rap")
	//删除
	db.Delete(&u) //故此处删掉的u为 ID：01的
}
```

# lesson 20 Gorm

指定表名

```go
// 使用方法更改表名
// 将 User 的表名设置为 `profiles`
func (User) TableName() string {
	return "profiles"
}
func main() {
	db, err := gorm.Open("mysql", "root:961024@tcp(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//字段迁移
	db.AutoMigrate(&User{})
	//表名默认就是结构体名称的复数 users

	// 使用User结构体 创建 名为`deleted_users`的表
	db.Table("deleted_users").CreateTable(&User{})
}
```

# lesson 21 Gorm 增

```go
package main

import (
   "database/sql"
   "fmt"
   "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
   ID   int64
   Name string `gorm:"default:'马保郭'"`
   Age  int64
}

type User2 struct {
   ID   int64
   Name *string `grom:"default:'小王'"`
   Age  int64
}

type User3 struct {
   ID   int64
   Name sql.NullString `grom:"default:'小王3'"`
   Age  int64
}

func main() {
   db, err := gorm.Open("mysql", "root:961024@tcp(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
   if err != nil {
      panic(err)
   }
   defer db.Close()
   db.AutoMigrate(User{})  //新建表user
   db.AutoMigrate(User2{}) // 新建表 user2
   db.AutoMigrate(User3{})
   user := User{
      Name: "掌门人",
      Age:  69,
   }
   fmt.Println(db.NewRecord(user)) //查询主键是否为空  true
   db.Create(&user)                //在users表中创建一条记录
   fmt.Println(db.NewRecord(user)) //查询主键是否为空 false
   user1 := User{
      Name: "",
      Age:  73,
   }
   db.Create(&user1) //在users表中创建一条记录

   user2 := User2{
      Name: new(string),
      Age:  18,
   }
   db.Create(&user2)

   user3 := User3{
      Name: sql.NullString{"", true},
      Age:  22,
   }
   db.Create(user3)

   //查询操作 根据主键查询第一条记录
   db.First(&user)
   fmt.Println(&user)
   //
}
```

# lesson 25 小清单



```go
//创建表
DB.AutoMigrate(values ...interface{})
//加载静态文件配置
r.Static(relativePath string, root string)
r.Static("/static", "F:\\goland\\go_project\\go_web\\websrc\\web_25\\static")

//模板解析
r.LoadHTMLGlob(pattern string)
r.LoadHTMLGlob("F:\\goland\\go_project\\go_web\\websrc\\web_25\\tmlplates/*")

//创建路由组
r.Group(relativePath string, handlers ...HandlerFunc) *RouterGroup
v1Group := r.Group("v1")
//Group命令用于创建新的路由器组。您应该添加所有具有相同中间件或相同路径前缀的路由。
//例如，可以对使用公共中间件进行授权的所有路由进行分组。

// 将 页面上的值放到todo中
c.BindJSON(&todo)

// 2、存入数据
err = DB.Create(&todo).Error
if err!= nil {
}

// 拿到要修改的id
id, ok := c.Params.Get("id")
//Get返回与给定名称匹配的第一个Param的值和一个布尔值true。
//如果没有找到匹配的参数，则返回一个空字符串和一个布尔值false。
```



```go
package main

import (
   "github.com/gin-gonic/gin"
   "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/mysql"
   "net/http"
)

type ToDo struct {
   ID     int    `json:"id"`
   Title  string `json:"title"`
   Status bool   `json:"status"`
}

var DB *gorm.DB

// 初始化数据库
func initMySQL() (err error) {
   dsn := "root:961024@tcp(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
   DB, err = gorm.Open("mysql", dsn)
   if err != nil {
      return
   }
   return DB.DB().Ping()
}

func main() {
   //创建数据库
   err := initMySQL()
   if err != nil {
      panic(err)
   }
   defer DB.Close()
   //模型绑定
   DB.AutoMigrate(&ToDo{})
   //加载中间件
   r := gin.Default()
   //加载静态文件
   r.Static("/static", "F:\\goland\\go_project\\go_web\\websrc\\web_25\\static")

   //模板解析
   r.LoadHTMLGlob("F:\\goland\\go_project\\go_web\\websrc\\web_25\\tmlplates/*")

   r.GET("/", func(c *gin.Context) {
      c.HTML(http.StatusOK, "index.html", nil)
   })

   // v1
   v1Group := r.Group("v1")
   //Group命令用于创建新的路由器组。您应该添加所有具有相同中间件或相同路径前缀的路由。
   //例如，可以对使用公共中间件进行授权的所有路由进行分组。

   //待办事项
   //添加
   {
      v1Group.POST("/todo", func(c *gin.Context) {
         //前端页面填写待办事项 点击提交 发送请求到此处
         var todo ToDo
         // 将 页面上的值放到todo中
         c.BindJSON(&todo)
         // 1、 从请求中把数据拿出来
         // 2、存入数据
         //err = DB.Create(&todo).Error
         //if err!= nil {
         //}
         // 3、返回响应
         if err = DB.Create(&todo).Error; err != nil {
            c.JSON(http.StatusOK, gin.H{
               "error": err.Error(),
            })
         } else {
            c.JSON(http.StatusOK, gin.H{
               "code": 2000,
               "msg":  "success",
               "data": todo,
            })
         }
      })
      // 查看所有的待办事项
      v1Group.GET("/todo", func(c *gin.Context) {
         //查询todo这个表里的所有数据
         var todoList []ToDo
         if err = DB.Find(&todoList).Error; err != nil {
            c.JSON(http.StatusOK, gin.H{
               "error": err.Error(),
            })
         } else {
            c.JSON(http.StatusOK, todoList)
         }
      })
       
      // 修改某一个待办事项
      v1Group.PUT("/todo/:id", func(c *gin.Context) {
         // 拿到要修改的id
         id, ok := c.Params.Get("id")
         //Get返回与给定名称匹配的第一个Param的值和一个布尔值true。
         //如果没有找到匹配的参数，则返回一个空字符串和一个布尔值false。
         if !ok {
            c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
            return
         }
         // 根据id 匹配数据
         var todo ToDo
         if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
            c.JSON(http.StatusOK, gin.H{"error": err.Error()})
            return
         }
         // 修改
         c.BindJSON(&todo)
         if err = DB.Save(&todo).Error; err != nil {
            c.JSON(http.StatusOK, gin.H{"error": err.Error()})
         } else {
            c.JSON(http.StatusOK, todo)
         }
      })
       
      v1Group.DELETE("/todo/:id", func(c *gin.Context) {
         id, ok := c.Params.Get("id")
         if !ok {
            c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
            return
         }
         if err = DB.Where("id=?", id).Delete(ToDo{}).Error; err != nil {
            c.JSON(http.StatusOK, gin.H{"error": err.Error()})
         } else {
            c.JSON(http.StatusOK, gin.H{id: "deleted"})
         }
      })
   }

   r.Run()
}
```

# Gorm 增删改查

## 查询

### 一般查询

```go
// 根据主键查询第一条记录
db.First(&user)
//// SELECT * FROM users ORDER BY id LIMIT 1;
```

```go
// 查询所有的记录
db.Find(&users)
//// SELECT * FROM users;
```

```go
// 查询指定的某条记录(仅当主键为整型时可用)
db.First(&user, 10)
//// SELECT * FROM users WHERE id = 10;
```

### Where 条件

```go
// Get all matched records
db.Where("name = ?", "jinzhu").Find(&users)
//// SELECT * FROM users WHERE name = 'jinzhu';
```

```go
// IN
db.Where("name IN (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
//// SELECT * FROM users WHERE name in ('jinzhu','jinzhu 2');
```

### Struct & Map查询

```go
// Struct
db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;

// Map
db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;
```

## 更新

### 更新所有字段

`Save()`默认会更新该对象的所有字段，即使你没有赋值。

```go
db.First(&user)

user.Name = "七米"
user.Age = 99
db.Save(&user)

////  UPDATE `users` SET `created_at` = '2020-02-16 12:52:20', `updated_at` = '2020-02-16 12:54:55', `deleted_at` = NULL, `n
```

### 更新修改字段

如果你只希望更新指定字段，可以使用`Update`或者`Updates`

```go
db.Model(&user).Update("name","修改名字")//修改所有
db.Model(&user).Where("age = ? ",69).Update("name","修改名字")//修改指定
```

```go
db.Model(&user).Update("age",gorm.Expr("age + ? + ?" , 2 , 100))
```

## 删除

```go
//删除现有记录
db.Delete(&email)
```

```go
db.Where("email LIKE ?","%jinzhu%").Delete(Email{})
db.Where("age = ?",20).Delete(&user)
```



