package solution

type Item struct {
	Index int
	Value int
}

type Tree struct {
	Left  *Tree
	item  Item
	Right *Tree
}

func oddEvenJumps(A []int) int {
	odds := make([]bool, len(A))
	evens := make([]bool, len(A))
	odds[len(A)-1] = true
	evens[len(A)-1] = true
	t := Tree{
		nil,
		Item{
			len(A) - 1,
			A[len(A)-1],
		},
		nil,
	}
	count := 1
	for i := len(A) - 2; i >= 0; i-- {
		p := &t
		node := Tree{
			nil,
			Item{
				i,
				A[i],
			},
			nil,
		}
		floor := Item{
			-1,
			-1,
		}
		ceiling := Item{
			-1,
			-1,
		}
		for {
			if node.item.Value < p.item.Value {
				floor = p.item
				if p.Left == nil {
					p.Left = &node
					break
				}
				p = p.Left
			} else if node.item.Value > p.item.Value {
				ceiling = p.item
				if p.Right == nil {
					p.Right = &node
					break
				}
				p = p.Right
			} else {
				floor = p.item
				ceiling = p.item
				p.item = node.item
				break
			}
		}
		if floor.Index != -1 {
			if odds[floor.Index] == true {
				evens[i] = true
				count++
			}
		}
		if ceiling.Index != -1 {
			if evens[ceiling.Index] == true {
				odds[i] = true
			}
		}
	}
	return count
}

func Run(A []int) int {
	return oddEvenJumps(A)
}
