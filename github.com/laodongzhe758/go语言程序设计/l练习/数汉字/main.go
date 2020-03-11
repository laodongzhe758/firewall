//编写代码统计出字符串""中汉字的数量。

//判断是否为汉字
// unicode.Is(unicode.Han, c)
// unicode.Is(unicode.Scripts["Han"], c)
package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "勤劳的劳动者hjakhdfjas132143125"
	num := 0
	for _, c := range s {
		if unicode.Is(unicode.Han, c) {  
			num++
		}
	}
	fmt.Printf("%s中有%d个汉字", s, num)
}
