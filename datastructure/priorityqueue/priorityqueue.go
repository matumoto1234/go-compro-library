package priorityqueue

type PriorityQueue[T any] struct {
	Heap []*T
	less func(T, T) bool
}

func New[T any](less func(T, T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		Heap: []*T{},
		less: less,
	}
}

func (pq *PriorityQueue[T]) Push(x T) {
	pq.Heap = append(pq.Heap, &x)
	i := len(pq.Heap) // 1-indexed

	for i > 1 {
		p := i / 2 // 1-indexed

		if pq.less(*pq.Heap[p-1], *pq.Heap[i-1]) {
			break
		}

		pq.Heap[p-1], pq.Heap[i-1] = pq.Heap[i-1], pq.Heap[p-1]
		i = p
	}
}

func (pq *PriorityQueue[T]) Pop() T {
	top := *pq.Heap[0]

	n := len(pq.Heap)

	pq.Heap[0] = pq.Heap[n-1]
	pq.Heap = pq.Heap[:n-1]
	n--

	i := 1 // 1-indexed
	for 2*i <= n {
		l := 2 * i
		r := 2*i + 1

		if r > n { // if only exists left child
			r = l
		}

		min := pq.Heap[i-1]
		minIndex := i

		if pq.less(*pq.Heap[l-1], *min) {
			min = pq.Heap[l-1]
			minIndex = l
		}

		if pq.less(*pq.Heap[r-1], *min) {
			min = pq.Heap[r-1]
			minIndex = r
		}

		next := minIndex

		if i == next {
			break
		}

		pq.Heap[i-1], pq.Heap[next-1] = pq.Heap[next-1], pq.Heap[i-1]
		i = next
	}

	return top
}

func (pq *PriorityQueue[T]) Len() int {
	return len(pq.Heap)
}

func (pq *PriorityQueue[T]) Top() T {
	return *pq.Heap[0]
}
