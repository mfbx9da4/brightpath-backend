package main

import (
	"container/heap"
	"fmt"
	"math"
	"sync"
)

// Node a single node that composes the tree
type Node Coordinate

// Graph Basic graph complete with concurrency safe lock
type Graph struct {
	edges    Edges
	numEdges int64
	lock     sync.RWMutex
}

// Edges maps nodes to set of other pointers to other nodes
type Edges map[Node]SetOfNodes

// SetOfNodes Set of nodes
type SetOfNodes = map[*Node]Exists

// Exists null value for sets
type Exists struct{}

var exists Exists

// Init initialize maps
func (graph *Graph) Init() {
	graph.lock.Lock()
	graph.edges = make(Edges)
	graph.numEdges = 0
	graph.lock.Unlock()
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", *n)
}

// AddEdge adds an edge to the graph
func (graph *Graph) AddEdge(n1 Node, n2 *Node) {
	graph.lock.Lock()
	if graph.edges[n1] == nil {
		graph.edges[n1] = make(SetOfNodes)
	}
	if graph.edges[*n2] == nil {
		graph.edges[*n2] = make(SetOfNodes)
	}
	graph.edges[n1][n2] = exists
	graph.numEdges++
	graph.lock.Unlock()
}

func (graph *Graph) String() string {
	s := ""
	for node, edges := range graph.edges {
		s += node.String() + " -> "
		for child := range edges {
			s += child.String() + " "
		}
		s += "\n"
	}
	return s
}

// Print graph
func (graph *Graph) Print() {
	graph.lock.RLock()
	fmt.Println(graph.String())
	graph.lock.RUnlock()
}

// QueueItemValue Best Path to node
type QueueItemValue struct {
	Node     *Node
	Path     []Node
	Distance float64
}

// NodeQueue sorted queue of paths
type NodeQueue struct {
	items []QueueItemValue
	lock  sync.RWMutex
}

// New creates a new NodeQueue
func (s *NodeQueue) New() *NodeQueue {
	s.lock.Lock()
	s.items = []QueueItemValue{}
	s.lock.Unlock()
	return s
}

// Enqueue adds an Node to the end of the queue
func (s *NodeQueue) Enqueue(t QueueItemValue) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// Dequeue removes an Node from the start of the queue
func (s *NodeQueue) Dequeue() *QueueItemValue {
	s.lock.Lock()
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	s.lock.Unlock()
	return &item
}

// Front returns the item next in the queue, without removing it
func (s *NodeQueue) Front() *QueueItemValue {
	s.lock.RLock()
	item := s.items[0]
	s.lock.RUnlock()
	return &item
}

// IsEmpty returns true if the queue is empty
func (s *NodeQueue) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items) == 0
}

// Size returns the number of Nodes in the queue
func (s *NodeQueue) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items)
}

// FindNode find node from graph
func (graph *Graph) FindNode(coords Coordinate) Node {
	var closest = Node{coords[0], coords[1]}
	if graph.edges[closest] == nil {
		var minDistance float64 = math.MaxFloat64
		// Probably there is a better algo for this, just doing the brute force sorry :(
		for node := range graph.edges {
			dx := coords[0] - node[0]
			dy := coords[1] - node[1]
			distance := math.Sqrt(dx*dx + dy*dy)
			if distance < minDistance {
				closest = node
				minDistance = distance
			}
		}
	}
	return closest
}

// Route Best route and distance of route
type Route struct {
	Path     []Node
	Distance float64
}

/*
FindPath Uses A* routing to find shortest path
(Keeps a min heap sorted by elapsed + remaing distance).
*/
func (graph *Graph) FindPath(src, dest Node) Route {
	graph.lock.RLock()

	// Init priority queue
	var pqueue = make(PriorityQueue, 1)
	var rootPath = []Node{src}
	var rootValue = QueueItemValue{&src, rootPath, 0}
	pqueue[0] = &QueueItem{
		Value:    &rootValue,
		Priority: 0,
		Index:    0,
	}
	heap.Init(&pqueue)

	// Keep track of visited
	visited := make(map[*Node]bool)
	for {
		if pqueue.Len() == 0 {
			break
		}
		// pqueue.Print()
		pqitem := heap.Pop(&pqueue).(*QueueItem)
		cur := pqitem.Value
		node := cur.Node
		visited[node] = true

		for child := range graph.edges[*node] {
			// TODO: Calculate in km
			// https://stackoverflow.com/a/1253545/1376627
			dx := (node[0] - child[0])
			dy := (node[1] - child[1])
			remaingDx := (dest[0] - child[0])
			remainingDy := (dest[1] - child[1])
			elapsed := math.Sqrt(dx*dx+dy*dy) + cur.Distance
			remaining := math.Sqrt(remaingDx*remaingDx + remainingDy*remainingDy)

			if *child == dest {
				path := append(cur.Path, *child)
				return Route{path, elapsed}
			}

			if !visited[child] {
				// TODO: Only add to path if different gradient
				path := make([]Node, len(cur.Path))
				copy(path, cur.Path)
				path = append(path, *child)
				queueItem := QueueItemValue{child, path, elapsed}
				newItem := QueueItem{
					Value:    &queueItem,
					Priority: elapsed + remaining,
				}
				heap.Push(&pqueue, &newItem)
				pqueue.update(&newItem, newItem.Value, newItem.Priority)
				visited[child] = true
			}
		}
	}
	graph.lock.RUnlock()
	// No path
	return Route{[]Node{}, -1}
}

// CalculatePath Finds closest nodes to start and end
func (graph *Graph) CalculatePath(startCoords Coordinate, endCoords Coordinate) Route {
	nodeStart := graph.FindNode(startCoords)
	nodeEnd := graph.FindNode(endCoords)
	pathFound := graph.FindPath(nodeStart, nodeEnd)
	return pathFound
}
