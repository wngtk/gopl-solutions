// 练习 4.7: 修改reverse函数用于原地反转UTF-8编码的[]byte。

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(string(reverse([]byte("你好，世界 hello world"))))
}

func reverse(s []byte) []byte {
	ret := make([]byte, len(s))
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRune(s[i:])
		if r == utf8.RuneError && size == 1 {
			break
		}
		copy(ret[len(s)-i-size:len(s)-i], s[i:i+size])
		i += size
	}
	return ret
}
