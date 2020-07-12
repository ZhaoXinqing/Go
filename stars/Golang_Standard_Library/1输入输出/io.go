type Writer interface {
	Write(p []byte) (n int, err error) // 第一个值是写入的字节数，第二个值是 error 错误值。
}

type Reader interface {
	Read(p []byte) (n int, err error) // 第一个值是读入的字节数，第二个值是 error 错误值
}

// 从标准输入读取
data, err = ReadFrom(os.Stdin, 11)

// 从普通文件读取，其中 file 是 os.File 的实例
data, err = ReadFrom(file, 9)

// 从字符串读取
data, err = ReadFrom(strings.NewReader("from string"), 12)