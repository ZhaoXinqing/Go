// Go语言通过内置的错误接口提供了非常简洁的错误处理机制
// err是一个接口类型，它的定义为：
type error interface {
	Error() string
}
// 我们可以在编码中通过error接口类型来生成错误信息。
// 函数通常在最后的返回值中返回错误信息。
// 使用errors.New 可返回一个错误信息，示例：
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
}