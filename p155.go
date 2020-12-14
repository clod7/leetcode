package main

import "math"

type MinStack struct {
	nums []int
	mins []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		nums: make([]int, 0),
		mins: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.nums = append(this.nums, x)
	top := this.mins[len(this.min)-1]
	this.mins = append(this.mins, min(top, x))
}

func (this *MinStack) Pop() {
	if len(this.nums) < 1 {
		return
	}

	this.nums = this.nums[:len(this.nums)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	if len(this.nums) < 1 {
		return 0
	}

	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.nums) < 1 {
		return 0
	}

	return this.mins[len(this.mins)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
