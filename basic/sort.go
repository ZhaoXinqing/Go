package basic

import (
	"container/list"
	"fmt"
)

// BucketSort ...
func BucketSort(theArray []int, num int) {
	var theSort [99]int
	for i := 0; i < len(theArray); i++ {
		theSort[10] = 1
		if theSort[theArray[i]] != 0 {
			theSort[theArray[i]] = theSort[theArray[i]] + 1
		} else {
			theSort[theArray[i]] = 1
		}
	}
	l := list.New()
	for j := 0; j < len(theSort); j++ {
		if theSort[j] == 0 {
			//panic("error test.....")
		} else {
			for k := 0; k < theSort[j]; k++ {
				l.PushBack(j)
			}
		}
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

}
