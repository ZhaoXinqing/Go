## k8s是什么? ##

将Docker应用于具体的业务实现，是存在困难的，编排、管理和调度等各个方面，都不容易。于是，人们迫切需要一套管理系统，对Docker及容器进行更高级更灵活的管理。就在这个时候，K8S出现了。K8S，就是基于容器的集群管理平台，它的全称，是kubernetes。

## 技术特点和结构 ##

一个K8S系统，通常称为一个K8S集群（Cluster）。这个集群主要包括两个部分：一个Master节点（主节点）和一群Node节点（计算节点）
Master节点主要还是负责管理和控制。Node节点是工作负载节点，里面是具体的容器。

Master节点包括**API Server、Scheduler、Controller manager、etcd**。

- API Server是整个系统的对外接口，供客户端和其它组件调用，相当于“营业厅”。
- Scheduler负责对集群内部的资源进行调度，相当于“调度室”。
- Controller manager负责管理控制器，相当于“大总管”。

然后是Node节点。pod、Docker、kuberlet、kube-proxy、fluentd

