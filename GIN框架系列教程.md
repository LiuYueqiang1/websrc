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

# lesson 18 中间件

浏览器 -路由（处理函数）

认证了就登陆，否则打回去

![image-20230511211016163](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230511211016163.png)

简单的中间件

![image-20230511211844509](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230511211844509.png)

![image-20230511212205430](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230511212205430.png)

![image-20230511212944368](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230511212944368.png)



![image-20230511213030175](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230511213030175.png)

![image-20230511213313502](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230511213313502.png)

![image-20230511213326704](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230511213326704.png)



![image-20230511215703412](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230511215703412.png)

![image-20230511220049309](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230511220049309.png)
