package main

import (
	"fmt"
	"time"
	"os"
	"path/filepath"
)

func main() {
	var myDate string

	if len(os.Args) != 2 {
		fmt.Println("Usage:%s data\n",
			filepath.Base(os.Args[0]))
		os.Exit(0)
	}
	myDate = os.Args[1]

	d,err := time.Parse("02 January 2006",myDate)
	if err == nil {
		fmt.Println("Full",d)
		fmt.Println("Time",d.Day(),d.Month(),d.Year())
	} else {
		fmt.Println(err)
	}
}

