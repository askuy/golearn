// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

// os包以跨平台的方式，提供了一些与操作系统交互的函数和变量。程序的命令行参数可从os包的Args变量获取；os包外部使用os.Args访问该变量。

/*
os.Args变量是一个字符串（string）的切片（slice）（译注：slice和Python语言中的切片类似，是一个简版的动态数组），切片是Go语言的基础概念，稍后详细介绍。现在先把切片s当作数组元素序列, 序列的长度动态变化, 用s[i]访问单个元素，用s[m:n]获取子序列(译注：和python里的语法差不多)。序列的元素数目为len(s)。和大多数编程语言类似，区间索引时，Go言里也采用左闭右开形式, 即，区间包括第一个索引元素，不包括最后一个, 因为这样可以简化逻辑。（译注：比如a = [1, 2, 3, 4, 5], a[0:3] = [1, 2, 3]，不包含最后一个元素）。比如s[m:n]这个切片，0 ≤ m ≤ n ≤ len(s)，包含n-m个元素。

os.Args的第一个元素，os.Args[0], 是命令本身的名字；其它的元素则是程序启动时传给它的参数。s[m:n]形式的切片表达式，产生从第m个元素到第n-1个元素的切片，下个例子用到的元素包含在os.Args[1:len(os.Args)]切片中。如果省略切片表达式的m或n，会默认传入0或len(s)，因此前面的切片可以简写成os.Args[1:]。

下面是Unix里echo命令的一份实现，echo把它的命令行参数打印成一行。程序导入了两个包，用括号把它们括起来写成列表形式, 而没有分开写成独立的import声明。两种形式都合法，列表形式习惯上用得多。包导入顺序并不重要；gofmt工具格式化时按照字母顺序对包名排序。（示例有多个版本时，我们会对示例编号, 这样可以明确当前正在讨论的是哪个。）
 */
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}