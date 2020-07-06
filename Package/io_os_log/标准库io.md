##Go语言中的IO操作封装在以下4个包中：
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
**官方说明：**

Read 将 len(p) 个字节读取到 p 中。它返回读取的字节数 n（0 <= n <= len(p)） 以及任何遇到的错误。
即使 Read 返回的 n < len(p)，它也会在调用过程中占用 len(p) 个字节作为暂存空间。
若可读取的数据不到 len(p) 个字节，Read 会返回可用数据，而不是等待更多数据。
当 Read 在成功读取 n > 0 个字节后遇到一个错误或 EOF (end-of-file)，它会返回读取的字节数。
它可能会同时在本次的调用中返回一个non-nil错误,或在下一次的调用中返回这个错误（且 n 为 0）。
一般情况下, Reader会返回一个非0字节数n, 若 n = len(p) 个字节从输入源的结尾处由 Read 返回，Read可能返回 err == EOF 或者 err == nil。并且之后的 Read() 都应该返回 (n:0, err:EOF)。
调用者在考虑错误之前应当首先处理返回的数据。这样做可以正确地处理在读取一些字节后产生的 I/O 错误，同时允许EOF的出现。

**2、io.Writer**

    type Writer interface {
    	Write(p []byte) (n int, err error)
    }
**官方说明：**

Write 将 len(p) 个字节从 p 中写入到基本数据流中。它返回从 p 中被写入的字节数 n（0 <= n <= len(p)）以及任何遇到的引起写入提前停止的错误。
若 Write 返回的 n < len(p)，它就必须返回一个 非nil 的错误。

 

这些interface被标准库中的哪些struct实现？下面罗列几个：

* os.File同时实现了io.Reader和io.Writer
* strings.Reader实现了io.Reader
* bufio.Reader/Writer分别实现了io.Reader和io.Writer
* bytes.Buffer同时实现了io.Reader和io.Writer
* bytes.Reader实现了io.Reader

更多信息可以访问http://docs.studygolang.com./pkg/查询。


其它一些列接口：

    type ReaderAt interface {
    	ReadAt(p []byte, off int64) (n int, err error)
    }
    
    type WriterAt interface {
    	WriteAt(p []byte, off int64) (n int, err error)
    }
    
    type ReaderFrom interface {
    	ReadFrom(r Reader) (n int64, err error)
    }
    
    type WriterTo interface {
    	WriteTo(w Writer) (n int64, err error)
    }
    
    type Seeker interface {
    	Seek(offset int64, whence int) (ret int64, err error)
    }
    
    type Closer interface {
    	Close() error
    }
    
    type ByteReader interface {
    	ReadByte() (c byte, err error)
    }
    
    type ByteWriter interface {
    	WriteByte(c byte) error
    }

此外还有ByteScanner、RuneReader、RuneScanner、ReadCloser、ReadSeeker、ReadWriteCloser、ReadWriteSeeker、ReadWriter、WriteCloser、WriteSeeker。

详情阅读https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter01/01.1.html


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
　　