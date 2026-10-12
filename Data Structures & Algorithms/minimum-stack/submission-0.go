type MinStack struct {
	minElement	int
	top			int
	stack		[]int
}

func Constructor() MinStack {
	return MinStack{
		minElement	: math.MaxInt,
		top			: 0,
		stack		: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if val < this.minElement{
		this.minElement = val
	}
	this.stack = append(this.stack, val)
	this.top++
}

func (this *MinStack) Pop() {
	if this.Top() == this.minElement {
		this.minElement = math.MaxInt
		for i:=0; i< this.top-1; i++ {
			if this.stack[i] < this.minElement {
				this.minElement = this.stack[i]
			}
		}
	}
	if this.top > -1 {
		this.stack = this.stack[:len(this.stack)-1]
		this.top--
	}
}

func (this *MinStack) Top() int {
	if this.top > -1 {
		return this.stack[this.top-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	return this.minElement
}
