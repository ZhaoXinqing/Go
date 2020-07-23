## 分布式系统 & 微服务架构 ##
微服务是**架构设计方式**，分布式是**系统部署方式**；

- 微服务是指很小的服务（小功能，服务可单独部署运行）不同服务之间通过rpc调用。
- 分布式是指服务部署在不同的机器上，一个服务提供一个或多个功能，服务之间也是通过rpc来交互或是webservice来交互。

https://blog.csdn.net/qq_41582642/article/details/81016500
https://www.jianshu.com/p/1f9455139a31


## 目前系统应用架构演变
https://blog.csdn.net/qq_36874292/article/details/82290890
https://www.jianshu.com/p/bbfda62a771c

## 系统拆分原则
https://www.cnblogs.com/guanghe/p/10978349.html
https://youzhixueyuan.com/the-principles-and-steps-of-distributed-architecture-system-resolution.html

## 水平拆分 & 垂直拆分 ##
互联网时代谈论最多的话题就是拆分，一种**分而治之**的思想和逻辑。
**水平拆分：** 单一节点扩展为多节点，多节点具有一致的功能，组成一个服务池，节点共同处理大规模高并发的请求量。


**垂直拆分：** 按照功能进行拆分，把一个复杂的功能拆分为多个单一简单的功能，使得维护和变更都变得更简单、容易、安全。所以更易于产品版本的迭代，还能够快速的进行敏捷发布和上线。

## 架构方面的设计模式 ##
工作中用到的一些架构方面的设计模式。总体而言，共有八种，分别是：

- 单库单应用模式：最简单的，可能大家都见过
- 内容分发模式：目前用的比较多
- 查询分离模式：对于大并发的查询、业务
- 微服务模式：适用于复杂的业务模式的拆解
- 多级缓存模式：可以把缓存玩的很好
- 分库分表模式：解决单机数据库瓶颈
- 弹性伸缩模式：解决波峰波谷业务流量不均匀的方法之一
- 多机房模式：解决高可用、高性能的一种方法
https://www.cnblogs.com/ZenoLiang/p/10891499.html

## ServiceMesh(服务网格) ##
https://www.jianshu.com/p/27a742e349f7


