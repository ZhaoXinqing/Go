## Go语言中的IO操作封装在以下4个包中：##

1. io: 为IO原语（I/O primitives）提供基本的接口。
1. io/ioutil: 封装一些实用的I/O函数。
1. fmt: 实现格式化I/O。
1. bufio: 实现带缓冲I/O。

## 1. io——基本的IO接口 ##

该包主要内容是定义了很多io相关的interface，作为标准库中其它函数的参数出现，其中最基础重要的两个接口是io.Reader和io.Writer。

**1、io.Reader**

    type Reader interface {
    	Read(p []byte) (n int, err error)
    }

**2、io.Writer**

    type Writer interface {
    	Write(p []byte) (n int, err error)
    }

**官方说明：**

这些interface被标准库中的哪些struct实现？下面罗列几个：

* os.File同时实现了io.Reader和io.Writer
* strings.Reader实现了io.Reader
* bufio.Reader/Writer分别实现了io.Reader和io.Writer
* bytes.Buffer同时实现了io.Reader和io.Writer
* bytes.Reader实现了io.Reader


## 2. ioutil——方便的IO操作函数集 ##

**1) NopCloser函数**
将io.Reader的实例转换为io.ReadCloser实例。

**2) ReadAll函数**
func ReadAll(r io.Reader) ([]byte, error)
一次性读取io.Reader中的数据

**3) ReadDir函数**
读取目录并返回排好序的文件和子目录名。

**4) ReadFile和WriteFile函数**
func ReadFile(filename string) ([]byte, error)
ReadFile 从filename指定的文件中读取数据并返回文件的内容。成功的调用返回的err为nil而非EOF。因为本函数定义为读取整个文件，它不会将读取返回的EOF视为应报告的错误。

func WriteFile(filename string, data []byte, perm os.FileMode) error
WriteFile 将data写入filename文件中，当文件不存在时会根据perm指定的权限进行创建一个,文件存在时会先清空文件内容。对于perm参数，我们一般可以指定为：0666，具体含义os包中讲解。

**5) TempDir和TempFile函数**

创建临时目录。

 

## 3. fmt —— 格式化IO ##


## 4. bufio——缓存IO ##

**1) bufio.Reader**

包装了一个io.Reader对象，提供缓存功能，同时实现了io.Reader接口。

	type Reader struct {
	    buf          []byte        // 缓存
	    rd           io.Reader    // 底层的io.Reader
	    // r:从buf中读走的字节（偏移）；w:buf中填充内容的偏移；
	    // w - r 是buf中可被读的长度（缓存数据的大小），也是Buffered()方法的返回值
	    r, w         int
	    err          error        // 读过程中遇到的错误
	    lastByte     int        // 最后一次读到的字节（ReadByte/UnreadByte)
	    lastRuneSize int        // 最后一次读到的Rune的大小(ReadRune/UnreadRune)
	}

**2) bufio.Writer**

包装了一个io.Writer对象，提供缓存功能，同时实现了io.Writer接口。

	type Writer struct {
	    err error        // 写过程中遇到的错误
	    buf []byte        // 缓存
	    n   int            // 当前缓存中的字节数
	    wr  io.Writer    // 底层的 io.Writer 对象
	}
　　
