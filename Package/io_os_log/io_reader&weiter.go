type Writer interface {
	Write(p []byte) (n int, err error) // 第一个值是写入的字节数，第二个值是 error 错误值。
}

type Reader interface {
	Read(p []byte) (n int, err error) // 第一个值是读入的字节数，第二个值是 error 错误值
}
