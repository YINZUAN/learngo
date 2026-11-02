// 第 3 课：切片（slice）
//
// 一句话心智模型：切片不是数组，是「指向底层数组的指针 + len + cap」三个字段。
//
//	len = 现在有几个元素（能读的范围）
//	cap = 从起点到数组末尾还剩几个位置（append 能直接写的范围）
package main

import "fmt"

// reverse 原地反转切片（挑战题 1）
//
// 思路：双指针，一头一尾，交换之后往中间收
// 注意：没有返回值。切片是引用式的，函数内改的就是外面那块底层数组
func reverse(s []int) {
	left, right := 0, len(s)-1 // left 指向头部，right 指向最后一个元素
	// 用 left < right 而不是 <= ：奇数长度时中间那个元素不用和自己交换
	// 边界自动安全：空切片 len=0 → right=-1 → 条件不成立，循环根本不进
	for left < right {
		// Go 支持多重赋值：右边先全部求值，再一起赋给左边
		// 所以交换两个值不需要临时变量（Java 里得写三行）
		s[left], s[right] = s[right], s[left]
		left++  // 左指针往右挪
		right-- // 右指针往左挪
	}
}

// unique 去重并保持原顺序（挑战题 2）
//
// 思路：用 map 记住"见过谁"，没见过就追加到结果里
// 时间复杂度 O(n)：查 map 平均是 O(1)；如果不用 map 而用两层循环，就是 O(n²)
func unique(nums []int) []int {
	// map[键类型]值类型。这里是「数字 -> 见没见过」
	// 关键特性：读一个不存在的键返回零值（bool 的零值是 false），所以下面能直接 if seen[n]
	seen := make(map[int]bool)
	// make 的三个参数含义：类型、长度 0（还没有元素）、容量 len(nums)
	// 预留容量是为了避免 append 过程中反复扩容 —— 第 3 课讲的 cap 在这里派上用场
	result := make([]int, 0, len(nums))
	for _, n := range nums { // _ 表示丢弃下标，只要元素值
		if seen[n] { // 见过 → 跳过
			continue
		}
		seen[n] = true             // 标记为"见过"
		result = append(result, n) // 追加到结果末尾，这样天然保持了原顺序
	}
	return result
}

func main() {
	// ── 练 1：三种创建方式与 len / cap ──

	a := []int{1, 2, 3}     // 切片字面量：没写容量，len=3 cap=3
	b := make([]int, 3)     // 第二个参数是长度：len=3 cap=3，元素全被填成零值 0
	c := make([]int, 3, 10) // 第三个参数是容量：len=3 cap=10，预留了 7 个空位

	fmt.Println("字面量的len是:", len(a), "字面量的cap是:", cap(a)) // 3 / 3
	fmt.Println("b的len是:", len(b), "b的cap是:", cap(b))     // 3 / 3
	fmt.Println("c的len是:", len(c), "c的cap是:", cap(c))     // 3 / 10

	// ── 练 2：切片的切片共享同一块底层数组 ──

	arr := []int{1, 2, 3, 4, 5}
	// arr[2:4] 是左闭右开：取下标 2 和 3，也就是 [3 4]
	// len = 4-2 = 2，cap = 原来 5 - 起点 2 = 3
	// 注意：这里没有复制任何数据，sub 和 arr 指向同一块内存
	sub := arr[2:4]
	sub[1] = 100     // sub[1] 就是 arr[3]，改它等于改 arr
	fmt.Println(arr) // [1 2 3 100 5] —— 原切片被改了

	// ── 练 3-A：append 时 cap 够 → 原地写，仍然共享 ──

	s := make([]int, 2, 5) // len=2 cap=5，还能再塞 3 个
	t := append(s, 1)      // 加 1 个：需要的长度 3 <= cap 5，不扩容，写到同一块内存
	t[0] = 88              // 改 t 等于改 s 的底层数组
	fmt.Println(s)         // [88 0] —— s 被影响了

	// ── 练 3-B：append 时 cap 不够 → 新数组 + 拷贝，彻底脱钩 ──

	q := make([]int, 2, 4) // len=2 cap=4，只能再塞 2 个
	// 加 3 个：需要的长度 5 > cap 4 → 触发扩容（4 < 256，于是翻倍到 8）
	// 此时 p 指向一块全新的内存，和 q 再无关系
	p := append(q, 4, 5, 6)
	p[0] = 66
	fmt.Println(q)                                           // [0 0] —— q 完全没动
	fmt.Println("p:", p, "p的cap:", cap(p), "p的len:", len(p)) // p: [66 0 4 5 6] cap 8 len 5

	// 判断是否扩容的唯一条件：append 之后需要的长度 > cap
	// 这就是"同一句 append，行为完全相反"的根源

	// ── 挑战题调用 ──

	nums := []int{1, 2, 3, 4, 5}
	reverse(nums)     // 不用接收返回值，函数里改的是同一个底层数组
	fmt.Println(nums) // [5 4 3 2 1]

	fmt.Println(unique([]int{1, 2, 2, 3, 1, 4})) // [1 2 3 4]
}
