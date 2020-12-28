## api：
      - 应用项目接口（Application Program Interface)
      - 提供一个软件到另一个的连接；
      - 由请求和响应构成；

## REST：
    资源（Resources）
       - 网络上的一个实体（一段文本、一张图片、一首歌曲、一种服务），要获取这个资源，访问这个资源对应的URI就可以；
       - URI只代表资源的实体，不代表它的形式。只代表"资源"的位置；
       
    表现层（Representation）
       - "资源"具体呈现出来的形式（文本txt格式、HTML、XML、JSON、二进制格式，图片JPG格式、PNG格式）
       - 它的具体表现形式，在HTTP请求的头信息中用 `Accept` 和 `Content-Type` 字段指定，这两个字段才是对"表现层"的描述。
    
    状态转化（State Transfer）
       - 客户端和服务器的一个互动过程，涉及到的数据和状态的变化。
            
    综述：
    （1）每一个URI代表一种资源；
    （2）客户端和服务器之间，传递这种资源的某种表现层；
    （3）客户端通过四个HTTP动词，对服务器端资源进行操作，实现"表现层状态转化"。
        
## HTTP methods：
    get:获取特定资源
    post:提交……
    put：更新……
    delete：删除……

    head：不返回body的get
    options：返回支持的http方法
    patch:更新部分资源

RESTful架构一种互联网软件架构模式，它结构清晰、符合标准、易于理解、扩展方便；

