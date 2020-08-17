Gin:
    - Gin是一个golang的微框架，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点
    - 对于golang而言，web框架的依赖要远比Python，Java之类的要小。自身的net/http足够简单，性能也非常不错
    - 借助框架开发，不仅可以省去很多常用的封装带来的时间，也有助于团队的编码风格和形成规范

    - Restful风格的API
        - gin支持Restful风格的API
        - 即Representational State Transfer的缩写。直接翻译的意思是"表现层状态转化"，
        - 是一种互联网应用程序的API设计理念：URL定位资源，用HTTP描述操作
            1.获取文章 /blog/getXxx Get blog/Xxx
            2.添加 /blog/addXxx POST blog/Xxx
            3.修改 /blog/updateXxx PUT blog/Xxx
            4.删除 /blog/delXxxx DELETE blog/Xxx

Gin 框架中文文档 2019 : https://learnku.com/docs/gin-gonic/2019