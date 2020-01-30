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
	l.addLast("Babulya")
	if l.size != 1 {
		t.Error("after adding one element size must be 1, got: ", l.size)
	}
	if l.First() != "Babulya" {
		t.Error("first in queue is Babulya, got:", l.First())
	}
	if l.Last() != "Babulya" {
		t.Error("last in queue is Babulya, got:", l.Last())
	}
	increment := l.queuing()
	if increment != "Babulya" {
		t.Error("the first person in line is Babulya, got: ", increment)
	}
	if l.size != 0 {
		t.Error("after deletion one element size must be 0, got: ", l.size)
	}
}
func Test_list_add_several(t *testing.T) {
	l := queue{}
	l.addLast("Babulya")
	l.addLast(1)
	l.addLast(2)
	if l.size != 3 {
		t.Error("after adding 3 element size must be 3, got: ", l.size)
	}
	if l.First() != "Babulya" {
		t.Error("first in queue is Babulya, got:", l.First())
	}
	if l.Last() != 2 {
		t.Error("last in queue is 2, got:", l.Last())
	}
	increment := l.queuing()
	if increment != "Babulya" {
		t.Error("the first person in line is Babulya, got: ", increment)
	}
	if l.size != 2 {
		t.Error("after deletion one element size must be 2, got: ", l.size)
	}
}
