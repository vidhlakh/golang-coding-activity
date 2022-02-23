package reversewords

type Stack []string

func (st *Stack) Push(word string) {
	*st = append(*st, word)

}

func (st *Stack) Pop() (popped string) {
	if !st.IsEmpty() {
		index := len(*st) - 1
		popped = (*st)[index]
		*st = (*st)[:index]
	}
	return popped
}
func (st *Stack) IsEmpty() bool {
	return len(*st) == 0
}
