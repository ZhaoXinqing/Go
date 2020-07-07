##cgroups来源
cgroups，其名称源自控制组群（control groups）的简写，是Linux内核的一个功能，用来限制、控制与分离一个进程组的资源（如**CPU、内存、磁盘输入输出**等）。这个项目最早是由Google的工程师（主要是Paul Menage和Rohit Seth）在2006年发起，最早的名称为进程容器（process containers）。在2007年时，因为在Linux内核中，容器（container）这个名词有许多不同的意义，为避免混乱，被重命名为cgroup
###功能作用
Cgroups最初的目标是为资源管理提供的一个统一的框架，既整合现有的cpuset等子系统，也为未来开发新的子系统提供接口。现在的cgroups适用于多种应用场景，从单个进程的资源控制，到实现操作系统层次的虚拟化（OS Level Virtualization）。<br>

Cgroups提供了以下功能：

1. 限制进程组可以使用的资源数量（Resource limiting ）。比如：memory子系统可以为进程组设定一个memory使用上限，一旦进程组使用的内存达到限额再申请内存，就会触发OOM（out of memory）。
1. 进程组的优先级控制（Prioritization ）。比如：可以使用cpu子系统为某个进程组分配特定cpu share。
1. 记录进程组使用的资源数量（Accounting ）。比如：可以使用cpuacct子系统记录某个进程组使用的cpu时间
1. 进程组隔离（Isolation）。比如：使用ns子系统可以使不同的进程组使用不同的namespace，以达到隔离的目的，不同的进程组有各自的进程、网络、文件系统挂载空间。
1. 进程组控制（Control）。比如：使用freezer子系统可以将进程组挂起和恢复。

##namespace 的概念
namespace 是 Linux 内核用来隔离内核资源的方式。通过 namespace 可以让一些进程只能看到与自己相关的一部分资源，而另外一些进程也只能看到与它们自己相关的资源，这两拨进程根本就感觉不到对方的存在。具体的实现方式是把一个或多个进程的相关资源指定在同一个 namespace 中。

Linux namespaces 是对全局系统资源的一种封装隔离，使得处于不同 namespace 的进程拥有独立的全局系统资源，改变一个 namespace 中的系统资源只会影响当前 namespace 里的进程，对其他 namespace 中的进程没有影响。
###namespace 的用途
可能绝大多数的使用者和我一样，是在使用 docker 后才开始了解 linux 的 namespace 技术的。实际上，Linux 内核实现 namespace 的一个主要目的就是实现轻量级虚拟化(容器)服务。在同一个 namespace 下的进程可以感知彼此的变化，而对外界的进程一无所知。这样就可以让容器中的进程产生错觉，认为自己置身于一个独立的系统中，从而达到隔离的目的。也就是说 linux 内核提供的 namespace 技术为 docker 等容器技术的出现和发展提供了基础条件。

### 查看进程所属的 namespace
从版本号为 3.8 的内核开始，/proc/[pid]/ns 目录下会包含进程所属的 namespace 信息，使用下面的命令可以查看当前进程所属的 namespace 信息：

1. 我们可以通过 clone() 在创建新进程的同时创建 namespace
1. 通过 setns() 函数可以将当前进程加入到已有的 namespace 中
1. 通过 unshare 函数可以在原进程上进行 namespace 隔离

###总结
namespace 是 linux 内核提供的特性，为虚拟化而生。随着 docker 的诞生引爆了容器技术，也把长期在后台默默奉献的 namespace 技术推到了大家的面前。

