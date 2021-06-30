// 统计输入中每个Unicode码点出现的次数
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // unicode字符数
	invalid := 0                    // 不合法的utf8字符
	var utflen [utf8.UTFMax + 1]int // utf8编码长度

	letter, number, other := 0, 0, 0

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() // 执行utf8解码，并返回3个值。解码的rune字符的值，字符UTF-8编码后的长度，和一个错误值

		if err == io.EOF { // 文件结尾的io.EOF
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v/n", err)
			os.Exit(1)
		}

		// 不合法字符的统计
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++

		if unicode.IsDigit(r) {
			number++
		} else if unicode.IsLetter(r) {
			letter++
		} else {
			other++
		}
	}

	// 输出不同字符的个数
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	// 输出各个长度的有多少个
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	// 输出invalid
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

	// 习题：输出不同类型的
	fmt.Printf("\nletter\tdigit\tother\n")
	fmt.Printf("%d\t%d\t%d\n", letter, number, other)
}

/*
中国
2 3
rune    count
'中'    1
'国'    1
'\n'    2
'2'     1
' '     1
'3'     1

len     count
1       5
2       0
3       2
4       0
*/
