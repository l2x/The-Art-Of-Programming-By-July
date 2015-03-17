// 1.1 单词翻转
// @author leiyonglin <leiyonglin@gmail.com>
package main

import "fmt"

// 翻转单词
// I am a student.
// 先翻转整个句子
// .tneduts a ma I
// 再逐个单词翻转
// student. a am I
func reverseWord(s []byte) {
	l := len(s) - 1

	// 先翻转整个句子
	reverse(s, 0, l)
	fmt.Println(string(s))

	// start和end标记单词的开始和结束位置
	// 空格分割
	start := 0
	end := 0
	for start < l {
		// 开始是空格
		if s[start] == 32 {
			start++
			end = start + 1
			continue
		}
		// 结束是空格
		// 翻转单词
		if s[end] == 32 {
			reverse(s, start, end-1)
			start = end + 1
			end = start + 1
			continue
		}
		// 查找到末尾
		if end == l {
			reverse(s, start, end)
			break
		}

		//否则start不动，end++继续寻找单词的结束位置
		end++
	}
}

func reverse(s []byte, start, end int) {
	var tmp byte

	for start < end {
		tmp = s[start]
		s[start] = s[end]
		s[end] = tmp

		start++
		end--
	}
}

func main() {
	s := []byte("I am a student.")
	fmt.Println(string(s))

	reverseWord(s)
	fmt.Println(string(s))
}
