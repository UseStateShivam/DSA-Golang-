package main

type MinStack struct {
    minSt []int
    st []int
}

func Constructor() MinStack {
    return MinStack{
        minSt: []int{},
        st: []int{},
    }
}

func (m *MinStack) Push(val int)  {
    m.st = append(m.st, val)
    if len(m.minSt) == 0 || val <= m.minSt[len(m.minSt) - 1] {
        m.minSt = append(m.minSt, val)
    }
}


func (m *MinStack) Pop()  {
    if m.st[len(m.st)-1] == m.minSt[len(m.minSt)-1] {
        m.minSt = m.minSt[:len(m.minSt)-1]
    }
    m.st = m.st[:len(m.st)-1]
}


func (m *MinStack) Top() int {
    return m.st[len(m.st)-1]
}


func (m *MinStack) GetMin() int {
    return m.minSt[len(m.minSt)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

 // 5ms
 // 11.11mb