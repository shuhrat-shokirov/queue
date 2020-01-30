package main

import "testing"

func Test_list_add(t *testing.T) {
	l := queue{}
	if l.size != 0 {
		t.Error("new queue length must be zero, got: ", l.size)
	}
}

func Test_list_add_one(t *testing.T) {
	l := queue{}
	l.addLast(0)
	if l.size != 1 {
		t.Error("after adding one element size must be 1, got: ", l.size)
	}
	if l.First() != 0 {
		t.Error("first in queue, want:", 0, " got:", l.First())
	}
	if l.Last() != 0 {
		t.Error("last in queue, want:", 0, " got:", l.Last())
	}
	l.queuing()
	if l.size != 0 {
		t.Error("after deletion one element size must be 0, got: ", l.size)
	}
}
func Test_list_add_several(t *testing.T) {
	l := queue{}
	l.addLast(0)
	l.addLast(1)
	l.addLast(2)
	if l.size != 3 {
		t.Error("after adding 3 element size must be 3, got: ", l.size)
	}
	if l.First() != 0 {
		t.Error("first in queue, want:", 0, " got:", l.First())
	}
	if l.Last() != 2 {
		t.Error("last in queue, want:", 2, " got:", l.Last())
	}
	l.queuing()
	if l.size != 2 {
		t.Error("after deletion one element size must be 2, got: ", l.size)
	}
}
