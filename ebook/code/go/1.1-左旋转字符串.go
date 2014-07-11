//1.1 旋转字符串
//@author leiyonglin <leiyonglin@gmail.com>
//http://play.golang.org/p/RFRscmzuMr
package main

import (
	"fmt"
)

//解法一: 暴力移位法
//左移一位
func leftShiftOne(s []byte, n int) {
	if n == 0 {
		return
	}

	t := s[0]
	for i := 1; i < n; i++ {
		s[i-1] = s[i]
	}
	s[n-1] = t
}

//左移m位
func leftRotateString(s []byte, n int, m int) {
	m %= n

	for m > 0 {
		leftShiftOne(s, n)
		m--
	}
}

//解法二:三步反转法
func reverse(s []byte, from int, to int) {
	for from < to {
		t := s[from]
		s[from] = s[to]
		s[to] = t
		from++
		to--
	}
}

func leftshift(s []byte, n int, m int) {
	m %= n
	reverse(s, 0, m-1)
	reverse(s, m, n-1)
	reverse(s, 0, n-1)
}

func main() {
	s := []byte("abcdef")
	n := len(s)
	m := 8

	leftRotateString(s, n, m)
	fmt.Println(string(s))

	leftshift(s, n, m)
	fmt.Println(string(s))
}
