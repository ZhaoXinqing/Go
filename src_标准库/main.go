package main

import (
	"fmt"
)


func main() {
	// 返回不小于x的最小整数（的浮点值）
	// 计算数字只入不舍
	math.Ceil(12.34)

	// 返回不大于x的最小整数（的浮点值）
	// 计算数字只舍不入
	math.Floor(12.34)

	// 四舍五入
	math.Round(12.34)

	// 返回x的整数部分（的浮点值）
	math.Trunc(12.34)

	// 返回f的整数部分和小数部分，结果的正负号都相同
	math.Modf(12.34)
		// 返回x的绝对值
	math.Abs(-12.34) // 12.34

	// 返回x和y中最大值
	math.Max(12.3, 22)

	// 返回x和y中最小值
	math.Min(12.3, 22)
	// 返回x的绝对值
	math.Abs(-12.34) // 12.34

	// 返回x和y中最大值
	math.Max(12.3, 22)

	// 返回x和y中最小值
	math.Min(12.3, 22)

	// 返回x的二次方根
	math.Sqrt(144) // 12

	// 返回x的三次方根
	math.Cbrt(27)
	// 返回x**y
	math.Pow(3, 4)

	// 返回10**e
	math.Pow10(2)
	
	// 使用给定的seed来初始化生成器到一个确定的状态
	// 如未调用，默认资源的行为就好像调用了Seed(1)
	// 如果需要实现随机数，seed值建议为 时间戳
	// 如果seed给定固定值，结果永远相同
	rand.Seed(time.Now().UnixNano())

	// 返回一个非负的伪随机int值
	fmt.Println(rand.Int())

	// 返回一个取值范围在[0,n]的伪随机int值，如果n<=0会panic
	fmt.Println(rand.Intn(100))

	// 返回一个int32类型的非负的31位伪随机数
	fmt.Println(rand.Int31())
}
