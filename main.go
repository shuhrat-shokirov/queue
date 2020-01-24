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

func (receiver *list) First() *listNode {
	return receiver.first
}

func (receiver *list) Last() *listNode {
	return receiver.last
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
	//l := list{}
	//l.addLast("", 3)
	//l.addLast("", 3)
	//l.addLast("", 2)
	//l.addLast("", 2)
	//l.addLast("", 1)
	//l.addLast("", 1)
	//l.addLast("", 0)
	//l.addLast("", 0)
	//l.addLast("", 4)
	//l.addLast("", 3)
	//l.addLast("", 2)
	//l.addLast("", 1)
	//l.addLast("", 0)
	//fmt.Println(l.first)
	//fmt.Println(l.last)
	//l.queuing()
	//fmt.Println(l.first)
	//l.queuing()
	//fmt.Println(l.first)
	//l.queuing()
	//fmt.Println(l.first)
	//l.queuing()
	//fmt.Println(l.first)
	//l.priorityQueue("", 0)
	//l.priorityQueue("", 0)
	//l.priorityQueue("", 1)
	///*current := l.first
	//for current.next != nil {
	//	fmt.Println(current)
	//	current = current.next
	//}
	//fmt.Println(current)
	//*/
}

/*
func (receiver *list) priorityQueue(elementPtr interface{}, priority int)  {
	if receiver.len() == 0 {
		receiver.first = &listNode{
			next:     nil,
			prev:     nil,
			value:    elementPtr,
			priority: priority,
		}
	}
	current := receiver.last
	for current.prev != nil{
		if current.priority < priority{
			prev := receiver.last
			for current.prev != prev{
				prev.prev = prev
				if prev == receiver.last{
					prev.next = nil}
				if prev != receiver.last{
					prev.next = prev.next.next
				}
				prev = prev.prev
			}
			return
		}
		prev := receiver.last
		for
		current = current.prev
	}
}
*/
