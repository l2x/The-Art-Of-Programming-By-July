// 1.1 链表翻转
// @author leiyonglin <leiyonglin@gmail.com>
package main

import "fmt"

// 单向链表
type List struct {
	first *element
	last  *element
	size  int
}

type element struct {
	value interface{}
	next  *element
}

func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		newElement := &element{value: value}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

func (list *List) Print() {
	el := list.first
	for {
		fmt.Println(el)
		if el.next == nil {
			break
		}
		el = el.next
	}
}

// 链表翻转
// list = 1->2->3->4->5->6, m = 2
// 步骤
// 1.翻转 1<-2<-3<-4<-5<-6
// 2.改变 第1, m+1项指针
func (list *List) Reverse(m int) {
	m = m % list.size

	var pReversedHead *element
	var pNode *element
	var pPrev *element
	var pNext *element
	var p *element
	var pm *element

	pNode = list.first

	i := 0
	for pNode != nil {
		if i == 0 {
			p = pNode
		}
		if i == m {
			pm = pNode
		}
		i++

		pNext = pNode.next
		if pNext == nil {
			pReversedHead = pNode
		}
		pNode.next = pPrev
		pPrev = pNode
		pNode = pNext
	}

	p.next = pReversedHead
	list.first = pm.next
	pm.next = nil
}

func New() *List {
	return &List{}
}

func main() {
	list := New()
	list.Add(1, 2, 3, 4, 5, 6)
	list.Print()

	list.Reverse(4)
	list.Print()
}
