type Node struct {
	Value int
	Next  *Node
}

var root = new(Node)

func addNode(t *Node)