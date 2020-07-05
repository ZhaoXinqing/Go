## 我所理解的Sync Pool ##

看gin源码时发现了sync.Pool的使用
    
    // gin.go:L144
    func New() *Engine {
    	...
    
    	engine.pool.New = func() interface{} {
    		return engine.allocateContext()
    	}
    	return engine
    }
    
    // gin.go: L346
    func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    	c := engine.pool.Get().(*Context)
    	c.writermem.reset(w)
    	c.Request = req
    	c.reset()
    
    	engine.handleHTTPRequest(c)
    
    	engine.pool.Put(c)
    }
那个时候其实不太明白这个Pool是在干啥用, 大致觉得应该是内存池之类的. 后面想仔细看下sync.Pool具体怎么用, 我就去直接看了下Pool的源码, 然后直接懵逼了

因为基本看不懂其逻辑, 因为Pool的源码涉及到Golang的调度相关的知识. 要是不明白Golang是如何调度的, 基本是看不懂Pool的源码的(虽说它只有短短的259行代码, 文件路径: SDK/src/sync/pool.go). 这里推荐Go调度器系列系列文章

**后面我就去看看别人是怎么理解和使用Pool的**, 我就搜到了这么一篇文章 go语言的官方包sync.Pool的实现原理和适用场景里面有这么一个例子

    package main
    
    import(
    	"fmt"
    	"sync"
    )
    
    func main() {
    	p := &sync.Pool{
    		New: func() interface{} {
    			return 0
    		},
    	}
    
    	a := p.Get().(int)
    	p.Put(1)
    	b := p.Get().(int)
    	fmt.Println(a, b)
    }
以及Golang中国论坛上一个人的提问, 我不禁陷入了深深的疑问, 这个Golang的Pool到底是用来做什么的?

我疑问的地方:

sync.Pool的运用场景是什么, 哪些地方能用到Pool, 哪些地方不能用?
为啥CSDN上先Put后Get, 就能拿到1; 而Golang中国上那个提问, 为啥得到的顺序就不确定了?
我后面想了很久, 以及再回头去看gin, logrus的源码, 我才想明白: 这两个例子都是"坑货". 这两个例子都是sync.Pool使用的反面例子, 都是不正确的用法. 其实这些在源码的注释中, 都是由明确说明的, 只是当时没能理解. Callers should not assume any relation between values passed to Put and the values returned by Get. 简单来说: 就是Get和Put没有任何关系, 我们用Pool的时候, 要时刻记得这个.

因此, 我的疑问2就迎刃而解了, 其实就是这个一直困惑着我

关于疑问1就更简单了

    // Pool's purpose is to cache allocated but unused items for later reuse,
    // relieving pressure on the garbage collector. That is, it makes it easy to
    // build efficient, thread-safe free lists. However, it is not suitable for all
    // free lists.
简单来说, Pool就是为了减少GC压力的, 重复利用内存. 千万不能把他当成内存池使用

其实Pool的用法很简单, 就是先Get, 用完之后Put, 如gin的使用.

再比如logrus的用法

    func (logger *Logger) Println(args ...interface{}) {
    	entry := logger.newEntry()
    	entry.Println(args...)
    	logger.releaseEntry(entry)
    }
    
    func (logger *Logger) newEntry() *Entry {
    	entry, ok := logger.entryPool.Get().(*Entry)
    	if ok {
    		return entry
    	}
    	return NewEntry(logger)
    }
    
    func (logger *Logger) releaseEntry(entry *Entry) {
    	entry.Data = map[string]interface{}{}
    	logger.entryPool.Put(entry)
    }
所以别被别的池子带跑了, golang里的sync.Pool就是GC优化的, 用法很简单

在gocn有这么一个问题: 想在pool的基础上做一个限制池中对象数量的功能, 发现还是多次执行pool.New. 期望是只执行一次NewBuffer，也就是只打印一次alloc。 实际上每次执行，会打印多次alloc

    const MaxFrameSize = 5000
    
    func main() {
    	for i := 0; i < 10; i++ {
    		// 多个协程想从pool中拿到对象
    		go func() {
    			c := getBuf()
    			putBuf(c)
    			log.Println("put done")
    		}()
    	}
    
    	time.Sleep(3 * time.Second)
    }
    
    var bufPool = sync.Pool{
    	New: func() interface{} {
    		log.Println("alloc")
    		return bytes.NewBuffer(make([]byte, 0, MaxFrameSize))
    	},
    }
    
    var bufPoolChan = make(chan bool, 1)
    
    func getBuf() *bytes.Buffer {
    	bufPoolChan <- true
    	b := bufPool.Get().(*bytes.Buffer)
    	b.Reset()
    	return b
    }
    
    func putBuf(b *bytes.Buffer) {
    	bufPool.Put(b)
    	<-bufPoolChan
    }
大神解答:

- sync.pool的源代码里说了，pool里的对象随时都有可能被自动移除，并且没有任何通知。sync.pool的数量是不可控制的。

- pool调用new与线程调度有关，pool内部有一个localpool的数组，每个p对应其中一个localpool，在当前p执行goroutine的时候，优先从当前的localpool的private变量取，娶不到在从shared列表里面取，再取不到就尝试从别的p的localpool的shared里面偷一个。最后实在取不到就new一个。
 
- 由于你的bufpoolchan限制基本上10个goroutine就在两个p后面排队轮流执行，所以alloc就会出现两次，后面的基本就是从这两个localpool的private取出来的。

- 如果取消这个限制，10个goroutine很快就被分配到10个p上去了，对应就有10个localpool，10次每次取private都取不到，取shared列表也取不到，别的localpool也没得偷，就会new10次，alloc就会出现10次。

[from：](https://studygolang.com/articles/21023)https://studygolang.com/articles/21023