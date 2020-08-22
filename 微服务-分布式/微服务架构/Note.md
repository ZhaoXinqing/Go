go-kit本身不是一个框架，而是一套微服务工具集；其设计思想：分层设计，组件化，可扩展；
三层结构：Transport层，Endpoint层，Service层。需求是什么？如何在满足需求的同时(1 只做高性能的HTTP 接口；2 需要完整的单元测试体系；3 可扩展，组件化；)，让框架和系统具有一定的弹性。GoKit CLi，根据我们的需求自动生成service，transport和endpoint模版，以及生成供调用方的使用的client library

Etcd比较多的应用场景是用于服务发现，服务发现(Service Discovery)要解决的是分布式系统中最常见的问题之一，即在同一个分布式集群中的进程或服务如何才能找到对方并建立连接。从本质上说，服务发现就是要了解集群中是否有进程在监听upd或者tcp端口，并且通过名字就可以进行查找和链接。

dubbo-go 中异步网络 I/O 模型的设计，这部分将 Go 语言轻量级协程的优势体现了出来。

Service Mesh 确实是微服务未来发展的的一个大方向，但是现阶段在国内大公司还没有看到非常成功的案例，很多中小公司自身微服务还未拆分完毕甚至于还未开始，目前 dubbo-go 社区优先解决这种类型企业微服务技术落地环节中遇到的问题，专注于补齐相关功能、优化整体性能和解决 bug。至于未来，我相信随着 Dubbo Mesh 在 Service Mesh 领域的探索，dubbo-go 肯定会跟进并扮演重要角色。

dubbo-go 何鑫铭：
当前发布的 v1.0.0 版本支持的功能如下：
    角色：Consumer(√)、Provider(√)
    传输协议：HTTP(√)、TCP(√)
    序列化协议：JsonRPC v2(√)、Hessian v2(√)
    注册中心：ZooKeeper(√)
    集群策略：Failover(√)
    负载均衡：Random(√)
    过滤器：Echo Health Check(√)
    extension 扩展机制
该版本沿用了 Dubbo 的代码分层解耦设计。Dubbo 2.6.x 的主要功能都会逐渐在 dubbo-go 中实现，包括 Dubbo 基于 SPI 的代码拓展机制，dubbo-go 也有对应的 extension 扩展机制与之对应。
 dubbo-go 现在使用的 TCP 异步网络 I/O 库,下一版本我们会针对 dubbo-go 和 Getty 的网络 I/O 与线程派发这一部分进行进一步优化。

下一步支持 Dubbo 的另外几大重要功能，如：
    routing rule 路由规则(dubbo v2.6.x)
    dynamic configuration 动态配置中心(dubbo v2.7.x)
    metrics 指标与监控(dubbo v2.7.x) 
    trace 链路监控(dubbo ecos) 

dubbo-go 现在已经开始被一些企业尝试应用于 Go 语言应用融入企业已有 Java & Dubbo 技术栈，以及搭建全新 Go 语言分布式应用等场景。比如中通快递内部 Go 调用 Java Dubbo 服务；作为携程 Go 语言应用的服务框架以及 Go、Java 应用互通。