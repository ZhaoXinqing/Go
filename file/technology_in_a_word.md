**缓存**：提高系统的响应速度，特别是在高并发场景下的热点数据的访问速度的提升；

**golang**:提供了array和slice两种序列结构。其中array是值类型。slice则是复合类型。slice是基于array实现的。

**文件读取**: ioutil一次性全部将内容读取到内存中，小文件读取占优势，因为物理内存的局限，不适合读取较大文件；
读取大文件，bufio读取会更快，无论是处理大文件还是小文件，bufio始终是最为平滑和高效的。
    
**递归**（recursion）：是一个树结构，以自相似方法重复事物的过程。

**迭代**（iteration）：是一个环结构，重复反馈过程的活动，每一次迭代的结果会作为下一次迭代的初始值。

**子网掩码**：只有一个作用，就是将某个IP地址划分成网络地址和主机地址两部分

交换机和路由
交换机是一个扩大网络的器材，能为子网中提供更多的连接端口，以便连接更多的电脑。交换机与路由器的区别：

- **工作层次不同**
   交换机主要工作在数据链路层（第二层）
   路由器工作在网络层（第三层）。
- **转发依据不同**
   交换机转发所依据的对象是：MAC地址。（物理地址）
   路由转发所依据的对象是：IP地址。（网络地址）
- **主要功能不同**
   交换机主要用于组建局域网，连接同属于一个(广播域)子网的所有设备，负责子网内部通信（广播）。
   路由主要功能是将由交换机组好的局域网相互连接起来，或者将他们接入Internet。
   交换机能做的，路由都能做。
   交换机不能分割广播域（子网），路由可以。
   路由还可以提供防火墙的功能。
   路由配置比交换机复杂。

这里我们还是需要说明一点：

- 交换机虽然主要依靠MAC地址查找工作在数据链路层，但是他也可以有IP地址，这样就可以进行远程登录等操作了。
- **LAN**，全称Local Area Network，中文名叫做**局域网**；**WAN**，全称Wide Area Network，中文名叫做**广域网**。WAN是一种跨越大的、地域性的计算机网络的集合。通常跨越省、市，甚至一个国家。广域网包括大大小小不同的子网，子网可以是局域网，也可以是小型的广域网；**WLAN**，全称Wireless LAN, **无线局域网**，通俗点讲就是WiFi。
- 家用的路由器，一般包括了交换机和路由器，因此他有两个接口——WAN端口用于连接至Internet；LAN端口用于连接至局域网设备。