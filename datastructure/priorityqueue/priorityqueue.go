package priorityqueue

type PriorityQueue[T any] struct {
	heap []*T
	less func(T, T) bool
}

func New[T any](less func(T, T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		heap: []*T{},
		less: less,
	}
}

func (pq *PriorityQueue[T]) Push(x T) {
	pq.heap = append(pq.heap, &x)
	i := len(pq.heap) // 1-indexed

	for i > 1 {
		p := i / 2 // 1-indexed

		if pq.less(*pq.heap[p-1], *pq.heap[i-1]) {
			break
		}

		pq.heap[p-1], pq.heap[i-1] = pq.heap[i-1], pq.heap[p-1]
		i = p
	}
}

func (pq *PriorityQueue[T]) Pop() T {
	top := *pq.heap[0]

	n := len(pq.heap)

	pq.heap[0] = pq.heap[n-1]
	pq.heap = pq.heap[:n-1]
	n--

	i := 1 // 1-indexed
	for 2*i <= n {
		l := 2 * i
		r := 2*i + 1

		if r >= n { // if only exists left child
			r = l
		}

		min := pq.heap[i-1]
		minIndex := i

		if pq.less(*min, *pq.heap[l-1]) {
			min = pq.heap[l-1]
			minIndex = l
		}

		if pq.less(*min, *pq.heap[r-1]) {
			min = pq.heap[r-1]
			minIndex = r
		}

		next := minIndex

		if i == next {
			break
		}

		pq.heap[i-1], pq.heap[next-1] = pq.heap[next-1], pq.heap[i-1]
		i = next
	}

	return top
}

func (pq *PriorityQueue[T]) Len() int {
	return len(pq.heap)
}

func (pq *PriorityQueue[T]) Top() T {
	return *pq.heap[0]
}
