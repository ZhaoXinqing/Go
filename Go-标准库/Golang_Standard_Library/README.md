## Go语言便准库——go语言中文网 ##
http://books.studygolang.com/The-Golang-Standard-Library-by-Example

Bufio	带缓冲的IO操作
Database	数据库驱动和接口
Encoding	常见算法如JSON、XML、Base64
io	实现I/O原始访问接口及访问封装
Math	数学库
Net	网络库，支持Socket、HTTP、邮件、RPC、SMTP等
Os	操作系统平台不依赖平台操作封装
Path	兼容各操作系统的路径操作使用函数
Plugin	Go 1.7 加入的插件系统，支持将代码编译成插件，按需加载
Reflect	语言发射支持，可以动态获得代码中的类型信息，获取和修改变量的值
Regexp	正则表达式封装
Runtime	运行时接口
Sort	排序接口
String	字符串转换、解析及使用函数
Time	时间接口，time包提供了时间的显示和测量用的函数。日历的计算采用的是公历。
Flag	命令行解析
Sync	基本的同步基元，如互斥锁。除了Once和WaitGroup类型，大部分都是适用于低水平程序线程，高水平的同步使用channel通信更好一些。