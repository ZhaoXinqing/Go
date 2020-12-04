import (
    "bufio"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "strings"
)

// read file
f, err := os.Open("index.html")
    if err != nil {
        return err
    }
    buf := bufio.NewReader(f)
    var rep = []string{"<meta name=\"testkey\" content=\"", arg1, "\" /> "}
    var result = ""
    for {
        a, _, c := buf.ReadLine()
        if c == io.EOF {
            break
        }
        if strings.Contains(string(a), "baidu-site-verification") {
            result += strings.Join(rep, "") + "\n"
        } else {
            result += string(a) + "\n"
        }
    }