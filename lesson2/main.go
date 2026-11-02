// 第 2 课：函数、if、for
package main

import "fmt"

// max 返回两个整数里较大的那个
//
// 语法要点：
//   - 类型写在变量名后面：a, b int，同类型可以合并写
//   - 返回值类型写在参数列表后面：最后的那个 int
func max(a, b int) int {
	if a > b {
		return a // return 一执行，函数立刻结束，后面的代码不会跑
	}
	return b // 所以这里不需要 else —— "提前返回"是 Go 的惯用写法，比 if/else 嵌套清晰
}

// sumTo 用 for 计算 1 + 2 + ... + n
func sumTo(n int) int {
	sum := 0 // 累加器，初值必须是 0
	// for 的三段式：初始化; 循环条件; 每轮结束后的操作
	// 注意：i 从 0 开始会多跑一轮（加个 0 不影响结果），语义上更准确的是 i := 1
	for i := 0; i <= n; i++ {
		sum += i // 等价于 sum = sum + i
	}
	return sum
}

// isEven 判断是否为偶数
//
// n%2 == 0 本身就是一个 bool 值，所以直接 return，不需要写成 if...return true else return false
func isEven(n int) bool {
	return n%2 == 0 // % 取余：余数为 0 就是偶数
}

// divmod 一次返回两个值：商和余数
//
// 多返回值是 Go 的核心特色之一，后面写错误处理全靠它：
//
//	result, err := doSomething()
//
// 返回值列表用括号包起来：(int, int)
func divmod(a, b int) (int, int) {
	return a / b, a % b // 整数除法会丢掉小数部分，7/2 = 3
}

func main() {
	// 函数调用：把返回值接到变量里
	result := max(6, 4)
	fmt.Println("最大值是:", result)

	// 单个返回值可以直接作为参数传给 Println
	fmt.Println(sumTo(100))

	// Go 没有 while、没有 do-while，所有循环都用 for 表达
	// 这里：i 从 1 递增到 20，是偶数的就打印
	for i := 1; i <= 20; i++ {
		if isEven(i) { // if 的条件不加括号，但结果必须是 bool，不能写 if 1
			fmt.Println(i)
		}
	}

	// 多返回值直接传给 Println：会被当成多个参数，输出 "3 1"
	// 但 Println(divmod(7,2), 1) 会编译失败，报 multiple-value in single-value context
	// 规则：多返回值只能独占一个参数位置，不能和其他值混着传
	fmt.Print(divmod(7, 2))
}
