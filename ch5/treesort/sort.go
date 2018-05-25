package treesort

type treeNode struct {
	val         int
	left, right *treeNode
}

// Sort the values
func Sort(values []int) {
	var root *treeNode
	for _, val := range values {
		root = add(root, val)
	}
	appendValues(values[:0], root)
}

func add(node *treeNode, val int) *treeNode {
	if node == nil {
		node = &treeNode{val: val}
	} else {
		if val < node.val {
			node.left = add(node.left, val)
		} else {
			node.right = add(node.right, val)
		}
	}
	return node
}

func appendValues(values []int, root *treeNode) []int {
	if root != nil {
		values = appendValues(values, root.left)
		values = append(values, root.val)
		values = appendValues(values, root.right)
	}
	return values
}
