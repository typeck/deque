package deque

type Deque struct {
	buf 		[]interface{}
	head 		uint
	tail 		uint
	size 		uint
}

func isPowerOf2(n uint) bool {
	return n != 0 && ((n & (n - 1)) == 0)
}

func min(x, y uint) uint {
	if x <= y {
		return x
	}
	return y
}


func roundupPowOf2(x uint) uint{
	if x == 0{
		return 0
	}
	var pos uint
	for i := x; i != 0; i >>= 1 {
		pos ++
	}
	return 1 << pos
}

func New(size uint) *Deque{
	var deque = &Deque{}
	deque.size = size
	if ! isPowerOf2(size) {
		deque.size = roundupPowOf2(size)
	}
	deque.buf = make([]interface{}, deque.size)
	return deque
}

//tail - head = buf data len.
//size - tail + head = free buf.
//tail == head : empty
//tail - head == size : full

func (q *Deque) PushBack(args... interface{}) {
	le := uint(len(args))
	q.growIfN(le)

	l := min(le, q.size - (q.tail & (q.size - 1)))
	copy(q.buf[(q.tail & (q.size - 1)):], args[:l])
	copy(q.buf[:le - l], args[l:le])

	q.tail += le
}

func (q *Deque) PushFront(arg interface{}) {
	q.growIfN(1)
	q.head --
	q.buf[q.head  & (q.size - 1)] = arg
}

//func (q *Deque) PopFront(le uint) []interface{} {
//
//	le = min(le, q.tail - q.head)
//	var res = make([]interface{}, le)
//
//	l := min(le, q.size - (q.head & (q.size - 1)))
//	copy(res, q.buf[(q.head & (q.size - 1)):(q.head & (q.size - 1))+l])
//	copy(res[l:], q.buf[:le - l])
//	q.head += le
//	return res
//}

func (q *Deque) PopFront() interface{} {
	if q.Empty() {
		return nil
	}
	q.head ++
	return q.buf[(q.head - 1) & (q.size - 1)]
}

func (q *Deque) PopBack() interface{} {
	if q.Empty() {
		return nil
	}
	q.tail --
	return q.buf[q.tail & (q.size - 1)]
}

func (q *Deque) Cap() uint {
	return q.size
}

func (q *Deque)Size() uint {
	return q.tail - q.head
}

func (q *Deque)Empty() bool {
	if q.tail == q.head {
		return true
	}
	return false
}

func (q *Deque)growIfN(n uint) {
	if q.tail - q.head >= n {
		return
	}

	buf := make([]interface{}, 2 * q.size)
	lHead := q.head & (q.size - 1)
	lTail := q.tail & (q.size - 1)
	le := q.tail - q.head
	if  le > q.size{
		copy(buf[0:], q.buf[lHead:])
		copy(buf[q.size	- lHead:], q.buf[0:lTail])
	}
	copy(buf[0:], q.buf[lHead:lTail])
	q.buf = buf
	q.size = 2 * q.size
	q.head = 0
	q.tail = le
}

func (q *Deque)Clear() {
	q.head = 0
	q.tail = 0
}