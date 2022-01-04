package _16_populating_next_right_pointers_in_each_node

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		var currentLevel []*Node
		for i := 0; i < len(queue); i++ {
			if i + 1 < len(queue) {
				queue[i].Next = queue[i + 1]
			}
			if queue[i].Left != nil {
				currentLevel = append(currentLevel, queue[i].Left)
			}
			if queue[i].Right != nil {
				currentLevel = append(currentLevel, queue[i].Right)
			}
		}
		queue = currentLevel
	}
	return root
}