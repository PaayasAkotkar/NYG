package dataset

//  Queue first in first out
type Queue[T any] struct {
	val []T
	len int64
}

// Front return pushes the value front
func (q *Queue[T]) Front(d T) {
	_f := []T{d}
	q.val = append(_f, q.val...)
	q.len += 1
}

// Pop return the value at front
func (q *Queue[T]) Pop() T {
	var count T
	if !q.IsEmpty() {
		count = q.val[0]
		q.len -= 1
	}
	return count
}
func (q *Queue[T]) Length() int64 {
	return q.len
}

// At return the value at given pos
func (q *Queue[T]) At(pos int) T {
	if len(q.val) <= pos {
		panic("index out of range")
	}
	return q.val[pos]
}

// Dequeue return removes all the associated value
func (q *Queue[T]) Dequeue() {
	q.val = append(q.val[:0], q.val[:0]...)
}

// IsEmpty return true if the collection is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.val) == 0
}

// AfterEraseFrom return erases all the element after given limit
// example: input: 1,2,3,4
// output: limit[2]: result: 1,2
// another one: output: limit[1]:result: 1
func (q *Queue[T]) AfterEraseFrom(pos int) {
	if pos != len(q.val)-1 {
		q.val = append(q.val[0:pos], q.val[pos:pos]...)
	}
}

// EraseOnPos return erases the element at the given index position
// input: 1,2,3
// output: at pos[0]: 2,3
// at post[1]: 1,3
func (q *Queue[T]) EraseOnPos(pos int) {
	// input:1 2 3 4 5
	// output: limit[2]: 1 2 4 5
	if pos != len(q.val) {
		q.val = append(q.val[:pos], q.val[pos+1:]...)
	} else {
		q.val = append(q.val[:pos], q.val[pos:]...)
	}
}

// EraseAfter return erases all the elements after given limit while keeping the element at given limit
// example: input: 1,2,3,4
// output: limit[2]: 123
func (q *Queue[T]) EraseAfter(pos int) {
	q.val = append(q.val[0:pos], q.val[pos])

}

// Range return collection of stored object
// tip: you can also use it like array
func (q *Queue[T]) Range() []T {
	return q.val
}

// EraseLimit return erases the elements in the given limit
// example: input: 1,2,3,4
// ouput: given limit [0,3]: result: 4
func (q *Queue[T]) EraseLimit(from int, to int) {
	if from > to {
		panic("from, to belongs to [0,...size of str] where from<to ")
	}
	q.val = append(q.val[to:], q.val[:from]...)
}
