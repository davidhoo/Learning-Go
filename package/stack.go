package stack

type Stack struct {
    i    int
    data [10]int
}

func (s *Stack) Push(i int) {
    s.data[s.i] = i
    s.i++
}

func (s *Stack) Pop() int {
    s.i--
    return s.data[s.i]
}
