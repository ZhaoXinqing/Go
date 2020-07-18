// 循环队列实现思路：
// 1.循环队列须要几个參数来确定
//   front，tail，length，capacity
//   front指向队列的第一个元素，tail指向队列最后一个元素的下一个位置
//   length表示当前队列的长度，capacity标示队列最多容纳的元素
// 2.循环队列各个參数的含义
// （1）队列初始化时，front和tail值都为零
// （2）当队列不为空时，front指向队列的第一个元素，tail指向队列最后一个元素的下一个位置；
// （3）当队列为空时，front与tail的值相等，但不一定为零
// （4）当（tail+1）% capacity == front ||  （length+1）== capacity 表示队列为满，
//     因此循环队列默认浪费1个空间
// 3.循环队列算法实现
// （1）把值存在tail所在的位置；
// （2）每插入1个元素，length+1，tail=（tail+1）% capacity
// （3）每取出1个元素，length-1，front=（front+1）% capacity
// （4）扩容功能，当队列容量满，即length+1==capacity时，capacity扩大为原来的2倍
// （5）缩容功能，当队列长度小于容量的1/4，即length<=capacity/4时，capacity缩短为原来的一半

type loopQueue struct {
	queues []interface{}
	front int
	tail int
	length int
	capacity int
}

func NewLoopQueue{
	loop := &loopQueue{
		queues : make([]interface{},0, 2)
		front : 0
		tail: 0
		length: 0
		capacity: 2
	}

	for i := 0; i < 2; i++ {
		loop.queues = append(loop.queues,"")
	}
	return loop
}

func (q *loopQueue) Len() int {
	return q.length
}

func (q *loopQueue) Cap() int {
	return q.capacity
}

func (q *loopQueue)IsEmpty() bool {
	return q.length == 0
}

func (q *loopQueue) IsFull() bool {
	return (q.length + 1) == q.capacity
}

func (q *loopQueue) GetFront() (interface{}, error) {
	if q.len() == 0 {
		return nil,error.New(
			"failed to getFront, queues is empty.")
	}
	return q.queues[q.front],nil
}

func (q *loopQueue)Enqueue(elem interface{}) {
	if q.IsFull() {
		buffer := new(loopQueue)
		for i := 0; i < 2*q.capacity;i++ {
			buffer.queues = append(buffer.queues,"")
		}
		for i := 0;i<q.length;i++ {
			buffer.queues[i] = q.queues[q.front]
			q.front = (q.front + 1) % q.capacity
		}
		q.queues = buffer.queues
		q.front = 0
		q.tail = q.length
		q.capacity = 2 * q.capacity
	}

	q.queues[q.tail] = elem
	q.tail = (q.tail + 1) % q.capacity
	q.length ++
}

func (q *loop Queue) Dequeue() (interface{},error) {
	if q.IsEmpty() {
		return nil,errors.New(
			"failed to dequeue,queues is empty."
		)
	}
	if q.length <= q.capcity/4{
		buffer := new(loopQueue)

		for i := 0;i < q.capacity/2;i++{
			buffer.queues = append(buffer.queues,"")
		}
		for i := 0;i < q.length;i++{
			buffer.queues[i] =q.queues[q.front]
			q.front = (q.front + 1) % q.capacity
		}
		q.queues = buffer.queues
		q.front = 0
		q.tail = q.length
		q.capacity = q.capacity/2
	}
	queue := q.queues(q.front)
	q.front = (q.front + 1) % q.capacity
	q.length --
	return queue,nil
}

func main() {
	q := queue4.NewLoopQueue()

	for i :=0;i < 5; i++ {
		q.Enqueue(fmt.Sprintln(strconv.Itoa(i) + "--world"))
	}
	
    fmt.Printf("isEmpty:%v, isFull:%v, len=%v, cap=%v, getFront=%v\n",
        q.IsEmpty(), q.IsFull(), q.Len(), q.Cap(), fmt.Sprintln(q.GetFront()))
    fmt.Printf("isEmpty:%v, isFull:%v, len=%v, cap=%v, dequeue=%v\n",
        q.IsEmpty(), q.IsFull(), q.Len(), q.Cap(), fmt.Sprintln(q.Dequeue()))
    fmt.Printf("isEmpty:%v, isFull:%v, len=%v, cap=%v, dequeue=%v\n",
        q.IsEmpty(), q.IsFull(), q.Len(), q.Cap(), fmt.Sprintln(q.Dequeue()))
    fmt.Printf("isEmpty:%v, isFull:%v, len=%v, cap=%v, dequeue=%v\n",
        q.IsEmpty(), q.IsFull(), q.Len(), q.Cap(), fmt.Sprintln(q.Dequeue()))
    fmt.Printf("isEmpty:%v, isFull:%v, len=%v, cap=%v, dequeue=%v\n",
        q.IsEmpty(), q.IsFull(), q.Len(), q.Cap(), fmt.Sprintln(q.Dequeue()))
    fmt.Printf("isEmpty:%v, isFull:%v, len=%v, cap=%v, dequeue=%v\n",
        q.IsEmpty(), q.IsFull(), q.Len(), q.Cap(), fmt.Sprintln(q.Dequeue()))
    fmt.Printf("isEmpty:%v, isFull:%v, len=%v, cap=%v, dequeue=%v\n",
        q.IsEmpty(), q.IsFull(), q.Len(), q.Cap(), fmt.Sprintln(q.Dequeue()))
    fmt.Printf("isEmpty:%v, isFull:%v, len=%v, cap=%v, getFront=%v\n",
        q.IsEmpty(), q.IsFull(), q.Len(), q.Cap(), fmt.Sprintln(q.GetFront()))
}