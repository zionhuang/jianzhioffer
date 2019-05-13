package leetcode


type (
	Queue struct {
		Head *node
		Tail *node
		Length int
	}
	node struct {
		Val interface{}
		Pre *node
		Aft *node
	}
)

