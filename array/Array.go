package array

type Array struct {
	length int
	data   map[int]string
}

func NewArray() *Array {
	return &Array{
		length: 0,
		data:   make(map[int]string),
	}
}

func (a *Array) Push(item string) {
	a.data[a.length] = item
	a.length++
}

func (a *Array) Get(index int) string {
	return a.data[index]
}

func (a *Array) Pop() string {
	item := a.data[a.length-1]
	delete(a.data, a.length-1)
	a.length--
	return item
}

func (a *Array) Delete(index int) string {
	item := a.data[index]
	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	delete(a.data, a.length-1)
	a.length--
	return item
}
