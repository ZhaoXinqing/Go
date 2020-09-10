# Docker：

1、是一个开源的引擎，也是一个容器化平台，可以轻松的为任何应用创建一个轻量级的、可移植的、自给自足的容器。
    (1)应用程序 + 依赖项（代码、运行时、系统工具、系统库） = 容器的形式打包在一起，以确保应用程序在任何环境（无论是开发环境、测试环境还是生产环境）中无缝运行；
    (2)加快了工作流程：通过将应用程序和其运行时的环境一起打包隔离，避免了应用程序之间运行时的相互影响，也解决了依赖环境冲突的问题，同时对开发人员来说，在不同的项目里可以有更大的创新和更多的选择，例如选择自己最适合的开发语言，或者性能最优的策略；

## 三大核心：

​    (1)镜像（Image）：类似于虚拟机镜像，可以理解为一个只读的模板，镜像可以用来创建容器；
​    (2)容器（Container）：类似于一个轻量级的沙箱，用来运行和隔离应用。容器是从镜像创建的应用运行实例，可以启动，开始，停止，删除等，可以理解为一个容器是一个简易版的linux环境以及运行在其中应用程序打包的盒子；
​    (3)仓库（Repository）：类似于代码仓库git，是Docker集中存放镜像文件的场所，和Git和Github是很相似的。
3、适合微服务：
​    (1)适合创建轻量级的隔离，应用程序和底层架构解耦（适应不同的服务环境）
​    (2)DockerSwarm和Kubernetes容器编排解决方案，让多个容器容易创建一套新的服务；

## Docker与LXC（Linux Container）:

​    (1)LXC通过强大的API和简单的工具，它可以让Linux用户轻松的创建和托管系统或者应用程序容器。
​    (2)Docker底层用的Linux的cgroup和namespace这两项技术来实现应用隔离，lxc是早期版本docker的一个基础组件，docker主要用到了它对Cgroup和Namespace两个内核特性的控制。新的Docker版本已经移除了对LXC的support。
​    (3)Docker的对LXC的改进：
​        ①更好的移植性：通过抽象容器配置，可以实现应用无障碍平台移植；
​        ②更丰富，便利的开发环境：镜像系统，仓库系统，多容器编排管控系统等；
​        ③虚拟机虚拟计算机硬件，而容器是针对操作系统的隔离，所以容器更轻量，便携，高效；
5、Docker友好的，基于CLI的工作流使构建、共享和运行容器化应用程序可供所有技能级别的开发人员访问；





## 命令：

\1. Docker pull： docker pull java，从Docker Hub下载java最新版镜像。

\2. Docker info：显示 Docker 系统信息，包括镜像和容器数

\3. Docker push：docker push myapache:v1，上传本地镜像myapache:v1到镜像仓库中

\4. Docker version：显示 Docker 版本信息。

\5. Docker diff：docker diff mymysql，查看容器mymysql的文件结构更改

\6. Docker save：

(1) docker save -o my_ubuntu_v3.tar w3cschool/ubuntu:v3，将镜像w3cschool/ubuntu:v3 生成my_ubuntu_v3.tar文档

\7. Docker login/logout：

(1) docker login -u 用户名 -p 密码，登陆到Docker Hub

(2) docker logout，登出Docker Hub

\8. Docker import：docker import  my_ubuntu_v3.tar w3cschool/ubuntu:v4，从镜像归档文件my_ubuntu_v3.tar 创建镜像，命名为w3cschool/ubuntu:v4

\9. Docker commit：docker commit -a "w3cschool.cn" -m "my apache" a404c6c174a2，将容器a404c6c174a2 保存为新的镜像,并添加提交人信息和说明信息；

\10. Docker build：

(1) docker build -t w3cschool/ubuntu:v1 . 使用当前目录的Dockerfile创建镜像。

(2) docker build github.com/creack/docker-firefox，使用URL github.com/creack/docker-firefox 的 Dockerfile 创建镜像。

\11. docker tag : 标记本地镜像，将其归入某一仓库。



## Docker 命令大全

 

容器生命周期管理

[run](https://www.w3cschool.cn/docker/docker-run-command.html)

[start/stop/restart](https://www.w3cschool.cn/docker/docker-start-stop-restart-command.html)

[kill](https://www.w3cschool.cn/docker/docker-kill-command.html)

[rm](https://www.w3cschool.cn/docker/docker-rm-command.html)

[pause/unpause](https://www.w3cschool.cn/docker/docker-pause-unpause-command.html)

[create](https://www.w3cschool.cn/docker/docker-create-command.html)

[exec](https://www.w3cschool.cn/docker/docker-exec-command.html)

容器操作

[ps](https://www.w3cschool.cn/docker/docker-ps-command.html)

[inspect](https://www.w3cschool.cn/docker/docker-inspect-command.html)

[top](https://www.w3cschool.cn/docker/docker-top-command.html)

[attach](https://www.w3cschool.cn/docker/docker-attach-command.html)

[events](https://www.w3cschool.cn/docker/docker-events-command.html)

[logs](https://www.w3cschool.cn/docker/docker-logs-command.html)

[wait](https://www.w3cschool.cn/docker/docker-wait-command.html)

[export](https://www.w3cschool.cn/docker/docker-export-command.html)

[port](https://www.w3cschool.cn/docker/docker-port-command.html)

容器rootfs命令

[commit](https://www.w3cschool.cn/docker/docker-commit-command.html)

[cp](https://www.w3cschool.cn/docker/docker-cp-command.html)

[diff](https://www.w3cschool.cn/docker/docker-diff-command.html)

镜像仓库

[login](https://www.w3cschool.cn/docker/docker-login-command.html)

[pull](https://www.w3cschool.cn/docker/docker-pull-command.html)

[push](https://www.w3cschool.cn/docker/docker-push-command.html)

[search](https://www.w3cschool.cn/docker/docker-search-command.html)

本地镜像管理

[images](https://www.w3cschool.cn/docker/docker-images-command.html)

[rmi](https://www.w3cschool.cn/docker/docker-rmi-command.html)

[tag](https://www.w3cschool.cn/docker/docker-tag-command.html)

[build](https://www.w3cschool.cn/docker/docker-build-command.html)

[history](https://www.w3cschool.cn/docker/docker-history-command.html)

[save](https://www.w3cschool.cn/docker/docker-save-command.html)

[import](https://www.w3cschool.cn/docker/docker-import-command.html)

info|version

[info](https://www.w3cschool.cn/docker/docker-info-command.html)

[version](https://www.w3cschool.cn/docker/docker-version-command.html)



## 用到的命令：

docker tag mysql:latest 192.168.14.30:433/csphere/mysql:latest

docker push 192.168.14.30:443/csphere/nginx:latest

docker rmi nginx，删除镜像

 

 

 