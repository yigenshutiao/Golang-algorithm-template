package _341_flatten_nested_list_iterator

type NestedInteger struct {
}

func (this NestedInteger) GetInteger() int {
	// 题里给了
	return 0
}

func (this NestedInteger) GetList() []*NestedInteger {
	// xx
	return nil
}

func (this NestedInteger) IsInteger() bool {
	// 题里给了
	return false
}

type NestedIterator struct {
	Vals []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var vals []int
	var dfs func(nestedList []*NestedInteger)
	dfs = func(nestedList []*NestedInteger) {
		for _, nest := range nestedList {
			if nest.IsInteger() {
				vals = append(vals, nest.GetInteger())
			} else {
				dfs(nest.GetList())
			}
		}
	}
	dfs(nestedList)
	return &NestedIterator{Vals: vals}
}

func (this *NestedIterator) Next() int {
	node := this.Vals[0]
	this.Vals = this.Vals[1:]
	return node
}

func (this *NestedIterator) HasNext() bool {
	return len(this.Vals) > 0
}
