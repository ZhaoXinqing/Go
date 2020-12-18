# Docker：
是一个开源的引擎，也是一个容器化平台，可以轻松的为任何应用创建一个轻量级的、可移植的、自给自足的容器。它将应用程序和其依赖项(代码、运行时、系统工具、系统库)以容器的形式打包在一起，以确保应用程序在任何环境（无论是开发环境、测试环境还是生产环境）中无缝运行，并避免了应用程序之间运行时的相互影响；

## 三大核心：
​(1)镜像（Image）一个只读的模板，镜像可以用来创建容器；
​(2)容器（Container）：类似于一个轻量级的沙箱，用来运行和隔离应用。容器是从镜像创建的应用运行实例，可以启动，开始，停止，删除等，可以理解为是一个简易版的linux环境以及运行在其中应用程序打包的盒子；
​(3)仓库（Repository）：是Docker集中存放镜像文件的场所，类似于代码仓库git，和Git和Github是很相似的。

## Docker与LXC（Linux Container）:
​(1)LXC通过强大的API和简单的工具，它可以让Linux用户轻松的创建和托管系统或者应用程序容器。Docker底层用的Linux的cgroup和namespace这两项技术来实现应用隔离，
​(2)Docker的对LXC的改进：
​    1.更好的移植性：可以实现应用无障碍平台移植；
​    2.容器更轻量，便携，高效。虚拟机虚拟计算机硬件，容器是针对操作系统的隔离；
    3.更丰富，便利的开发环境：镜像系统，仓库系统，多容器编排管控系统等

## 常用命令：
容器生命周期管理：run、start/stop/restart、kill、rm、pause/unpause、create、exec
容器操作：ps、inspect、top、attach、events、logs、wait、export、port
容器rootfs命令：commit、cp、diff
镜像仓库：login、pull、push、search
本地镜像管理：images、rmi、tag、build、bistory、save、import、info|version

## 用到的命令：
docker tag mysql:latest 192.168.14.30:433/csphere/mysql:latest
docker push 192.168.14.30:443/csphere/nginx:latest
docker rmi nginx，删除镜像

 

 

 