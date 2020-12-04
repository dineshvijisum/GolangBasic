package main

import "fmt"

var count int
//Node represents the components of a binary search tree
type Node struct {
	Key   int
	Right *Node
	Left  *Node
}

//insert will add a node to the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		//move right
      if n.Right == nil{
      n.Right = &Node{Key: k}       
	  }else{
		  n.Right.Insert(k)
	  }
	} else if n.Key > k {
		//move to left
		if n.Left == nil{
			n.Left = &Node{Key: k}       
			}else{
				n.Left.Insert(k)
			}
		}
}

//Search will take in a key value
//and Return true if there is a node with that value
func (n *Node)Search(k int)bool{

	//Count the tree value
	count++

	if n == nil{
		return false
	}
	if n.Key < k {
		//move right
		return n.Right.Search(k)
	}else if n.Key > k{
       return n.Left.Search(k)
	}
	return true
}
func main() {
	tree := &Node{Key: 100}
	
	fmt.Println(tree)

	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(90)
	tree.Insert(22)
	tree.Insert(37)
	tree.Insert(932)
	tree.Insert(77)
	fmt.Println(tree.Search(9))
	fmt.Println(count)
}
