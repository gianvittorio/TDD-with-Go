package generics

type StackNode[T any] struct {
	value T
	prev, next *StackNode[T]
}

type Stack [T any] interface {
	IsEmpty() bool
	Push(value T)
	Pop() (T, error)
}

type ErrStack string

func (err ErrStack) Error() string {
	return string(err)
}

const (
	ErrStackIsEmpty = ErrStack("Stack is empty")
)

type StackImpl[T any] struct {
	head, tail *StackNode[T]
}

func (st *StackImpl[T]) IsEmpty() bool {
	return st.head == st.tail && st.head == nil
}

func (st *StackImpl[T]) Push(value T) {
	newNode := &StackNode[T]{value: value}
	if st.IsEmpty() {
		st.tail = newNode
		st.head = st.tail

		return
	}

	st.tail.next = newNode
	newNode.prev = st.tail
	st.tail = newNode
}

func (st *StackImpl[T]) Pop() (T, error) {
	var value T
	if st.IsEmpty() {
		return value, ErrStackIsEmpty
	}

	if st.head == st.tail {
		value = st.tail.value
		st.tail = nil
		st.head = nil

		return value, nil
	}

	value = st.tail.value
	prev := st.tail.prev
	prev.next = nil
	st.tail.prev = nil
	st.tail = prev

	return value, nil
}

func NewStack[T any]() Stack[T] {
	return new(StackImpl[T])
} 
