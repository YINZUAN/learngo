// 第 4 课：map 与结构体
//
// map  = 哈希表，按键查值，平均 O(1)；键必须是可比较类型（不能是 slice/map/func）
// struct = 字段的集合，Go 没有 class，用结构体 + 方法代替
package main

import "fmt"
import "strings"

// countWords 统计每个单词出现的次数
//
// 思路：切词 → 遍历 → 用 map 计数
// 关键：count[word]++ 能一行搞定，是因为读一个不存在的键会返回零值 0，
// 0+1=1 再写回去，不需要先判断键在不在（其他语言通常要写 if/else 三行）
func countWords(s string) map[string]int {
	words := strings.Fields(s)    // 按空白切分成 []string
	count := make(map[string]int) // 键是单词，值是次数
	for _, word := range words {
		count[word]++ // 自增：不存在时先读成 0
	}
	return count
}

// Student 学生信息
//
// 注意导出规则：字段名首字母大写 = 包外可见（json 序列化也只看得到大写字段）
// 小写开头的字段是"私有"的，别的包读不到
type Student struct {
	Name  string
	Score float64
}

// IsPass 判断是否及格
//
// 值接收者 (s Student)：拿到的是副本，只读不改，适合这种"查询"方法
func (s Student) IsPass() bool {
	return s.Score >= 60
}

// AddScore 加分
//
// 指针接收者 (s *Student)：拿到的是地址，改的是原对象
// 如果这里写成值接收者，s.Score += n 改的只是副本，外面的学生分数不会变
func (s *Student) AddScore(n float64) {
	s.Score += n
}

func main() {
	// ── 练习 1：单词计数 ──
	fmt.Println(countWords("go is great go is fast"))

	// ── 练习 2：成绩表（增、删、查）──
	scores := map[string]int{"joe": 85, "amy": 92}
	scores["joe"] = 100    // 键存在 → 覆盖
	delete(scores, "amy")  // 删掉 amy
	_, ok := scores["tom"] // comma-ok：ok 表示键是否存在
	if !ok {               // 只需要"存不存在"，所以第一个返回值用 _ 丢掉
		fmt.Println("不存在")
	}
	fmt.Println(scores) // map[joe:100]

	// ── 练习 3：结构体 + 两种接收者 ──
	// 故意用 55 分：先不及格，加分后及格，才能看出 AddScore 真的改了原对象
	s := Student{Name: "joe", Score: 55}
	fmt.Println("及格了吗:", s.IsPass()) // false

	s.AddScore(9)                    // s 是值，但方法是 *Student —— Go 自动帮你取地址 (&s).AddScore(9)
	fmt.Println("加分后是:", s.Score)    // 64
	fmt.Println("及格了吗:", s.IsPass()) // true
}
