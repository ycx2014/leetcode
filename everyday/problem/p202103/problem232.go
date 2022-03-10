package p202103

import "errors"

type MyQueue struct {
	size int
	as   Stack
	bs   Stack
}

/** Initialize your data structure here. */
func Constructor1() MyQueue {
	return MyQueue{
		size: 0,
		as:   Stack{items: []int{}},
		bs:   Stack{items: []int{}},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	for i := 0; i < this.size; i++ {
		b, _ := this.bs.pop()
		this.as.push(b)
	}
	this.bs.push(x)
	for i := 0; i <this.size ; i++ {
		a,_:=this.as.pop()
		this.bs.push(a)
	}
	this.size++

}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	b,_:= this.bs.pop()
	this.size--
	return  b
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	b,_:=this.bs.peek()
	return  b
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.size == 0
}

type Stack struct {
	items []int
}

func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) peek() (int, error) {
	if len(s.items) != 0 {
		n := len(s.items)
		item := s.items[n-1]
		return item, nil
	}
	return 0, errors.New("stack is empty ")

}

func (s *Stack) pop() (int, error) {
	if len(s.items) != 0 {
		n := len(s.items)
		item := s.items[n-1]
		s.items = s.items[0 : n-1]
		return item, nil
	}
	return 0, errors.New("stack is empty ")
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

//请你仅使用两个栈实现先入先出队列。队列应当支持一般队列的支持的所有操作（push、pop、peek、empty）：
//
//实现 MyQueue 类：
//
//void push(int x) 将元素 x 推到队列的末尾
//int pop() 从队列的开头移除并返回元素
//int peek() 返回队列开头的元素
//boolean empty() 如果队列为空，返回 true ；否则，返回 false
// 
//
//说明：
//
//你只能使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
//你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
// 
//
//进阶：
//
//你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。
// 
//
//示例：
//
//输入：
//["MyQueue", "push", "push", "peek", "pop", "empty"]
//[[], [1], [2], [], [], []]
//输出：
//[null, null, null, 1, 1, false]
//
//解释：
//MyQueue myQueue = new MyQueue();
//myQueue.push(1); // queue is: [1]
//myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
//myQueue.peek(); // return 1
//myQueue.pop(); // return 1, queue is [2]
//myQueue.empty(); // return false
