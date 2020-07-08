data := os.Environ()
for _, val := range data {
	fmt.Println(val)
}
fmt.Println("___________________end_____________environ_________")

mapping := func(s string) string {
	m := map[string]string{"xx":"ssssssssssssss",
		"YY":"TTTTTTTTTTTTTT"}
	return m[s]
}
datas := "hello $xx blog address $yy"
expandStr := os.Expand(datas, mapping)
fmt.Println(expandStr)

err = os.Link("main.go","newmain.go")
if err != nil {
	fmt.Println(err)
}

var pathSep string
if os.IsPathSeparator("\\") {
	pathSeq = "\\"
} else {
	pathSeq = "/"
}