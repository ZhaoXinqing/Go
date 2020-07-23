## k8s是什么? ##

是一套管理系统，对Docker及容器进行更高级更灵活的管理。就是基于容器的集群管理平台，它的全称是kubernetes。
一个K8S系统，通常称为一个K8S集群（Cluster）。这个集群主要包括两个部分：一个Master节点（主节点）和一群Node节点（计算节点）
Master节点主要还是负责管理和控制。Node节点是工作负载节点，里面是具体的容器。

## Kubernetes 组件 ## 
1 Master 组件
    1.1 kube-apiserver
    1.2 ETCD
    1.3 kube-controller-manager
    1.4 cloud-controller-manager
    1.5 kube-scheduler
    1.6 插件 addons
	    1.6.1 DNS
	    1.6.2 用户界面
	    1.6.3 容器资源监测
	    1.6.4 Cluster-level Logging
2 节点（Node）组件
    2.1 kubelet
    2.2 kube-proxy
    2.3 docker
    2.4 RKT
    2.5 supervisord
    2.6 fluentd
[http://docs.kubernetes.org.cn/230.html](http://docs.kubernetes.org.cn/230.html)

**kubectl** 用于运行Kubernetes集群命令的管理工具
Kubernetes最初源于谷歌内部的大规模集群管理系统**Borg**，提供了面向应用的容器集群部署和管理系统，负责对谷歌内部很多核心服务的调度和管理。
Borg的目的是让用户能够不必操心资源管理的问题，让他们专注于自己的核心业务，并且做到跨多个数据中心的资源利用率最大化。

## Pod ##
是kubernetes对象模型中可部署的最小对象，一个Pod代表集群上正在运行的一个进程。
一个Pod封装一个应用容器（也可以有多个容器），存储资源、一个独立的网络IP以及管理控制容器运行方式的策略选项。
Pod代表部署的一个单位：Kubernetes中单个应用的实例，它可能由单个容器或多个容器共享组成的资源。
Pod中的容器在集群中Node上被自动分配，容器之间可以共享资源、网络和相互依赖关系，并同时被调度使用。

**Pods提供两种共享资源：网络和存储**
网络：每个Pod被分配一个独立的IP地址，Pod中的每个容器共享网络命名空间，包括IP地址和网络端口。Pod内的容器可以使用localhost相互通信。当Pod中的容器与Pod 外部通信时，他们必须协调如何使用共享网络资源（如端口）。
存储：Pod可以指定一组共享存储volumes。Pod中的所有容器都可以访问共享volumes，允许这些容器共享数据。volumes 还用于Pod中的数据持久化，以防其中一个容器需要重新启动而丢失数据。有关Kubernetes如何在Pod中实现共享存储的更多信息，请参考Volumes。
Controller可以创建和管理多个Pod，提供副本管理、滚动升级和集群级别的自愈能力。例如，如果一个Node故障，Controller就能自动将该节点上的Pod调度到其他健康的Node上。

## Kubernetes三大落地姿势，你pick谁？ ##
https://www.kubernetes.org.cn/7774.html

## 发布：“2020行业云原生应用报告指南”，传统企业面临的又一次关键技术抉择 ##
https://www.kubernetes.org.cn/7862.html
