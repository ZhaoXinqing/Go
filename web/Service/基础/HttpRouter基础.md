# HttpRouter

#### 1、什么是httprouter

`httprouter`谈不上是一个框架，它的性能非常优秀。完全可以用来代替默认的 `ServeMux`。所以比较适合对性能要求高，路由相对简单扁平的业务，或者作为高层HTTP框架的内部模块。

httproute内部通过实现一个trie树来提高性能。核心代码就是golang标准库中 http.Handler 接口，在该函数中实现自己的请求路由分发策略。

> gorilla/mux也是一个第三方的路由器，但是性能一般

#### 2、为什么用httprouter

##### 说说net/http的不足：

- 不能单独的对请求方法(POST,GET等)注册特定的处理函数
- 不支持路由变量参数
- 不能自动对Path进行校准
- 性能一般、扩展性不足

##### httprouter带来哪些便利？

- 仅支持`精确匹配`，及只匹配一个模式或不会匹配到任何模式。相对于其他一些mux，例如go原生的 [http.ServerMux](http://golang.org/pkg/net/http/#ServeMux), 会使得一个请求URL匹配多个模式，从而需要有优先级顺序，例如最长匹配、最先匹配等等。
- 不需要关心URL`结尾的斜杠`
- 路径自动归一化和矫正
- 零内存分配、高性能

#### 3、httprouter路由匹配

##### net/http的路由匹配

```go
//	/api,可以访问到
//	/api/,不可以
http.HandleFunc("/api",func(w http.ResponseWriter,r *http.Request){
	fmt.Println("/api")
})

//	/api,可以
//	/api/,也可以
http.HandleFunc("/api/",func(w http.ResponseWriter,r *http.Request){
	fmt.Println("/api")
})
```

##### httprouter的路径匹配规则

两个路由命名捕获方式：（==路由命名的捕获，不是路由参数的捕获==）

- `:name` named parameter (只能根据路由命名进行捕获)
- `*name` catch-all parameter (*表示捕获任意内容)

其中`:name`的捕获方式是匹配内容直到下一个斜线或者路径的结尾。例如要为如下路径注册 handler:

```go
Path: /blog/:category/:post  
router.GET("/blog/:category/:post", Hello) //（category/post可以看成是一个变量）
	
当请求路径为：
/blog/go/request-routers            match: category="go", post="request-routers"
/blog/go/request-routers/           no match, but the router would redirect
/blog/go/                           no match
/blog/go/request-routers/comments   no match
```

`*name`的捕获方式是从指定位置开始 (包含前缀 "/") 匹配到结尾：

```go
Path: /files/*filepath     */
router.GET("/files/*filepath", Hello) //filepath可以看成是一个变量

当请求路径为：
/files/                             match: filepath="/"
/files/LICENSE                      match: filepath="/LICENSE"
/files/templates/article.html       match: filepath="/templates/article.html"
/files                              no match, but the router would redirect
```

#### 4、httprouter重定向

```go
func New() *Router {
	return &Router{
		//是否启用自动重定向。
        //比如请求/foo/,但是只有/foo的handler，则httprouter内部使用301重定向到/foo进行GET处理。
        //查看浏览器可以看到301的请求，可以看到请求头的location:/foo
        
        //直接get，然后响应一个字符串，这里的文本类型是text/plain，其实就是txt文件
        //这也就为什么，只有/foo的handler时，请求/foo/会404，因为/foo/代表着foo目录，默认会请求/foo/index.html，不就404了
        RedirectTrailingSlash:  true,
        
	RedirectFixedPath:      true,
		
        HandleMethodNotAllowed: true,
		
        //如果启用，则路由器会自动回复OPTIONS请求
        HandleOPTIONS:          true,
	}
}
```

#### 5、lookup

```go
func (r *Router) Lookup(method, path string) (Handle, Params, bool)
```

Lookup根据method+path检索对应的Handle，以及Params，并可以通过第三个返回值判断是否会进行重定向。

#### 6、关于参数

除了路由参数，还有URL的`query`参数，也就是`?a=b&c=d`这样的格式

##### query参数

```go
//url参数解析
http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
    values := r.URL.Query()
    fmt.Println(values)
})

//result
//http://127.0.0.1:9090/api/?a=1&b=2&c=3
//map[a:[1] b:[2] c:[3]]
复制代码
```

##### 路由参数

```go
//使用场景有哪些呢？？
/users/123
/users/456
/users/23456
比如有很多用户，如上，我们需要为一个个用户去注册路由？显然不现实
这时就可以使用路由参数了，/users/id
复制代码
```

