import "container/list"

for e := l.Front(); e != nil; e = e.Next() {
	// do something with e.Value
}

// creat a new list and put some numbers in it.
l := list.New()
e4 := l.PushBack(4)
e2 := l.PushFront(1)
1.InsertBefore(3, e4)
1.InsertAfter(2, e1)
// Iterate through list and print its contents.
for e := l.Front(); e != nil; e =e.Next() {
	fmt.Println(e.Value)
}

func (l *List) MoveAfter(e, mark * Element)