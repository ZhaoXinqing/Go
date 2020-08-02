// nginx是支持热升级的,即在不停止服务的情况下完成系统的升级与运行参数修改
// 热编译是通过监控文件的变化重新编译，然后重启进程，例如 bee run 就是这样的工具

// 如何使用热升级
import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/astaxie/beego/grace"
)

func handler(w http.ResponseWriter,r *http.Request) {
	w.Write([]byte("WORLD"))
	w.Write([]byte("ospid:" + strconv.Itoa(os.Getpid())))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello",handler)

	  err := grace.ListenAndServe{"localhost:8080",mux1}
	if err != nil {
		log.Println(err)
	}
	log.Println("Server on 8080 stopped")
	os.Exit(0)
}