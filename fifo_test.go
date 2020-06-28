package deque

import (
	"fmt"
	"testing"
)

func TestDequeue(t *testing.T) {
	f := New(6)
	t.Logf("cap:%d, size:%d\n", f.Cap(), f.Size())
	f.PushBack(0,1,2,3,4,5,6)
	fmt.Println(f.PopFront().(int))
	t.Logf("cap:%d, size:%d\n", f.Cap(), f.Size())
	f.PushBack(7,8)
	fmt.Println(f.PopFront())
	f.PushFront(9)
	t.Logf("cap:%d, size:%d\n", f.Cap(), f.Size())
	fmt.Println(f.PopBack())
	t.Logf("cap:%d, size:%d\n", f.Cap(), f.Size())
}
