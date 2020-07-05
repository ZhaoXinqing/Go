// zip包提供了zip档案文件的读写服务。
// 本包不支持跨硬盘的压缩
// Writer类型实现了zip文件的写入器
// Crate a buffer to write our archive to.
buf := zip.NewWriter(buf)
// creat a new files to the archive
w := zip.NewWriter(buf)
// Add some files to the archive.
var files = []string {
	Name, Body string
} {
	{"readme.txt", "This archive ontains some text files."},
	{"gopher.txt", "Gopher name:\nGeorge\nGeoffary\nGonzo"}
	{"todo.txt", "get animal handing licence.\nWrite more examples."}
}
for _,file := range files {
	f, err := w.Create(file.Name)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write([]byte(file.Body))
	if err != nil {
		log.Fatal(err)
	}
}
// Make sure to check the error on Close.
err := w.Close()
if err != nil {
	log.Fatal(err)
}
