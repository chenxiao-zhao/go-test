package stack

import "strconv"

const LIMIT = 10

type Stack struct {
	ix   int
	data [LIMIT]int
}

func (st *Stack) Push(n int) {
	if st.ix+1 > LIMIT {
		return
	}
	st.data[st.ix] = n
	st.ix++
}

func (st *Stack) Pop() int {
	st.ix--
	return st.data[st.ix]
}

func (st Stack) String() string {
	str := ""
	for ix := 0; ix < st.ix; ix++ {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(st.data[ix]) + "] "
	}
	return str
}
