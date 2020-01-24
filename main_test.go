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
	gotFirst := l.first
	gotLast := l.last
	if l.First() != gotFirst {
		t.Error("first in line, want:", l.First(), " got:", gotFirst)
	}
	if l.Last() != gotLast {
		t.Error("last in line, want:", l.Last(), " got:", gotLast)
	}
	l.queuing()
	if l.len() != 0 {
		t.Error("after deletion one element size must be 0, got: ", l.len())
	}
}
func Test_list_with_one_item(t *testing.T) {
	l := list{}
	l.addLast(0)
	l.addLast(1)
	l.addLast(2)
	if l.len() != 3 {
		t.Error("after adding 3 element size must be 3, got: ", l.len())
	}
	gotFirst := l.first
	gotLast := l.last
	if l.First() != gotFirst {
		t.Error("first in line, want:", l.First(), " got:", gotFirst)
	}
	if l.Last() != gotLast {
		t.Error("last in line, want:", l.Last(), " got:", gotLast)
	}
	l.queuing()
	if l.len() != 2 {
		t.Error("after deletion one element size must be 2, got: ", l.len())
	}
}