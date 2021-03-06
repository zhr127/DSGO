package bptree

import (
	"math/rand"
	"testing"
	"time"
)

func assert(t *testing.T, state bool) {
	if !state {
		t.Fail()
	}
}
func guardUT(t *testing.T) {
	if err := recover(); err != nil {
		t.Fail()
	}
}

func Test_Tree(t *testing.T) {
	defer guardUT(t)

	var tree Tree
	const size = 5000
	list := new([size]int)
	cnt := 0

	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		list[i] = rand.Int()
	}

	for i := 0; i < size; i++ {
		if tree.Insert(list[i]) {
			cnt++
		}
	}
	for i := 0; i < size; i++ {
		assert(t, tree.Search(list[i]))
		assert(t, !tree.Insert(list[i]))
	}

	utMark = -(int(^uint(0) >> 1)) - 1
	tree.Travel(checkNum)

	for i := 0; i < size; i++ {
		if tree.Remove(list[i]) {
			cnt--
		}
		assert(t, !tree.Search(list[i]))
	}
	assert(t, tree.IsEmpty() && cnt == 0)
	assert(t, !tree.Remove(0))
}

var utMark = 0

func checkNum(val int) {
	if val < utMark {
		panic(val)
	}
	utMark = val
}
