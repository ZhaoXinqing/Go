import (
	"github.com/astaxie/beego/cache"
)

// 初始化全局变量：
bm, err := cache.NewCache("memory", `{"interval":60}`)

bm.Put("astaxie", 1, 10)
bm.Get("astaxie")
bm.IsExist("astaxie")
bm.Delete("astaxie")

// 开发自己的引擎
// cache 模块采用了接口的方式实现，因此用户可以很方便的实现接口，然后注册就可以实现自己的 Cache 引擎：

type Cache interface {
	Get(key string) interface{}
	GetMulti(keys []string) []interface{}
	Put(key string, val interface{}, timeout int64) error
	Delete(key string) error
	Incr(key string) error
	Decr(key string) error
	IsExist(key string) bool
	ClearAll() error
	StartAndGC(config string) error
} 
// 用户开发完毕在最后写类似这样的：
func init() {
	Register("myowncache", NewOwnCache())
}