// 第 1 课：第一个 Go 程序 + 变量与类型
//
// 一个 Go 文件永远是这三段结构：
//  1. package  —— 声明这个文件属于哪个包（一个目录 = 一个包）
//  2. import   —— 引入要用到的其他包
//  3. func     —— 具体的函数
package main // main 是特殊包：Go 只能从 main 包里的 main 函数启动程序

import "fmt" // fmt = format，负责格式化输入输出。Go 没有内置的 print 语句，打印全靠它

// main 是程序入口，签名固定：没有参数、没有返回值
// 规则：整个 main 包里只能有一个 main 函数，多写一个就报 main redeclared
func main() {
	// 作业 1：打印一句话
	// Println 的特点：多个参数用空格分隔，末尾自动换行
	fmt.Println("你好，我是 Joe，正在从零学 Go")

	// 作业 2：变量与类型
	// := 是短声明，一步完成"声明 + 赋值 + 推断类型"，只能在函数内用
	// 它等价于 var name string = "joe"，但类型不用手写
	name := "joe"     // string：字符串，必须用双引号
	age := 20         // int：整数。Go 整数默认推断成 int（不是 int32/int64）
	height := 1.78    // float64：小数默认推断成 float64（不是 float32）
	isStudent := true // bool：只有 true / false 两个值

	// Println 会把每个参数按顺序打印，中间加空格
	fmt.Println(name, age)

	// Printf 按"占位符"格式化输出，不会自动换行，所以要自己写 \n
	//   %s  字符串    %d  整数    %.2f  保留两位小数    %v  万能值（什么类型都能打）
	// 占位符个数和后面的参数个数必须一一对应，少一个会输出 %!d(MISSING) 这种怪东西
	fmt.Printf("我叫%s，今年%d岁，身高%.2f米，是学生：%v\n", name, age, height, isStudent)

	// 作业 3：变量声明了却不用 —— 编译直接报错，这是 Go 的零容忍规则
	// 去掉下面这行的注释试试，报错信息是：declared and not used: unused
	// var unused = 1
}
