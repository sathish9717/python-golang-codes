//circularlinkedlist in golang
package main

import "fmt"


type node struct{
     data  int
	 next *node

}
type circular struct{
	head *node
	tail *node
	length int
}
func (L *circular) Length() int{
	return L.length
}
//  func (L *circular) isempty()int{
//  	return L.length == 0
// }
func (L *circular) addlast(data int){
	newest:=&node{data:data}
	if L.length==0{
		newest.next =newest
		L.head=newest
	}else{
		newest.next=L.tail
		L.tail.next=newest
		}
	L.tail=newest
	L.length++	
	}
func (L *circular) display(){
	temp:=L.head
	
	for i:=0; i<L.length ;i++{
        fmt.Printf("%+v ->",temp.data)
		temp:=temp.next
	}
	fmt.Print()
}				
func(L *circular) GetHead(){
	fmt.Println("Head: ",L.head)
}
func(L *circular) GetTail(){
	fmt.Println("Head: ",L.tail)
}
		
	
	


