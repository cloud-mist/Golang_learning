// 格式化字符
package main

import "fmt"
func main() {
	// 1.各符号示例
	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 15)
	fmt.Printf("A string: %s\n", "hello")
	fmt.Printf("A boolean: %t\n", false)
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)	// 根据提供类型选择合适格式 
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true) // 根据代码中显示的格式进行格式化
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n")
	
	fmt.Println()
	// 2.可以在%后添加一个数字表示最小宽度，若比最小宽度少，就用空格填充。
	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Printf("-----------------------------\n")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n","Paper Clips", 5)
	fmt.Printf("%12s | %2d\n","Tape", 99)

	fmt.Println()
	// 3.小数位数保留
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456)
	fmt.Printf("%%.2f: %.2f\n", 12.3456)
	fmt.Printf("%%.1f: %.1f\n", 12.3456)

}

/*
A float: 3.141500
An integer: 15
A string: hello
A boolean: false
Values: 1.2	  true
Values: 1.2 "\t" true
Types: float64 string bool
Percent sign: %     

	 Product | Cost in Cents
-----------------------------
      Stamps | 50
 Paper Clips |  5
        Tape | 99

%7.2f:   12.35
%7.1f:    12.3
%.2f: 12.35
%.1f: 12.3
*/

