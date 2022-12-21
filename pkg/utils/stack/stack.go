package stack

func Stack[T any]() IStack[T] {
	return &stack[T]{}
}

type (
	IStack[T any] interface {
		Clear()
		IsEmpty() bool
		Pop() T
		Push(v any)
	}

	stack[T any] struct {
		root *item
	}
	item struct {
		value any
		prev  *item
	}
)

func (this *stack[T]) Clear() {
	this.root = nil
}

func (this *stack[T]) Push(v any) {
	this.root = &item{
		value: v,
		prev:  this.root,
	}
}

func (this *stack[T]) Pop() T {
	if this.root == nil {
		panic("stack is empty")
	}
	v := this.root.value.(T)
	this.root = this.root.prev
	return v
}

func (this *stack[T]) IsEmpty() bool {
	return this.root == nil
}
