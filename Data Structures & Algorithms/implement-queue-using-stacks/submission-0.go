type MyQueue struct {
	stackOne []int
	stackTwo []int
}

func Constructor() MyQueue {
	return MyQueue{
		stackOne: []int{},
		stackTwo: []int{},
	}
}


func (this *MyQueue) Push(x int) {
	this.stackTwo = append(this.stackTwo, x)

	for _, val := range this.stackOne {
		this.stackTwo = append(this.stackTwo, val)
	}

	temp := this.stackTwo
	this.stackOne = temp
	this.stackTwo = []int{}
}

func (this *MyQueue) Pop() int {
	pos := len(this.stackOne)-1
	pop := this.stackOne[pos]
	this.stackOne = this.stackOne[:pos]
	return pop
}

func (this *MyQueue) Peek() int {
	return this.stackOne[len(this.stackOne)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stackOne) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
