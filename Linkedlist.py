
 
# lnked list in python

class _Node:
   __slots__='_element','_Next'
   def __init__(self,element,Next):
       self._element=element
       self._Next=Next
class Linkedlist:
   def __init__(self):
       self._head=None
       self._tail=None
       self._size=0
   def __len__(self):
       return self._size
   def isempty(self):
       return self._size==0 
   def addlast(self,e):
       newest =_Node(e,None)
 
       if self.isempty():
           self._head = newest
       else:
           self._tail._Next = newest
       self._tail = newest
       self._size += 1
   def display(self):
       p=self._head
       while p:
           print(p._element,end='-->')
           p =p._Next
 
       print() 
L=Linkedlist()         
L.addlast(5)
L.addlast(66)
L.addlast(2)
L.addlast(9)
L.display()
print('Size:',len(L))

 
 
 
