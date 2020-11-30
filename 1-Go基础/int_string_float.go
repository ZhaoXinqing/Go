// Go int、int64、string类型转换

// string 到 int
int, err: = strconv.Atoi(string)

// string 到 int64
// 第三个参数位大小表示期望转换的结果类型，其值可以为0, 8, 16, 32和64，分别对应 int, int8, int16, int32和int64
int64, err := strconv.ParseInt(string, 10, 64)


// int 到 string
// 等价于 string := strconv.FormatInt(int64(int), 10)
string := strconv.Itoa(int)

// int64 到 string
string := strconv.FormatInt(int64, 10)

// int 到 int64
int64_ := int64(1234)

// float 到 string
string := strconv.FormatFloat(float32, 'E', -1, 32)
string := strconv.FormatFloat(float64, 'E', -1, 64)

// string 到 float
float, err := strconv.ParseFloat(string, 64)
float, err := strconv.ParseFloat(string, 32)