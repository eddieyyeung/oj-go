package duiliedezuidazhilcof

import (
	"container/list"
)

// 剑指 Offer 59 - II. 队列的最大值
// 请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

// 若队列为空，pop_front 和 max_value 需要返回 -1

// 示例 1：

// 输入:
// ["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
// [[],[1],[2],[],[],[]]
// 输出: [null,null,null,2,1,2]
// 示例 2：

// 输入:
// ["MaxQueue","pop_front","max_value"]
// [[],[],[]]
// 输出: [null,-1,-1]

// 限制：

// 1 <= push_back,pop_front,max_value的总操作数 <= 10000
// 1 <= value <= 10^5

// 方法二：维护一个单调的双端队列
// 思路

// 本算法基于问题的一个重要性质：当一个元素进入队列的时候，它前面所有比它小的元素就不会再对答案产生影响。

// 举个例子，如果我们向队列中插入数字序列 1 1 1 1 2，那么在第一个数字 2 被插入后，数字 2 前面的所有数字 1 将不会对结果产生影响。因为按照队列的取出顺序，数字 2 只能在所有的数字 1 被取出之后才能被取出，因此如果数字 1 如果在队列中，那么数字 2 必然也在队列中，使得数字 1 对结果没有影响。

// 按照上面的思路，我们可以设计这样的方法：从队列尾部插入元素时，我们可以提前取出队列中所有比这个元素小的元素，使得队列中只保留对结果有影响的数字。这样的方法等价于要求维持队列单调递减，即要保证每个元素的前面都没有比它小的元素。

// 那么如何高效实现一个始终递减的队列呢？我们只需要在插入每一个元素 value 时，从队列尾部依次取出比当前元素 value 小的元素，直到遇到一个比当前元素大的元素 value' 即可。

// 上面的过程保证了只要在元素 value 被插入之前队列递减，那么在 value 被插入之后队列依然递减。
// 而队列的初始状态（空队列）符合单调递减的定义。
// 由数学归纳法可知队列将会始终保持单调递减。
// 上面的过程需要从队列尾部取出元素，因此需要使用双端队列来实现。另外我们也需要一个辅助队列来记录所有被插入的值，以确定 pop_front 函数的返回值。

// 保证了队列单调递减后，求最大值时只需要直接取双端队列中的第一项即可。

// 复杂度分析

// 时间复杂度：O(1)（插入，删除，求最大值） 删除操作于求最大值操作显然只需要 O(1) 的时间。 而插入操作虽然看起来有循环，做一个插入操作时最多可能会有 nnn 次出队操作。但要注意，由于每个数字只会出队一次，因此对于所有的 n 个数字的插入过程，对应的所有出队操作也不会大于 n 次。因此将出队的时间均摊到每个插入操作上，时间复杂度为 O(1)。
// 空间复杂度：O(n)，需要用队列存储所有插入的元素。

type MaxQueue struct {
	Queue *list.List
	Deque *list.List
}

func Constructor() MaxQueue {

	return MaxQueue{
		Queue: list.New(),
		Deque: list.New(),
	}
}

func (this *MaxQueue) Max_value() int {
	if this.Queue.Len() == 0 {
		return -1
	}
	return this.Deque.Front().Value.(int)
}

func (this *MaxQueue) Push_back(value int) {
	this.Queue.PushBack(value)
	for e := this.Deque.Front(); e != nil; e = e.Next() {
		dv := e.Value.(int)
		if dv < value {
			this.Deque.Remove(e)
		}
	}
	this.Deque.PushBack(value)
}

func (this *MaxQueue) Pop_front() int {
	if this.Queue.Len() == 0 {
		return -1
	}
	f := this.Queue.Front().Value.(int)
	df := this.Deque.Front().Value.(int)
	this.Queue.Remove(this.Queue.Front())
	if f == df {
		this.Deque.Remove(this.Deque.Front())
	}
	return f
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
