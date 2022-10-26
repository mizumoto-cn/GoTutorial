// in package container/heap

type Interface interface {
	sort.Interface      //嵌入字段sort.Interface
	Push(x interface{}) //a Push method to push elements into the heap
	Pop() interface{}   //a Pop elements that pops elements from the heap
}

//We see that sort.Interface is an embedded interface,
//so the above Interface has the three methods
//contained within the sort.Interface implicitly.
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less returns whether the element with index i should sort
	// before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}