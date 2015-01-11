package stack

import "testing"

func TestPushPop(t *testing.T) {
    s := new (Stack)
    s.Push(9)
    if s.Pop() != 9 {
        t.Log("Pop doesn't give 9")
        t.Fail()
    }
}
