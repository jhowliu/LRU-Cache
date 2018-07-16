package algorithm

const MAX = 5

type LRUCache struct {
	q []int
}

func (lru *LRUCache) refer(x int) {
	// x not present in cache
	if t := lru.search(x); t == -1 {
		// cache fulled
		if lru.size() == MAX {
			lru.deleteLast()
		}
	} else {
		lru.deleteAt(t)
	}

	lru.insertFront(x)
}

func (lru *LRUCache) insertFront(x int) {
	lru.q = append([]int{x}, lru.q...)
}

func (lru *LRUCache) deleteAt(index int) {
	if index != MAX-1 {
		lru.q = append(lru.q[:index], lru.q[index+1:]...)
	} else {
		lru.deleteLast()
	}
}

func (lru *LRUCache) deleteLast() {
	lru.q = lru.q[:len(lru.q)-1]
}

func (lru *LRUCache) size() int {
	return len(lru.q)
}

func (lru *LRUCache) search(target int) int {
	for i, v := range lru.q {
		if v == target {
			return i
		}
	}

	return -1
}
