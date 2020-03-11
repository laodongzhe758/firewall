// Copyright © 2010-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//分3部分了解这个示例程序：首先介绍import部分，然后是静态数据，
//再之后是程序处理过程
package main

import (
	"fmt" //fmt包提供了格式化文本和读取格式化文本的相关函数（参见3.5节）。
	"log" //log包提供了日志功能。
	"os"  //os 包提供的是平台无关的操作系统级别变量和函数，包括用于保存命令行参数的类型为[]string的os.Args变量（即字符串类型的切片）。

	//os.Args[0]存放的是程序名字，
	"path/filepath" //而path包中的filepath子包则提供了一系列可跨平台的对文件名和路径操作的函数。需要
	// 注意的是，对于位于其他包内的子包，在我们的代码中用到时只需要指定其包名称
	// 的最后一部分即可（对于此例而言就是filepath）。
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		//filepath.Base()函数会返回传入路径的基础名（其实就是文件名）。
		// 输出消息后，程序通过调用os.Exit函数退出，返回1给操作系统.
		// 在类Unix系统中，程序返回0表示成功，非零值表示用法问题或执行失败。
		os.Exit(1)
	}

	stringOfDigits := os.Args[1]
	for row := range bigDigits[0] { //ROW（行）与COLUMN（列）
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			a := 1
			//当我们按索引位置查询一个字符串的内容时，我们将得到索引位置对应的一个
			// byte类型的值（在Go语言中，byte类型等同于uint8类型）。所以，我们可以对命
			// 令行传入的参数按索引位置取相应的byte类型值，然后将该值和数字0对应的byte
			// 类型值相减，以得知对应的数字。在UTF-8和ASCII中，字符‘0’对应的是48，字
			// 符‘1’对应的是49，以此类推。因此，假如我们得到的是一个字符‘3’（对应数
			// 值为51），那么我们可以通过运算‘3’-‘0’（也就是51-48）来获取相应的整型
			// 值，也就是一个byte类型的整型数，值为3。
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + "  "
			} else {
				log.Fatal("invalid whole number")
				//我们调用log.Fatal()函
				// 数记录一条错误信息，包括日期、时间和错误信息，如果没有显式指定记录到哪
				// 里，那么默认是打印到os.Stderr，并调用os.Exit(1)终止程序的执行。另外还有
				// 一个log.FatalF()函数可以接受%格式的占位符。在第一个if语句里我们没有使用
				// log.Fatal()函数，因为我们只需要输出程序的帮助信息，而不需要日期和时间这
				// 些通常log.Fatal()函数的输出会包含的信息。

			}
			fmt.Println(a)
		}
		fmt.Println(line)
	}

}

var bigDigits = [][]string{
	{"  000  ", " 0   0 ", "0     0", "0     0", "0     0", " 0   0 ", "  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
	{"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ", "   4  "},
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	{" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}

// 通常情况下声明和定义的顺序并不会带来影响。因此在
// bigdigits/bigdigits.go文件中，我们可以在main()函数前后声明bigDigits变
// 量。在这个例子里，我们将main()函数放在前面，因为本书所有的例子我们都趋向
// 于用自上而下的方式来组织内容。
