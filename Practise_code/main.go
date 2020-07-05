package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

line := inpit.Text()
counts[line] = counts[line] + 1

input := buffio.NewScanner(os.Stdin)

package main 

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",Handle)
	log.Fatal(http.ListenAndServer("localhost:8080",nil))
}

func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}


f, err := os.Open(name)
if err != nil {
	return err
}
f.Close()

in, err := os.Open(file)
// ...
out, err := os.Crate(outfile)

p := new(int)
fmt.Println(*p)

age := make(map[string]int)

ages := map[string]int {
	"alice": 31,
	"charlie": 34
}

ages["bob"] = ages["bob"] + 1
ages["bob"] += 1
ages["bob"] ++


for name, age := range ages {
	fmt.Printf("%s\t%d\n", name, age)
}

import "sort"

var names []string
for name := range ages {
	names = append(names,name)
}
sort.Strings(names)
for _, name := range names {
	fmt.Printf("%s\t%d\n",name, ages[name])
}

gopl.io/ch5/title1
func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// Check Content-Type is HTML (e.g., "text/html;charset=utf-8").
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct,"text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html",url, ct)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url,err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title"&&n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}

func title(url string) err {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct,"text/html;") {
		return fmt.Errorf("%s has type %s, not text/html",url,ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url,err)
	}
	// ...print doc's title element
}
io/ioutil
package ioutil
func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadAll(f)
}

// 互斥锁
var mu sync.Mutex
var m = make(map[string]int)
func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}
Reader/ Closer
