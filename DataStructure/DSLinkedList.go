package main
import "fmt"
type node struct{
	data int
	next *node
}

type linkedList struct{
	head *node
	length int
}

//2nd node going to 1st 
func (l *linkedList) prepend (n *node){
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

//list out the all the node values
func (l linkedList) printListData(){
	toPrint := l.head
	for l.length != 0{
		fmt.Printf("%d ",toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

//Delate the node from the list
func (l *linkedList) deleteWithValue(value int){
//
if l.length == 0{
	return
}
//delate head node in the list
if l.head.data == value{
	l.head = l.head.next
    l.length--
}

	toDelete := l.head
	for toDelete.next.data != value{
		//node is nil o/p does not shows the error message
		if toDelete.next.next == nil{
			return
		}
	toDelete = toDelete.next
}
toDelete.next = toDelete.next.next
l.length--
}


func main(){
  mylist := linkedList{}
   node1 := &node{data:54}
   node2 := &node{data:48}
   node3 := &node{data:77}
   node4 := &node{data:27}
   node5 := &node{data:76}
   node6 := &node{data:45}
   mylist.prepend(node1)
   mylist.prepend(node2)
   mylist.prepend(node3)
   mylist.prepend(node4)
   mylist.prepend(node5)
   mylist.prepend(node6)
   mylist.printListData()
   mylist.deleteWithValue(27)
   mylist.printListData()
  /* mylist.deleteWithValue(76)
   mylist.deleteWithValue(100)//before if case we had error,aftr using if case in delete func for loop err not come
   mylist.deleteWithValue(45)//1st node not posible to delete,aftr we use if case in delete func it can be done
   mylist.printListData()
   emptyList := linkedList{}//77 and 78 line ignored because 1st node is delete
   emptyList.deleteWithValue(10)/*
}