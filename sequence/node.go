package sequence


type Node struct {
	Id int16
}

func (n *Node) get() int16 {
	return n.Id & maxNode
}