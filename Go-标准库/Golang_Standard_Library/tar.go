package tar

import (
	"fmt"
	"tar"
)

buf := new(bytes.Buffer)
tw := tar.NewWriter(buf)
var files = []struct{
	Name, Body string
}{
	{"readme.txt", "This archive contains some text files."},
    {"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
    {"todo.txt", "Get animal handling licence."},
}
for _, file := range files{
	hdr := &tar.Header{
		Name: file.Name
		Size: int64(len(file.Body)),
	}
	if err := tw.WriteHeader(hdr);err != nil {
		log.Fatalln(err)
	}
}

// make sure to check the error on Close
if err := tw.Close(); err != nil {
	log.Fatalln(err)
}
// open the tar archive for reading.
r := bytes.NewReader(buf.Bytes())
tr := tar.NewReader(r)
// Iterate through the file in the archive
for {
	hdr, err := tr.Next()
	if err == io.EOF {
		// end of tar archive
		break
	}
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Content of %s:\n", hdr.Name)
	if _, err := io.Copy(os.Stdout, tr); err != nil {
		log.Fatalln(err)
	}
	fmt.Println()
}

// Writer实现了为io.Writer接口对象提供缓冲。如果在向一个Writer类型值写入时遇到了错误，
// 该对象将不再接受任何数据，且所有写操作都会返回该错误。
// 在说有数据都写入后，调用者有义务调用Flush方法以保证所有的数据都交给了下层的io.Writer。
type Writer struct {
	// 内含隐藏或非导出字段
}
w := bufio.New.Writer(os.Stdout)
fmt.Fprint(w, "hello,")
fmt.Fprint(w, "world")
w.Flush() // Don't forget to flush