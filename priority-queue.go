// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)

// An QueueItem is something we manage in a priority queue.
type QueueItem struct {
	Value    *QueueItemValue // The value of the item; arbitrary.
	Priority float64         // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*QueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push pushes item onto priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*QueueItem)
	item.Index = n
	*pq = append(*pq, item)
}

// Pop item from priority queue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Print prints items
func (pq *PriorityQueue) Print() {
	pqueue := *pq
	fmt.Println("* Printing priority queue of length", len(pqueue))
	for i := 0; i < len(pqueue); i++ {
		item := pqueue[i]
		fmt.Println(item.Value.Node.Hash, "prio, dista", item.Priority, item.Value.Distance)
	}
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *QueueItem, value *QueueItemValue, priority float64) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
// func test() {
// 	// Some items and their priorities.
// 	var coords = Coordinate{1.0, 1.0}
// 	var QueueItem3 = CreateNode(coords)

// 	items := map[int]Node{
// 		3: node3,
// 		2: node3,
// 		4: node3,
// 	}

// 	// Create a priority queue, put the items in it, and
// 	// establish the priority queue (heap) invariants.
// 	pq := make(PriorityQueue, len(items))
// 	i := 0
// 	for priority, value := range items {
// 		pq[i] = &Item{
// 			Value:    &value,
// 			Priority: priority,
// 			Index:    i,
// 		}
// 		i++
// 	}
// 	heap.Init(&pq)

// 	// Insert a new item and then modify its priority.
// 	item := &Item{
// 		Value:    &node3,
// 		Priority: 1,
// 	}
// 	heap.Push(&pq, item)
// 	pq.update(item, item.Value, 5)

// 	// Take the items out; they arrive in decreasing priority order.
// 	for pq.Len() > 0 {
// 		item := heap.Pop(&pq).(*Item)
// 		fmt.Printf("%.2d:%s ", item.Priority, item.Value)
// 	}
// }
