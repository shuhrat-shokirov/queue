package main

type list struct {
	first *listNode
	last  *listNode
	size  int
}

type listNode struct {
	next  *listNode
	prev  *listNode
	value interface{}
	//priority int
}

func (receiver *list) First() interface{} {
	return receiver.first.value
}

func (receiver *list) Last() interface{} {
	return receiver.last.value
}

func (receiver *list) Len() int {
	return receiver.len()
}

func (receiver *list) queuing() {
	if receiver.len() == 0 {
		return
	}
	if receiver.len() == 1 {
		receiver.first = nil
		receiver.last = nil
		receiver.size--
		return
	}
	receiver.first = receiver.first.next
	receiver.first.prev = nil
	receiver.size--
}
func (receiver *list) len() int {
	return receiver.size
}

func (receiver *list) addLast(elementPtr interface{}) {
	if receiver.len() == 0 {
		current := &listNode{
			next:  nil,
			prev:  nil,
			value: elementPtr,
			//priority: priority,
		}
		receiver.first, receiver.last = current, current
		receiver.size++
		return
	}
	prev := receiver.last
	current := &listNode{
		next:  nil,
		prev:  prev,
		value: elementPtr,
		//priority: priority,
	}
	prev.next, receiver.last = current, current
	receiver.size++
}
func main() {
}
