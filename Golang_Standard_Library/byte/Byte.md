
1. bytes  bytes包实现了操作[]byte的常用函数.
1. encoding	encoding包定义了供其它包使用的可以将数据在字节水平和文本表示之间转换的接口.
1. strconv	strconv包实现了基本数据类型和其字符串表示的相互转换.
1. strings	strings包实现了用于操作字符的简单函数.
1. text	scanner对utf-8文本的token扫描服务，tabwriter缩进修正
1. unicode	unicode 包提供了一些测试Unicode码点属性的数据和函数.

## golang中的bytes包 ##

golang标准库中提供了bytes包，该包几乎和strings包给string提供的功能，只不过bytes包对应的是[]byte。和strings一样，并不修改传入变量，而是返回其副本修改之后的内容。

整个包分为以下几种操作：**转换、比较、去除、分割、查找、替换**

## 一、转换 ##
将s的副本中所有字符修改成大写（小写），然后返回

	func ToUpper(s []byte) []byte  //  全部转大写
	func ToLower(s []byte) []byte  // 全部转小写
	func ToTitle(s []byte) []byte  // 全部转大写
使用指定的映射表将对s的副本中的所有字符进行修改，然后返回

	func ToUpperSpecial(_case unicode.SpecialCase, s []byte) []byte
	func ToLowerSpecial(_case unicode.SpecialCase, s []byte) []byte
	func ToTitleSpecial(_case unicode.SpecialCase, s []byte)
将s的副本中的所有单词的首字符修改为Title格式返回。

	func Title(s []byte) []byte

## 二、比较 ##
比较两个[]byte，nil参数相当于[]byte{}，

	a < b  返回-1
	a == b 返回0
	a > b   返回1
	func Compare(a, b []byte) int
判断a、b是否相等，nil参数相当[]byte{}。

	func Equal(a, b []byte) bool
判断s、t是否相似，忽略大写、小写、标题三种格式的区别。

	func EqualFold(s, t []byte) bool
## 三、去除 ##
去除s的副本左右包含cutset中的字符（返回s的副本）

	func Trim(s []byte, cutset string) []byte
	func TrimLeft(s []byte, cutset string) []byte
	func TrimRight(s []byte, cutset string) []byte
去掉s副本的左右边中f函数中返回true的字符

	func TrimFunc(s []byte, f func(r rune) bool) []byte
	func TrimLeftFunc(s []byte, f func(r rune) boo) []byte
	func TrimRightFunc(s []byte, f func(r rune) bool) []byte
去除s副本两边的空白（unicode.IsSpace)

	func TrimSpace(s []byte) []byte
去掉s副本的前缀prefix（后缀suffix）

	func TrimPrefix(s, prefix []byte) []byte
	func TrimSuffix(s, suffix []byte) []byte
## 四、分割 ##
Split函数以sep为分隔符将s的副本切割分成多个子串，结果不包含分隔符。
如果sep为空，则将s的副本分割成Unicode字符列表。
SplitN可以指定分割次数，超出n的部分将不进行分割，n小于0，表示分割所有。

	func Split(s, sep []byte) [][]byte
	func SplitN(s, sep []byte, n int) [][]byte
功能同Spelit，只不过结果包含分隔符（在各个子串尾部）

	func SplitAfter(s, sep []byte) [][]byte
	func SplitAfterN(s, sep []byte, n int) [][]byte
以连续空白为分隔符将s的副本分隔成多个子串，结果不包含分隔符。

	func Fields(s []byte) [][]byte
以符合f的字符为分隔符将s的副本分割成多个子串，结果不包含分割符。

	func FieldsFunc(s []byte, f func(rune) bool) [][]byte
以sep为连接符，将子串列表s的副本连接成一个字节串。

	func Join(s [][]byte, sep []byte) []byte
将子串b重复count次后返回。

	func Repeat(b []byte, count int) []byte

## 五、查找 ##

判断s是否有前缀prefix（后缀suffix）

	func HasPrefix(s, prefix []byte) bool
	func HasSuffix(s, suffix []byte) bool
判断b中是否包含子串subslice（字符r）

	func Contains(b, subslice []byte) bool
	func ContainsRune(b []byte, r rune) bool
判断b中是否包含chars中的任何一个字符

	func ContainsAny(b []byte, chars string) bool
查找子串sep（字节c、字符r）在s中第一次出现的位置，找不到则返回-1。

	func Index(s, sep []byte) int
	func IndexByte(s []byte, c byte) int
	func IndexRune(s []byte, r rune) int
查找chars中的任何一个字符在s中第一次出现的位置，找不到则返回-1。

	ndexAny(s []byte, chars strings) int
查找符合f的字符在s中第一次出现的位置，找不到则返回-1。

	func IndexFunc(s []byte, f func(r rune) bool) int
功能同上，只不过查找最后一次出现的位置。

	func LastIndex(s, sep []byte) int
	func LastIndexByte(s []byte, c byte) int
	func LastIndexAny(s []byte, chars string) int
	func LastIndexFunc(s []byte, f func(r rune) bool) int
获取sep在s中出现的次数

	func Count(s, sep []byte) int
## 六、替换 ##
将s副本中的前n个old替换为new，n<0则替换全部。

	func Replace(s, old, new []byte, n int) []byte
将s副本中字符替换为mapping(r)的返回值，如果mapping返回负值，则丢弃该字符。

	func Map(mapping func(r rune) rune, s []byte) []byte
将s副本转换为[]rune类型返回

	func Runes(s []byte) []rune
	<h3>type Reader</h3>
将切片b封装成bytes.Reader对象

	wReader(b []byte) *Reader

**bytes.Reader实现了如下接口：**

	1）io.ReadeSeeker
	2）io.ReaderAt
	3）io.WriterTo
	4）io.ByteScanner
	5）io.RuneScanner

	// 返回未读取部分的数据长度
	func (r *Reader) Len() int

	// 返回底层数据的总长度，方便ReadAt使用，返回值不变。
	func (r *Reader) Size() int64


type Buffer struct {...}
将buf包装成bytes.Buffer对象。
	func NewBuffer(buf []byte) *Buffer
转化成bytes.Buffer对象

	func NewBufferString(s string) *Buffer
**Buffer是一个缓存，没有底层数据，缓存的容量会根据需要自动调整**。大多数情况下，使用new(Buffer)就足以初始化一个Buffer了。
bytes.Buffer实现了如下接口：

	1）io.ReadWrite
	2）io.ReaderFrom
	3）io.WriterTo
	4）io.ByteWriter
	5）io.ByteScanner
	6）io.RuneScanner
未读取部分的数据长度

	func (b *Buffer) Len() int
缓存的容量

	func (b *Buffer) Cap() int
读取前n字节的数据并以切片形式返回，如果数据长度小于n，则全部读取。切片只在下一次读写操作前合法。
	
	func (b *Buffer) Next(n int) []byte
读取第一个delim及其之前的内容，返回遇到的错误（一般是io.EOF）。

	func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)
	func (b *Buffer) ReadString(delim byte) (line string, err error)
写入r的UTF-8编码，返回写入的字节数和error。
保留err是为了匹配bufio.Write的WriteRune
	
	func (b *Buffer) WriteRune(r rune) (n int, err error)
写入s，返回写入的字节数和error。

	func (b *Buffer) WriteString(s string) (n int, err error)
引用未读取部分的数据部分的数据切片（不移动读取位置）

	func (b *Buffer) Bytes() []byte
返回未读取部分的数据字符串（不移动读取位置）

	func (b *Buffer) String() string
自动增加缓存容量，以保证有n字节的剩余空间。
如果n小于0或无法增加则会panic。

	func (b *Buffer) Grow(n int)
将数据长度截短到n字节，如果n小于0或大于Cap则panic。

	func (b *Buffer) Teuncate(n int)
重设缓冲区，清空所有数据（包括初始内容）。

	func (b *Buffer) Reset