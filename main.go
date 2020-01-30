package main

type queue struct {
	first *queueNode
	last  *queueNode
	size  int
}

type queueNode struct {
	next  *queueNode
	prev  *queueNode
	value interface{}
	//priority int
}

func (receiver *queue) First() interface{} {
	return receiver.first.value
}

func (receiver *queue) Last() interface{} {
	return receiver.last.value
}


func (receiver *queue) queuing() interface{} {
	if receiver.size == 0 {
		return -1
	}
	if receiver.size == 1 {
		increment := receiver.first.value
		receiver.first = nil
		receiver.last = nil
		receiver.size--
		return increment
	}
	increment := receiver.first.value
	receiver.first = receiver.first.next
	receiver.first.prev = nil
	receiver.size--
	return increment
}

func (receiver *queue) addLast(elementPtr interface{}) {
	if receiver.size == 0 {
		current := &queueNode{
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
	current := &queueNode{
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
