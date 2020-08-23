// 基础的算法，olang提供了chan类型，可以很方便的实现
// golang的官方文档，只使用chan就可以实现生产者和消费者之间的数据和状态同步

// 解chan操作接口和语意，其实就是send、recv和close，
// 类似于文件操作而没有open，文件read通过EOF来返回文件尾，
// 而chan则可以通过额外的返回值(ok)来获取close的信息。

// 生产者和消费者是并行工作的，而且在工作完成后通过chan同步状态。

package main
 
import "fmt"
 
func Producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}
 
func Consumer(id int, ch chan int, done chan bool) {
	for {
		value, ok := <-ch
		if ok {
			fmt.Printf("id: %d, recv: %d\n", id, value)
		} else {
			fmt.Printf("id: %d, closed\n", id)
			break
		}
	}
	done <- true
}
 
func main() {
	ch := make(chan int, 3)
 
	coNum := 2
	done := make(chan bool, coNum)
	for i := 1; i <= coNum; i++ {
		go Consumer(i, ch, done)
	}
 
	go Producer(ch)
 
	for i := 1; i <= coNum; i++ {
		<-done
	}
}



$ ./chan2 
id: 2, recv: 1
id: 1, recv: 2
id: 1, recv: 4
id: 1, recv: 5
id: 1, recv: 6
id: 2, recv: 3
id: 2, recv: 8
id: 2, recv: 9
id: 2, recv: 10
id: 2, closed
id: 1, recv: 7
id: 1, closed