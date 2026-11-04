// 第 5 课：错误处理与接口
//
// 两个核心观念：
//  1. Go 没有 try-catch，错误是一个普通返回值（error 类型）
//  2. Go 没有 implements，只要"会这个方法"就自动满足接口（鸭子类型）
package main

import (
	"errors"
	"fmt"
	"math"
)

// divide 演示"错误是返回值"：出错时返回 error，不出错时返回 nil
// 惯例：error 永远是最后一个返回值；printf 时 err 会自动调用它的 Error()
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil // nil 表示"没有错误"
}

// Shape 接口：一块写着"会算面积才能进"的牌子
// 接口只规定"要会什么"，不关心"你是谁"
type Shape interface {
	Area() float64
}

type Circle struct {
	R float64
}

// Circle 只是老老实实写了个 Area 方法，就自动成了 Shape
// 注意：没有写任何 "implements Shape" 之类的话
func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R // math.Pi 比手写 3.14 精确
}

type Rect struct {
	W, H float64
}

func (r Rect) Area() float64 {
	return r.W * r.H
}

// printArea 只认 Shape 这块牌子，不关心传进来的是圆还是矩形
func printArea(s Shape) {
	fmt.Println(s.Area())
}

// NotFoundError 自定义错误类型：给错误加了一个"格子"
// 普通错误（errors.New）只有一句话，这个能带结构化信息
type NotFoundError struct {
	Resource string
}

// Error 让 NotFoundError 学会"说明自己是什么错" —— 只有这样它才算 error
// 和 Circle 有 Area() 就是 Shape，是同一个规则
func (e NotFoundError) Error() string {
	return e.Resource + " 不存在"
}

// findUser 找不到人时返回自定义错误
func findUser(id int) (string, error) {
	if id != 1 {
		return "", NotFoundError{Resource: "用户"}
	}
	return "joe", nil
}

func main() {
	// ① 接口：同一个函数处理两种不同的形状
	printArea(Circle{R: 3})     // 28.274...
	printArea(Rect{W: 4, H: 5}) // 20

	// ② 普通错误：用 errors.New 造一个
	// 这里不用 return —— main 里 return 会结束整个程序，后面的代码就不跑了
	_, err := divide(4, 0)
	if err != nil {
		fmt.Println(err)
	}

	// ③ 自定义错误：先判断"是不是这个类型"，是的话就能读格子
	// err.(NotFoundError) 和 map 的 v, ok 是同一个套路：第二个返回值表示"成不成功"
	_, err = findUser(99)
	if err != nil {
		e, ok := err.(NotFoundError)
		if ok {
			fmt.Println("缺的是:", e.Resource)
		} else {
			fmt.Println("其他错误:", err)
		}
	}
}
