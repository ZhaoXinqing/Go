## k8s是什么? ##

是一套管理系统，对Docker及容器进行更高级更灵活的管理。就是基于容器的集群管理平台，它的全称是kubernetes。


## 技术特点和结构 ##

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

http://docs.kubernetes.org.cn/230.html
