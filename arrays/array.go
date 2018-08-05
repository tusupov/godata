package arrays

type Array struct {
	items []int
}

func (a *Array) Append(n int) {
	a.items = append(a.items, n)
}

func (a Array) Get(i int) int {
	return a.items[i]
}

func (a Array) Len() int {
	return len(a.items)
}

func (a Array) Items() []int {
	tmp := make([]int, len(a.items))
	copy(tmp, a.items)
	return tmp
}

func (a *Array) RotateLeft(d int) {

	size := len(a.items)
	if size == 0 {
		return
	}

	gcd := func(x, y int) int {
		for {
			if y <= 0 {
				break
			}
			x, y = y, x % y
		}
		return x
	}

	d = d % size

	for i := 0; i < gcd(size, d); i++ {
		tmp, j := a.items[i], i
		for {
			k := j + d
			if k >= size {
				k -= size
			}
			if k == i {
				break
			}
			a.items[j] = a.items[k]
			j = k
		}
		a.items[j] = tmp
	}

}

func New() Array {
	return Array{
		items: make([]int, 0),
	}
}
