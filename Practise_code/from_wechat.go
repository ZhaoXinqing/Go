//切片扩容
func main () {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
// [0 0 0 0 0 1 2 3]
func main() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt,Println(s)
}
// [1 2 3 4]

// new() 和 make() 的区别？
/* 
	1、new(T) 和 make(T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。
	2、new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T 的值。换句话
说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。
	3、make(T,args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，是经过
	初始化之后的 T 的引用。make() 只适用于 slice、map 和 channel.
*/

const (
	_ = iota
	c1 int = (10 * iota)
	c2
	d = iota
)
func main() {
	fmt.Printf("%d - %d - %d", c1,c2,d)
}
// 10 - 20 - 3

//并发，引用，输出10
const N = 10

func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i 
			mu.Unlock()
		}()
	}
		
}


type Student struct{
     Age int
 }

 func main(){
    s := Student{30}
    modify1(s)
	fmt.Println(s) //{30}
	
    modify2(&s)
    fmt.Println(s) //{32}
 }

 func modify1(s Student){
    s.Age = 31
 }

 func modify2(s *Student){
    s.Age = 32
 }