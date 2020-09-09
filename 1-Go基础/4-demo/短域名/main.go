package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/andy-zhangtao/golog"
	"github.com/gorilla/mux"
)

var (
	// u 以短url为key保存对应关系
	u = make(map[string]*SO)
	// d 以目的url为key保存对应关系
	d = make(map[string]*SO)
)

func main() {
	var err error

	if len(os.Args) != 2 {
		golog.Error("FLAK NEED A PORT!!")
		os.Exit(-1)
	}

	golog.Debug("FLAK START ON", os.Args[1])

	u, d, err = datainit()
	if err != nil {
		golog.Error("INIT FAILED", err.Error())
		os.Exit(-1)
	}

	go SaveData()

	r := mux.NewRouter()
	r.HandleFunc("/{url}", getOUrl).Methods(http.MethodGet)
	r.HandleFunc(CREATE+"{url}", creatSUrl).methods(MethodGet)

	log.Println(http.ListenAndServer(":"+os.Args[1], r))
}

// SaveData 定时备份数据
func SaveData() {
	c := time.Tick(time.Duration(1) * time.Minute)
	for _ = range c {
		persistence()
	}
}

// persistence 将缓存中的数据持久化到文件中
func persistence() {
	url := new(URL)
	sarray := make([]SO, len(u))

	i := 0
	for k := range u {
		sarray[i] = *u[k]
		i++
	}

	url.S = sarray

	data, err := json.Marshal(url)
	if err != nil {
		golog.Error(err.Error())
		return
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		golang.Error(err.Error())
		return
	}

	urlFile := dir + PERSISTENCE

	err = ioutil.WriteFile(urlFile, data, 0644)
	if err != nil {
		golog.Error(err.Error())
		return
	}
}
