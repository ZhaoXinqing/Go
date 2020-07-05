// session 模块是用来存储客户端用户，session 模块目前只支持 cookie 方式的请求，如果客户端不支持
// cookie，那么就无法使用该模块。

import (
	"github.com/astaxie/beego/session"
)
// 初始化一个全局的变量用来存储 session 控制器：
var globalSessions *session.Manager

// 在入口函数中初始化数据
func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessioni
d", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600,
"secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLi
feTime": 3600, "providerConfig": ""}`)
	go globalSessions.GC()
}

// 业务逻辑处理函数
func login(w http.ResponseWriter, r *http.Request) {
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	username := sess.Get("username")
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	}else {
		sess.Set("username",r.Form["username"])
	}
}

// beego 的这个 session 模块设计的时候就是采用了
// interface，所以你可以根据接口实现任意的引擎，例如 memcache 的引擎
type SessionStore interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete{key interface{}} error
	SessionID() string
	SessionRelease()
	Flush() error
}

type Provider interface {
	SessionInit(maxLifetime int64, savePath string) error
	SessionRead(sid string) (SessionStore, error)
	SessionExist(sid string) bool
	SessionRegenerate(oldsid, sid string) (SessionStore, error)
	SessionDestroy(sid string) error
	SessionAll() int // GET ALL active session
	SessionGC()
}

// 注册自己写的引擎
func init() {
	Register("own",ownadaper)
}
