package main

import "testing"

func Test_list_add(t *testing.T) {
	l := list{}
	if l.len() != 0 {
		t.Error("new list length must be zero, got: ", l.len())
	}
}

func Test_list_add_one(t *testing.T) {
	l := list{}
	l.addLast(0)
	if l.len() != 1 {
		t.Error("after adding one element size must be 1, got: ", l.len())
	}
	if l.First() != 0 {
		t.Error("first in line, want:", 0, " got:", l.First())
	}
	if l.Last() != 0 {
		t.Error("last in line, want:", 0, " got:", l.Last())
	}
	l.queuing()
	if l.len() != 0 {
		t.Error("after deletion one element size must be 0, got: ", l.len())
	}
}
func Test_list_add_several(t *testing.T) {
	l := list{}
	l.addLast(0)
	l.addLast(1)
	l.addLast(2)
	if l.len() != 3 {
		t.Error("after adding 3 element size must be 3, got: ", l.len())
	}
	if l.First() != 0 {
		t.Error("first in line, want:", 0, " got:", l.First())
	}
	if l.Last() != 2 {
		t.Error("last in line, want:", 2, " got:", l.Last())
	}
	l.queuing()
	if l.len() != 2 {
		t.Error("after deletion one element size must be 2, got: ", l.len())
	}
}
