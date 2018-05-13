package treesort

type tree struct {
	val         int
	left, right *tree
}

func Sort(values []int) []int {
	var t *tree
	for _, val := range values {
		add(t, val)
	}
	return appendValues(values[:0], t)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.val)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		return &tree{val: value}
	}
	if value < t.val {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
