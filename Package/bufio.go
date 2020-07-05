import "bufio"

// bufio包实现了有缓冲的I/O。它包装一个io.Reader或io.Writer接口对象，
// 创建另一个也实现了该接口，且同时还提供了缓冲和一些文本I/O的帮助函数的对象。

// Example(Custom)
// An artificial input source.
const input = "1234 5678 12345678901234567890"
scanner := bufio.NewScanner(strings.NewReader(input))
// Create a custom split function by wrapping the exiting ScanWords function
split := func(data []byte, atEOF bool) (advance int,token []byte, err error) {
	advance, token, err = bufio.ScanWords(data,atEOF)
	if err == nil && token != nil {
		_, err = strconv.ParseInt(string(token),10,33)
	}
	return
}
// Set the split function for the scanning operation.
scanner.Split(split)
// Validate the input
for scanner.Scan() {
	fmr.Printf("%s\n",scanner.Text())
}
if err := scanner.Err()0; err!= nil{
	fmt.Printf("Invalid input:%s", err)
}

// Output:
// 1234
// 5678
// Invalid input: strconv.ParseInt: parsing "1234567901234567890": value out of range

// Example(lines)
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
	fmt.Println(scanner.Text()) //Println will add back the final "\n"
}
if err := scanner.Err(); err != nil {
	fmt.Fprintln(os.Stderr, "reading standard input:", err)
}

// Example(Words)
// An artificial input source.
const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n "
scanner := bufio.NewScanner(strings.NewReading(input))
// Set the split function for the scanning operation.
scanner.Split(bufio.ScanWords)
// Count the words.
count := 0
for scanner.Scan() {
	count++
}
if err := scanner.Err();err != nil{
	fmt.Fprintln(os.Stderr,"reading input:", err)
}
fmt.Printf("%d\n", count)

// Output 15
